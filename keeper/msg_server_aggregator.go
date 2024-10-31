// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"github.com/noble-assets/halo/v2/types/aggregator"
)

var _ aggregator.MsgServer = &aggregatorMsgServer{}

type aggregatorMsgServer struct {
	*Keeper
}

func NewAggregatorMsgServer(keeper *Keeper) aggregator.MsgServer {
	return &aggregatorMsgServer{Keeper: keeper}
}

func (k aggregatorMsgServer) Transmit(ctx context.Context, msg *aggregator.MsgTransmit) (*aggregator.MsgTransmitResponse, error) {
	_, err := k.EnsureReporter(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	id, err := k.LastRoundId.Next(ctx)
	if err != nil {
		return nil, err
	}

	round := aggregator.RoundData{
		Answer:    msg.Answer,
		UpdatedAt: msg.UpdatedAt,
	}

	if err = k.Rounds.Set(ctx, id, round); err != nil {
		return nil, err
	}

	return &aggregator.MsgTransmitResponse{RoundId: id}, k.eventService.EventManager(ctx).Emit(ctx, &aggregator.AnswerUpdated{
		Current:   msg.Answer,
		RoundId:   id,
		UpdatedAt: msg.UpdatedAt,
	})
}

func (k aggregatorMsgServer) SetNextPrice(ctx context.Context, msg *aggregator.MsgSetNextPrice) (*aggregator.MsgSetNextPriceResponse, error) {
	_, err := k.EnsureReporter(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if msg.NextPrice.IsZero() {
		return nil, aggregator.ErrInvalidNextPrice
	}

	id, err := k.LastRoundId.Peek(ctx)
	if err != nil {
		return nil, err
	}

	if err = k.NextPrices.Set(ctx, id, msg.NextPrice); err != nil {
		return nil, err
	}

	return &aggregator.MsgSetNextPriceResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &aggregator.NextPriceReported{
		RoundId:   id,
		NextPrice: msg.NextPrice,
	})
}

func (k aggregatorMsgServer) TransferOwnership(ctx context.Context, msg *aggregator.MsgTransferOwnership) (*aggregator.MsgTransferOwnershipResponse, error) {
	reporter, err := k.EnsureReporter(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if msg.NewReporter == reporter {
		return nil, aggregator.ErrSameReporter
	}

	if err = k.Reporter.Set(ctx, msg.NewReporter); err != nil {
		return nil, err
	}

	return &aggregator.MsgTransferOwnershipResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &aggregator.OwnershipTransferred{
		PreviousReporter: reporter,
		NewReporter:      msg.NewReporter,
	})
}

//

func (k aggregatorMsgServer) EnsureReporter(ctx context.Context, signer string) (string, error) {
	reporter, _ := k.Reporter.Get(ctx)
	if reporter == "" {
		return "", aggregator.ErrNoReporter
	}
	if signer != reporter {
		return "", errors.Wrapf(aggregator.ErrInvalidReporter, "expected %s, got %s", reporter, signer)
	}
	return reporter, nil
}
