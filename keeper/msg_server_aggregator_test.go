// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	"github.com/noble-assets/halo/v2/keeper"
	"github.com/noble-assets/halo/v2/types"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/noble-assets/halo/v2/utils"
	"github.com/noble-assets/halo/v2/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestAggregatorTransferOwnership(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewAggregatorMsgServer(k)

	// ACT: Attempt to transfer ownership with no reporter set.
	_, err := server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{})
	// ASSERT: The action should've failed due to no reporter set.
	require.ErrorContains(t, err, "there is no reporter")

	// ARRANGE: Set reporter in state.
	reporter := utils.TestAccount()
	require.NoError(t, k.Reporter.Set(ctx, reporter.Address))

	// ACT: Attempt to transfer ownership with invalid signer.
	_, err = server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, aggregator.ErrInvalidReporter.Error())

	// ACT: Attempt to transfer ownership to same address.
	_, err = server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{
		Signer:      reporter.Address,
		NewReporter: reporter.Address,
	})
	// ASSERT: The action should've failed due to same address.
	require.ErrorContains(t, err, aggregator.ErrSameReporter.Error())

	// ARRANGE: Generate a new reporter account.
	newReporter := utils.TestAccount()

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmp := k.Reporter
	k.Reporter = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		aggregator.ReporterKey, "aggregator_reporter", collections.StringValue,
	)

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{
		Signer:      reporter.Address,
		NewReporter: newReporter.Address,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.Reporter = tmp

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(ctx, &aggregator.MsgTransferOwnership{
		Signer:      reporter.Address,
		NewReporter: newReporter.Address,
	})
	// ASSERT: The action should've succeeded, and set reporter in state.
	require.NoError(t, err)
	res, err := k.Reporter.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, newReporter.Address, res)
}
