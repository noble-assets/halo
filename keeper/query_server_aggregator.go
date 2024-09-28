// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/halo/v2/types/aggregator"
)

var _ aggregator.QueryServer = &aggregatorQueryServer{}

type aggregatorQueryServer struct {
	*Keeper
}

func NewAggregatorQueryServer(keeper *Keeper) aggregator.QueryServer {
	return &aggregatorQueryServer{Keeper: keeper}
}

func (k aggregatorQueryServer) Owner(ctx context.Context, req *aggregator.QueryOwner) (*aggregator.QueryOwnerResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	return &aggregator.QueryOwnerResponse{
		Owner: k.GetAggregatorOwner(ctx),
	}, nil
}

func (k aggregatorQueryServer) NextPrice(ctx context.Context, req *aggregator.QueryNextPrice) (*aggregator.QueryNextPriceResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	return &aggregator.QueryNextPriceResponse{
		NextPrice: k.GetNextPrice(ctx),
	}, nil
}

func (k aggregatorQueryServer) RoundData(ctx context.Context, req *aggregator.QueryRoundData) (*aggregator.QueryRoundDataResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	round, found := k.Keeper.GetRound(ctx, req.RoundId)
	if !found {
		return nil, fmt.Errorf("unknown round %d", req.RoundId)
	}

	return &aggregator.QueryRoundDataResponse{
		RoundId:         req.RoundId,
		Answer:          round.Answer,
		StartedAt:       round.UpdatedAt,
		UpdatedAt:       round.UpdatedAt,
		AnsweredInRound: req.RoundId,
	}, nil
}

func (k aggregatorQueryServer) LatestRoundData(ctx context.Context, req *aggregator.QueryLatestRoundData) (*aggregator.QueryRoundDataResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	id := k.GetLastRoundId(ctx)
	round, _ := k.GetRound(ctx, id)

	return &aggregator.QueryRoundDataResponse{
		RoundId:         id,
		Answer:          round.Answer,
		StartedAt:       round.UpdatedAt,
		UpdatedAt:       round.UpdatedAt,
		AnsweredInRound: id,
	}, nil
}

func (k aggregatorQueryServer) RoundDetails(ctx context.Context, req *aggregator.QueryRoundDetails) (*aggregator.QueryRoundDetailsResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	round, found := k.Keeper.GetRound(ctx, req.RoundId)
	if !found {
		return nil, fmt.Errorf("unknown round %d", req.RoundId)
	}

	return &aggregator.QueryRoundDetailsResponse{
		RoundId:     req.RoundId,
		Balance:     round.Balance,
		Interest:    round.Interest,
		TotalSupply: round.Supply,
		UpdatedAt:   round.UpdatedAt,
	}, nil
}

func (k aggregatorQueryServer) LatestRoundDetails(ctx context.Context, req *aggregator.QueryLatestRoundDetails) (*aggregator.QueryRoundDetailsResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	id := k.GetLastRoundId(ctx)
	round, _ := k.Keeper.GetRound(ctx, id)

	return &aggregator.QueryRoundDetailsResponse{
		RoundId:     id,
		Balance:     round.Balance,
		Interest:    round.Interest,
		TotalSupply: round.Supply,
		UpdatedAt:   round.UpdatedAt,
	}, nil
}
