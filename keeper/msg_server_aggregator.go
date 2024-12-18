// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"github.com/noble-assets/halo/v3/types/aggregator"
)

var _ aggregator.MsgServer = &aggregatorMsgServer{}

type aggregatorMsgServer struct {
	*Keeper
}

func NewAggregatorMsgServer(keeper *Keeper) aggregator.MsgServer {
	return &aggregatorMsgServer{Keeper: keeper}
}

func (k aggregatorMsgServer) ReportBalance(ctx context.Context, msg *aggregator.MsgReportBalance) (*aggregator.MsgReportBalanceResponse, error) {
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	balance := msg.Principal.Add(msg.Interest)

	id, err := k.IncrementLastRoundId(ctx)
	if err != nil {
		return nil, err
	}
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
		UpdatedAt: k.headerService.GetHeaderInfo(ctx).Time.Unix(),
	}
	if err = k.SetRound(ctx, id, round); err != nil {
		return nil, err
	}

	if !msg.NextPrice.IsPositive() || msg.NextPrice.LTE(answer) {
		return nil, aggregator.ErrInvalidNextPrice
	}
	if err = k.Keeper.SetNextPrice(ctx, msg.NextPrice); err != nil {
		return nil, err
	}

	_ = k.eventService.EventManager(ctx).Emit(ctx, &aggregator.BalanceReported{
		RoundId:   id,
		Balance:   balance,
		Interest:  msg.Interest,
		Price:     answer,
		UpdatedAt: k.headerService.GetHeaderInfo(ctx).Time.Unix(),
	})
	_ = k.eventService.EventManager(ctx).Emit(ctx, &aggregator.NextPriceReported{Price: msg.NextPrice})

	return &aggregator.MsgReportBalanceResponse{
		RoundId: id,
	}, nil
}

func (k aggregatorMsgServer) SetNextPrice(ctx context.Context, msg *aggregator.MsgSetNextPrice) (*aggregator.MsgSetNextPriceResponse, error) {
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if !msg.NextPrice.IsPositive() {
		return nil, aggregator.ErrInvalidNextPrice
	}

	if err = k.Keeper.SetNextPrice(ctx, msg.NextPrice); err != nil {
		return nil, err
	}

	return &aggregator.MsgSetNextPriceResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &aggregator.NextPriceReported{
		Price: msg.NextPrice,
	})
}

func (k aggregatorMsgServer) TransferOwnership(ctx context.Context, msg *aggregator.MsgTransferOwnership) (*aggregator.MsgTransferOwnershipResponse, error) {
	owner, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if msg.NewOwner == owner {
		return nil, aggregator.ErrSameOwner
	}

	if err = k.SetAggregatorOwner(ctx, msg.NewOwner); err != nil {
		return nil, err
	}

	return &aggregator.MsgTransferOwnershipResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &aggregator.OwnershipTransferred{
		PreviousOwner: owner,
		NewOwner:      msg.NewOwner,
	})
}

//

func (k aggregatorMsgServer) EnsureOwner(ctx context.Context, signer string) (string, error) {
	owner := k.GetAggregatorOwner(ctx)
	if owner == "" {
		return "", aggregator.ErrNoOwner
	}
	if signer != owner {
		return "", errors.Wrapf(aggregator.ErrInvalidOwner, "expected %s, got %s", owner, signer)
	}
	return owner, nil
}
