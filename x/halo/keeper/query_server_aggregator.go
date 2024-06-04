package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
)

var _ aggregator.QueryServer = &aggregatorQueryServer{}

type aggregatorQueryServer struct {
	*Keeper
}

func NewAggregatorQueryServer(keeper *Keeper) aggregator.QueryServer {
	return &aggregatorQueryServer{Keeper: keeper}
}

func (k aggregatorQueryServer) Owner(goCtx context.Context, req *aggregator.QueryOwner) (*aggregator.QueryOwnerResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &aggregator.QueryOwnerResponse{
		Owner: k.GetAggregatorOwner(ctx),
	}, nil
}

func (k aggregatorQueryServer) NextPrice(goCtx context.Context, req *aggregator.QueryNextPrice) (*aggregator.QueryNextPriceResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &aggregator.QueryNextPriceResponse{
		NextPrice: k.GetNextPrice(ctx),
	}, nil
}

func (k aggregatorQueryServer) RoundData(goCtx context.Context, req *aggregator.QueryRoundData) (*aggregator.QueryRoundDataResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
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

func (k aggregatorQueryServer) LatestRoundData(goCtx context.Context, req *aggregator.QueryLatestRoundData) (*aggregator.QueryRoundDataResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
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

func (k aggregatorQueryServer) RoundDetails(goCtx context.Context, req *aggregator.QueryRoundDetails) (*aggregator.QueryRoundDetailsResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
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

func (k aggregatorQueryServer) LatestRoundDetails(goCtx context.Context, req *aggregator.QueryLatestRoundDetails) (*aggregator.QueryRoundDetailsResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
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
