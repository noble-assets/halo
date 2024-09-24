package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/halo/v2/keeper"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/noble-assets/halo/v2/utils"
	"github.com/noble-assets/halo/v2/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestAggregatorOwnerQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorQueryServer(k)

	// ACT: Attempt to query aggregator owner with invalid request.
	_, err := server.Owner(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set aggregator owner in state.
	owner := utils.TestAccount()
	k.SetAggregatorOwner(ctx, owner.Address)

	// ACT: Attempt to query aggregator owner.
	res, err := server.Owner(ctx, &aggregator.QueryOwner{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Equal(t, owner.Address, res.Owner)
}

func TestNextPriceQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorQueryServer(k)

	// ACT: Attempt to query next price with invalid request.
	_, err := server.NextPrice(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query next price with no state.
	res, err := server.NextPrice(ctx, &aggregator.QueryNextPrice{})
	// ASSERT: The query should've succeeded, and returned nothing.
	require.NoError(t, err)
	require.True(t, res.NextPrice.IsZero())

	// ARRANGE: Set next price in state.
	// https://etherscan.io/tx/0xfd21979418ce5e6686c624841f48d11ed241b387b08eb60e2bd361de5ed1a061
	expected := math.NewInt(103780600)
	k.SetNextPrice(ctx, expected)

	// ACT: Attempt to query next price with state.
	res, err = server.NextPrice(ctx, &aggregator.QueryNextPrice{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Equal(t, expected, res.NextPrice)
}

func TestRoundDataQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorQueryServer(k)

	// ACT: Attempt to query round data with invalid request.
	_, err := server.RoundData(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set a round in state.
	// https://etherscan.io/tx/0xbcba4db502a72e51a05b378d2e5867be4c60936585cedaf5aad90002f0599428
	k.SetRound(ctx, 187, aggregator.RoundData{
		Answer:    math.NewInt(103780685),
		Balance:   math.NewInt(4791541000),
		Interest:  math.NewInt(701123),
		Supply:    math.NewInt(46169872257060),
		UpdatedAt: 1712071487,
	})

	// ACT: Attempt to query round data of unknown round.
	_, err = server.RoundData(ctx, &aggregator.QueryRoundData{RoundId: 0})
	// ASSERT: The query should've failed due to unknown round.
	require.ErrorContains(t, err, "unknown round")

	// ACT: Attempt to query round data.
	res, err := server.RoundData(ctx, &aggregator.QueryRoundData{RoundId: 187})
	// ASSERT: The query should've successfully returned round data.
	require.NoError(t, err)
	require.Equal(t, uint64(187), res.RoundId)
	require.Equal(t, math.NewInt(103780685), res.Answer)
	require.Equal(t, int64(1712071487), res.StartedAt)
	require.Equal(t, int64(1712071487), res.UpdatedAt)
	require.Equal(t, uint64(187), res.AnsweredInRound)
}

func TestLatestRoundDataQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorQueryServer(k)

	// ACT: Attempt to query latest round data with invalid request.
	_, err := server.LatestRoundData(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set a round in state.
	// https://etherscan.io/tx/0xbcba4db502a72e51a05b378d2e5867be4c60936585cedaf5aad90002f0599428
	k.SetRound(ctx, 187, aggregator.RoundData{
		Answer:    math.NewInt(103780685),
		Balance:   math.NewInt(4791541000),
		Interest:  math.NewInt(701123),
		Supply:    math.NewInt(46169872257060),
		UpdatedAt: 1712071487,
	})
	k.SetLastRoundId(ctx, 187)

	// ACT: Attempt to query latest round details.
	res, err := server.LatestRoundData(ctx, &aggregator.QueryLatestRoundData{})
	// ASSERT: The query should've successfully returned round data.
	require.NoError(t, err)
	require.Equal(t, uint64(187), res.RoundId)
	require.Equal(t, math.NewInt(103780685), res.Answer)
	require.Equal(t, int64(1712071487), res.StartedAt)
	require.Equal(t, int64(1712071487), res.UpdatedAt)
	require.Equal(t, uint64(187), res.AnsweredInRound)
}

func TestRoundDetailsQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorQueryServer(k)

	// ACT: Attempt to query round details with invalid request.
	_, err := server.RoundDetails(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set a round in state.
	// https://etherscan.io/tx/0xbcba4db502a72e51a05b378d2e5867be4c60936585cedaf5aad90002f0599428
	k.SetRound(ctx, 187, aggregator.RoundData{
		Answer:    math.NewInt(103780685),
		Balance:   math.NewInt(4791541000),
		Interest:  math.NewInt(701123),
		Supply:    math.NewInt(46169872257060),
		UpdatedAt: 1712071487,
	})

	// ACT: Attempt to query round details of unknown round.
	_, err = server.RoundDetails(ctx, &aggregator.QueryRoundDetails{RoundId: 0})
	// ASSERT: The query should've failed due to unknown round.
	require.ErrorContains(t, err, "unknown round")

	// ACT: Attempt to query round details.
	res, err := server.RoundDetails(ctx, &aggregator.QueryRoundDetails{RoundId: 187})
	// ASSERT: The query should've successfully returned round details.
	require.NoError(t, err)
	require.Equal(t, uint64(187), res.RoundId)
	require.Equal(t, math.NewInt(4791541000), res.Balance)
	require.Equal(t, math.NewInt(701123), res.Interest)
	require.Equal(t, math.NewInt(46169872257060), res.TotalSupply)
	require.Equal(t, int64(1712071487), res.UpdatedAt)
}

func TestLatestRoundDetailsQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorQueryServer(k)

	// ACT: Attempt to query latest round details with invalid request.
	_, err := server.LatestRoundDetails(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set a round in state.
	// https://etherscan.io/tx/0xbcba4db502a72e51a05b378d2e5867be4c60936585cedaf5aad90002f0599428
	k.SetRound(ctx, 187, aggregator.RoundData{
		Answer:    math.NewInt(103780685),
		Balance:   math.NewInt(4791541000),
		Interest:  math.NewInt(701123),
		Supply:    math.NewInt(46169872257060),
		UpdatedAt: 1712071487,
	})
	k.SetLastRoundId(ctx, 187)

	// ACT: Attempt to query latest round details.
	res, err := server.LatestRoundDetails(ctx, &aggregator.QueryLatestRoundDetails{})
	// ASSERT: The query should've successfully returned round details.
	require.NoError(t, err)
	require.Equal(t, uint64(187), res.RoundId)
	require.Equal(t, math.NewInt(4791541000), res.Balance)
	require.Equal(t, math.NewInt(701123), res.Interest)
	require.Equal(t, math.NewInt(46169872257060), res.TotalSupply)
	require.Equal(t, int64(1712071487), res.UpdatedAt)
}
