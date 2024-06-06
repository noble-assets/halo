package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/utils"
	"github.com/noble-assets/halo/utils/data"
	"github.com/noble-assets/halo/utils/mocks"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
	"github.com/stretchr/testify/require"
)

func TestReportBalance(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewAggregatorMsgServer(k)

	// ACT: Attempt to report balance with no owner set.
	_, err := server.ReportBalance(goCtx, &aggregator.MsgReportBalance{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set aggregator owner in state.
	owner := utils.TestAccount()
	k.SetAggregatorOwner(ctx, owner.Address)

	// ACT: Attempt to report balance with invalid signer.
	_, err = server.ReportBalance(goCtx, &aggregator.MsgReportBalance{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, aggregator.ErrInvalidOwner.Error())

	//
	for _, testCase := range data.EthereumRounds {
		// ACT: Attempt to report balance.
		testCase.Msg.Signer = owner.Address
		res, err := server.ReportBalance(goCtx, &testCase.Msg)

		// ASSERT: The action should've succeeded.
		require.NoError(t, err)
		_, found := k.GetRound(ctx, res.RoundId)
		require.True(t, found)
	}

	// ASSERT: All rounds can be retrieved.
	rounds := k.GetRounds(ctx)
	require.Len(t, rounds, 10)

	// ACT: Attempt to report balance with identical round.
	msg := data.EthereumRounds[len(data.EthereumRounds)-1].Msg
	msg.Signer = owner.Address
	_, err = server.ReportBalance(goCtx, &msg)
	// ASSERT: The action should've failed due to identical round.
	require.ErrorContains(t, err, aggregator.ErrAlreadyReported.Error())

	// ACT: Attempt to report balance with invalid next price.
	// https://etherscan.io/tx/0xe8154863cf89175c9ec361999ec7ddeebf7c29297ee62325f777067409071303
	msg = aggregator.MsgReportBalance{
		Signer:      owner.Address,
		Principal:   utils.MustParseInt("4600092532"),
		Interest:    utils.MustParseInt("658503"),
		TotalSupply: utils.MustParseInt("44285046691710"),
		NextPrice:   sdk.ZeroInt(),
	}
	_, err = server.ReportBalance(goCtx, &msg)
	// ASSERT: The action should've failed due to invalid next price.
	require.ErrorContains(t, err, aggregator.ErrInvalidNextPrice.Error())

	// ARRANGE: Set the next round in state.
	k.SetRound(ctx, k.GetLastRoundId(ctx)+1, aggregator.RoundData{})

	// ACT: Attempt to report balance with existing next round.
	// https://etherscan.io/tx/0x1a628856cb74de37357a35c29ec22509b72b1fc826ac7bb1020c73a99f9f80fc
	msg = aggregator.MsgReportBalance{
		Signer:      owner.Address,
		Principal:   utils.MustParseInt("4600751035"),
		Interest:    utils.MustParseInt("657348"),
		TotalSupply: utils.MustParseInt("44285680550257"),
		NextPrice:   utils.MustParseInt("103914726"),
	}
	_, err = server.ReportBalance(goCtx, &msg)
	// ASSERT: The action should've failed due to existing round.
	require.ErrorContains(t, err, aggregator.ErrAlreadyReported.Error())
}

func TestSetNextPrice(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewAggregatorMsgServer(k)

	// ACT: Attempt to set next price with no owner set.
	_, err := server.SetNextPrice(goCtx, &aggregator.MsgSetNextPrice{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set aggregator owner in state.
	owner := utils.TestAccount()
	k.SetAggregatorOwner(ctx, owner.Address)

	// ACT: Attempt to set next price with invalid signer.
	_, err = server.SetNextPrice(goCtx, &aggregator.MsgSetNextPrice{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, aggregator.ErrInvalidOwner.Error())

	// ACT: Attempt to set next price with invalid price.
	_, err = server.SetNextPrice(goCtx, &aggregator.MsgSetNextPrice{
		Signer:    owner.Address,
		NextPrice: sdk.ZeroInt(),
	})
	// ASSERT: The action should've failed due to invalid price.
	require.ErrorContains(t, err, aggregator.ErrInvalidNextPrice.Error())

	// ACT: Attempt to set next price.
	// https://etherscan.io/tx/0xfd21979418ce5e6686c624841f48d11ed241b387b08eb60e2bd361de5ed1a061
	price := utils.MustParseInt("103780600")
	_, err = server.SetNextPrice(goCtx, &aggregator.MsgSetNextPrice{
		Signer:    owner.Address,
		NextPrice: price,
	})
	// ASSERT: The action should've succeeded, and set next price in state.
	require.NoError(t, err)
	require.Equal(t, price, k.GetNextPrice(ctx))
}

func TestAggregatorTransferOwnership(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewAggregatorMsgServer(k)

	// ACT: Attempt to transfer ownership with no owner set.
	_, err := server.TransferOwnership(goCtx, &aggregator.MsgTransferOwnership{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set aggregator owner in state.
	owner := utils.TestAccount()
	k.SetAggregatorOwner(ctx, owner.Address)

	// ACT: Attempt to transfer ownership with invalid signer.
	_, err = server.TransferOwnership(goCtx, &aggregator.MsgTransferOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, aggregator.ErrInvalidOwner.Error())

	// ACT: Attempt to transfer ownership to same address.
	_, err = server.TransferOwnership(goCtx, &aggregator.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: owner.Address,
	})
	// ASSERT: The action should've failed due to same address.
	require.ErrorContains(t, err, aggregator.ErrSameOwner.Error())

	// ARRANGE: Generate a new owner account.
	newOwner := utils.TestAccount()

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(goCtx, &aggregator.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've succeeded, and set owner in state.
	require.NoError(t, err)
	require.Equal(t, newOwner.Address, k.GetAggregatorOwner(ctx))
}
