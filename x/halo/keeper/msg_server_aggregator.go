package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
)

var _ aggregator.MsgServer = &aggregatorMsgServer{}

type aggregatorMsgServer struct {
	*Keeper
}

func NewAggregatorMsgServer(keeper *Keeper) aggregator.MsgServer {
	return &aggregatorMsgServer{Keeper: keeper}
}

func (k aggregatorMsgServer) ReportBalance(goCtx context.Context, msg *aggregator.MsgReportBalance) (*aggregator.MsgReportBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	balance := msg.Principal.Add(msg.Interest)

	id := k.IncrementLastRoundId(ctx)
	round, found := k.GetRound(ctx, id)
	if found && round.Balance.Equal(balance) && round.Interest.Equal(msg.Interest) && round.Supply.Equal(msg.TotalSupply) {
		return nil, aggregator.ErrAlreadyReported
	}

	id += 1
	if _, found := k.GetRound(ctx, id); found {
		return nil, aggregator.ErrAlreadyReported
	}

	answer := balance.MulRaw(1_000_000_000_000).Quo(msg.TotalSupply)
	round = aggregator.RoundData{
		Answer:    answer,
		Balance:   balance,
		Interest:  msg.Interest,
		Supply:    msg.TotalSupply,
		UpdatedAt: ctx.BlockTime().Unix(),
	}
	k.SetRound(ctx, id, round)

	if !msg.NextPrice.IsPositive() || msg.NextPrice.LTE(answer) {
		return nil, aggregator.ErrInvalidNextPrice
	}
	k.Keeper.SetNextPrice(ctx, msg.NextPrice)

	return &aggregator.MsgReportBalanceResponse{
			RoundId: id,
		}, ctx.EventManager().EmitTypedEvents(
			&aggregator.BalanceReported{
				RoundId:   id,
				Balance:   balance,
				Interest:  msg.Interest,
				Price:     answer,
				UpdatedAt: ctx.BlockTime().Unix(),
			},
			&aggregator.NextPriceReported{Price: msg.NextPrice},
		)
}

func (k aggregatorMsgServer) SetNextPrice(goCtx context.Context, msg *aggregator.MsgSetNextPrice) (*aggregator.MsgSetNextPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if !msg.NextPrice.IsPositive() {
		return nil, aggregator.ErrInvalidNextPrice
	}

	k.Keeper.SetNextPrice(ctx, msg.NextPrice)

	return &aggregator.MsgSetNextPriceResponse{}, ctx.EventManager().EmitTypedEvent(&aggregator.NextPriceReported{
		Price: msg.NextPrice,
	})
}

func (k aggregatorMsgServer) TransferOwnership(goCtx context.Context, msg *aggregator.MsgTransferOwnership) (*aggregator.MsgTransferOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	owner, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if msg.NewOwner == owner {
		return nil, aggregator.ErrSameOwner
	}

	k.SetAggregatorOwner(ctx, msg.NewOwner)

	return &aggregator.MsgTransferOwnershipResponse{}, ctx.EventManager().EmitTypedEvent(&aggregator.OwnershipTransferred{
		PreviousOwner: owner,
		NewOwner:      msg.NewOwner,
	})
}

//

func (k aggregatorMsgServer) EnsureOwner(ctx sdk.Context, signer string) (string, error) {
	owner := k.GetAggregatorOwner(ctx)
	if owner == "" {
		return "", aggregator.ErrNoOwner
	}
	if signer != owner {
		return "", errors.Wrapf(aggregator.ErrInvalidOwner, "expected %s, got %s", owner, signer)
	}
	return owner, nil
}
