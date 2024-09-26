package keeper_test

import (
	"testing"

	"github.com/noble-assets/halo/v2/utils"
	"github.com/noble-assets/halo/v2/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestGetNonces(t *testing.T) {
	keeper, ctx := mocks.HaloKeeper(t)

	// ACT: Retrieve all nonces with no state.
	nonces := keeper.GetNonces(ctx)
	// ASSERT: No nonces returned.
	require.Empty(t, nonces)

	// ARRANGE: Set nonces in state.
	user1, user2 := utils.TestAccount(), utils.TestAccount()
	err := keeper.SetNonce(ctx, user1.Bytes, 1)
	require.NoError(t, err)
	err = keeper.SetNonce(ctx, user2.Bytes, 2)
	require.NoError(t, err)

	// ACT: Retrieve all nonces.
	nonces = keeper.GetNonces(ctx)
	// ASSERT: Nonces returned.
	require.Len(t, nonces, 2)
	require.Equal(t, uint64(1), nonces[user1.Address])
	require.Equal(t, uint64(2), nonces[user2.Address])
}
