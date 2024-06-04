package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/halo/utils"
	"github.com/noble-assets/halo/utils/mocks"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types"
	"github.com/stretchr/testify/require"
)

func TestOwnerQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewQueryServer(k)

	// ACT: Attempt to query owner with invalid request.
	_, err := server.Owner(goCtx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set owner in state.
	owner := utils.TestAccount()
	k.SetOwner(ctx, owner.Address)

	// ACT: Attempt to query owner.
	res, err := server.Owner(goCtx, &types.QueryOwner{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Equal(t, owner.Address, res.Owner)
}

func TestNoncesQuery(t *testing.T) {
	// NOTE: Query pagination is assumed working, so isn't testing here.

	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewQueryServer(k)

	// ACT: Attempt to query nonces with invalid request.
	_, err := server.Nonces(goCtx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query nonces with no state.
	res, err := server.Nonces(goCtx, &types.QueryNonces{})
	// ASSERT: The query should've succeeded, with empty nonces.
	require.NoError(t, err)
	require.Empty(t, res.Nonces)

	// ARRANGE: Set nonces in state.
	user1, user2 := utils.TestAccount(), utils.TestAccount()
	k.SetNonce(ctx, user1.Bytes, 1)
	k.SetNonce(ctx, user2.Bytes, 2)

	// ACT: Attempt to query nonces with state.
	res, err = server.Nonces(goCtx, &types.QueryNonces{})
	// ASSERT: The query should've succeeded, with nonces.
	require.NoError(t, err)
	require.Len(t, res.Nonces, 2)
	require.Equal(t, uint64(1), res.Nonces[user1.Address])
	require.Equal(t, uint64(2), res.Nonces[user2.Address])
}

func TestNonceQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewQueryServer(k)

	// ACT: Attempt to query nonce with invalid request.
	_, err := server.Nonce(goCtx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query nonce with invalid address.
	_, err = server.Nonce(goCtx, &types.QueryNonce{
		Address: "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn",
	})
	// ASSERT: The query should've failed due to invalid address.
	require.ErrorContains(t, err, "unable to decode address")

	// ACT: Attempt to query nonce of unused account.
	res, err := server.Nonce(goCtx, &types.QueryNonce{
		Address: utils.TestAccount().Address,
	})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Zero(t, res.Nonce)

	// ARRANGE: Set nonce in state.
	user := utils.TestAccount()
	k.SetNonce(ctx, user.Bytes, 1)

	// ACT: Attempt to query nonce of used account.
	res, err = server.Nonce(goCtx, &types.QueryNonce{
		Address: user.Address,
	})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Equal(t, uint64(1), res.Nonce)
}
