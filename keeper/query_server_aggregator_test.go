// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/halo/v2/keeper"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/noble-assets/halo/v2/utils"
	"github.com/noble-assets/halo/v2/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestReporterQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorQueryServer(k)

	// ACT: Attempt to reporter with invalid request.
	_, err := server.Reporter(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set reporter in state.
	reporter := utils.TestAccount()
	require.NoError(t, k.Reporter.Set(ctx, reporter.Address))

	// ACT: Attempt to query reporter.
	res, err := server.Reporter(ctx, &aggregator.QueryReporter{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Equal(t, reporter.Address, res.Reporter)
}
