// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/v3/keeper"
	"github.com/noble-assets/halo/v3/types"
	"github.com/noble-assets/halo/v3/types/aggregator"
	"github.com/noble-assets/halo/v3/utils"
	"github.com/noble-assets/halo/v3/utils/data"
	"github.com/noble-assets/halo/v3/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestReportBalance(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorMsgServer(k)

	// ACT: Attempt to report balance with no owner set.
	_, err := server.ReportBalance(ctx, &aggregator.MsgReportBalance{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set aggregator owner in state.
	owner := utils.TestAccount()
	err = k.SetAggregatorOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ARRANGE: Save the original LastRoundIDKey and reset it to an empty byte slice.
	tmpLastRoundIDKey := aggregator.LastRoundIDKey
	aggregator.LastRoundIDKey = []byte("")

	// ACT: Verify the LastRoundIDKey is set to zero when the key is reset to an empty slice.
	require.Equal(t, k.GetLastRoundId(ctx), uint64(0))

	// ARRANGE: Restore the original LastRoundIDKey to its previous value.
	aggregator.LastRoundIDKey = tmpLastRoundIDKey

	// ACT: Attempt to report balance with invalid signer.
	_, err = server.ReportBalance(ctx, &aggregator.MsgReportBalance{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, aggregator.ErrInvalidOwner.Error())

	//
	for _, testCase := range data.EthereumRounds {
		// ACT: Attempt to report balance.
		testCase.Msg.Signer = owner.Address
		res, err := server.ReportBalance(ctx, &testCase.Msg)

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
	_, err = server.ReportBalance(ctx, &msg)
	// ASSERT: The action should've failed due to identical round.
	require.ErrorContains(t, err, aggregator.ErrAlreadyReported.Error())

	// ACT: Attempt to report balance with a negative next price.
	msg = aggregator.MsgReportBalance{
		Signer:      owner.Address,
		Principal:   utils.MustParseInt("4600092531"),
		Interest:    utils.MustParseInt("658502"),
		TotalSupply: utils.MustParseInt("44285046691709"),
		NextPrice:   utils.MustParseInt("-1"),
	}
	_, err = server.ReportBalance(ctx, &msg)
	// ASSERT: The action should've failed due to invalid next price.
	require.ErrorContains(t, err, aggregator.ErrInvalidNextPrice.Error())

	// ACT: Attempt to report balance with invalid next price.
	// https://etherscan.io/tx/0xe8154863cf89175c9ec361999ec7ddeebf7c29297ee62325f777067409071303
	msg = aggregator.MsgReportBalance{
		Signer:      owner.Address,
		Principal:   utils.MustParseInt("4600092532"),
		Interest:    utils.MustParseInt("658503"),
		TotalSupply: utils.MustParseInt("44285046691710"),
		NextPrice:   math.ZeroInt(),
	}
	_, err = server.ReportBalance(ctx, &msg)
	// ASSERT: The action should've failed due to invalid next price.
	require.ErrorContains(t, err, aggregator.ErrInvalidNextPrice.Error())

	// ARRANGE: Set the next round in state.
	err = k.SetRound(ctx, k.GetLastRoundId(ctx)+1, aggregator.RoundData{})
	require.NoError(t, err)

	// ACT: Attempt to report balance with existing next round.
	// https://etherscan.io/tx/0x1a628856cb74de37357a35c29ec22509b72b1fc826ac7bb1020c73a99f9f80fc
	msg = aggregator.MsgReportBalance{
		Signer:      owner.Address,
		Principal:   utils.MustParseInt("4600751035"),
		Interest:    utils.MustParseInt("657348"),
		TotalSupply: utils.MustParseInt("44285680550257"),
		NextPrice:   utils.MustParseInt("103914726"),
	}
	_, err = server.ReportBalance(ctx, &msg)
	// ASSERT: The action should've failed due to existing round.
	require.ErrorContains(t, err, aggregator.ErrAlreadyReported.Error())

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmpRounds := k.Rounds
	k.Rounds = collections.NewMap(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		aggregator.RoundPrefix, "aggregator_rounds", collections.Uint64Key, codec.CollValue[aggregator.RoundData](
			mocks.MakeTestEncodingConfig("noble").Codec,
		),
	)

	// ACT: Attempt to report balance with failing Rounds store.
	_, err = server.ReportBalance(ctx, &msg)
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.Rounds = tmpRounds

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmpLastRoundId := k.LastRoundId
	k.LastRoundId = collections.NewSequence(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		aggregator.LastRoundIDKey, "aggregator_last_round_id",
	)

	// ACT: Attempt to report balance with failing LastRoundId store.
	_, err = server.ReportBalance(ctx, &msg)
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.LastRoundId = tmpLastRoundId

	// ARRANGE: Set up a failing collection store for the attribute setter.
	k.NextPrice = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		aggregator.NextPriceKey, "aggregator_next_price", sdk.IntValue,
	)

	// ACT: Attempt to report balance with failing NextPrice store.
	_, err = server.ReportBalance(ctx, &msg)
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
}

func TestSetNextPrice(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorMsgServer(k)

	// ACT: Attempt to set next price with no owner set.
	_, err := server.SetNextPrice(ctx, &aggregator.MsgSetNextPrice{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set aggregator owner in state.
	owner := utils.TestAccount()
	err = k.SetAggregatorOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt to set next price with invalid signer.
	_, err = server.SetNextPrice(ctx, &aggregator.MsgSetNextPrice{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, aggregator.ErrInvalidOwner.Error())

	// ACT: Attempt to set next price with invalid price.
	_, err = server.SetNextPrice(ctx, &aggregator.MsgSetNextPrice{
		Signer:    owner.Address,
		NextPrice: math.ZeroInt(),
	})
	// ASSERT: The action should've failed due to invalid price.
	require.ErrorContains(t, err, aggregator.ErrInvalidNextPrice.Error())

	// ACT: Attempt to set next price.
	// https://etherscan.io/tx/0xfd21979418ce5e6686c624841f48d11ed241b387b08eb60e2bd361de5ed1a061
	price := utils.MustParseInt("103780600")
	_, err = server.SetNextPrice(ctx, &aggregator.MsgSetNextPrice{
		Signer:    owner.Address,
		NextPrice: price,
	})
	// ASSERT: The action should've succeeded, and set next price in state.
	require.NoError(t, err)
	require.Equal(t, price, k.GetNextPrice(ctx))

	// ARRANGE: Set up a failing collection store for the attribute setter.
	k.NextPrice = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		aggregator.NextPriceKey, "aggregator_next_price", sdk.IntValue,
	)

	// ACT: Attempt to set next price.
	_, err = server.SetNextPrice(ctx, &aggregator.MsgSetNextPrice{
		Signer:    owner.Address,
		NextPrice: price,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
}

func TestAggregatorTransferOwnership(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorMsgServer(k)

	// ACT: Attempt to transfer ownership with no owner set.
	_, err := server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set aggregator owner in state.
	owner := utils.TestAccount()
	err = k.SetAggregatorOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt to transfer ownership with invalid signer.
	_, err = server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, aggregator.ErrInvalidOwner.Error())

	// ACT: Attempt to transfer ownership to same address.
	_, err = server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: owner.Address,
	})
	// ASSERT: The action should've failed due to same address.
	require.ErrorContains(t, err, aggregator.ErrSameOwner.Error())

	// ARRANGE: Generate a new owner account.
	newOwner := utils.TestAccount()

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmp := k.AggregatorOwner
	k.AggregatorOwner = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		aggregator.OwnerKey, "aggregator_owner", collections.StringValue,
	)

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.AggregatorOwner = tmp

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've succeeded, and set owner in state.
	require.NoError(t, err)
	require.Equal(t, newOwner.Address, k.GetAggregatorOwner(ctx))
}
