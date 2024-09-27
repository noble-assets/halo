package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/v2/keeper"
	"github.com/noble-assets/halo/v2/types"
	"github.com/noble-assets/halo/v2/types/entitlements"
	"github.com/noble-assets/halo/v2/utils"
	"github.com/noble-assets/halo/v2/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestSendRestrictionBurn(t *testing.T) {
	user := utils.TestAccount()
	k, ctx := mocks.HaloKeeper(t)
	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, ONE))

	testCases := []struct {
		name    string
		paused  bool
		allowed bool
		err     string
	}{
		{
			name:    "PausedAndAllowed",
			paused:  true,
			allowed: true,
			err:     "cannot transfer",
		},
		{
			name:    "PausedAndNotAllowed",
			paused:  true,
			allowed: false,
			err:     "cannot transfer",
		},
		{
			name:    "UnpausedAndAllowed",
			paused:  false,
			allowed: true,
		},
		{
			name:    "UnpausedAndNotAllowed",
			paused:  false,
			allowed: false,
			err:     "cannot transfer",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// ARRANGE: Set paused state.
			err := k.SetPaused(ctx, testCase.paused)
			require.NoError(t, err)
			// ARRANGE: Set allowed state.
			err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, testCase.allowed)
			require.NoError(t, err)

			// ACT: Attempt to burn.
			_, err = k.SendRestrictionFn(ctx, user.Bytes, types.ModuleAddress, coins)

			// ASSERT: Send restriction correctly handled test case.
			if testCase.err != "" {
				require.ErrorContains(t, err, testCase.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSendRestrictionMint(t *testing.T) {
	user := utils.TestAccount()
	k, ctx := mocks.HaloKeeper(t)
	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, ONE))

	testCases := []struct {
		name    string
		paused  bool
		allowed bool
		err     string
	}{
		{
			name:    "PausedAndAllowed",
			paused:  true,
			allowed: true,
			err:     "cannot transfer",
		},
		{
			name:    "PausedAndNotAllowed",
			paused:  true,
			allowed: false,
			err:     "cannot transfer",
		},
		{
			name:    "UnpausedAndAllowed",
			paused:  false,
			allowed: true,
		},
		{
			name:    "UnpausedAndNotAllowed",
			paused:  false,
			allowed: false,
			err:     "cannot transfer",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// ARRANGE: Set paused state.
			err := k.SetPaused(ctx, testCase.paused)
			require.NoError(t, err)
			// ARRANGE: Set allowed state.
			err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, testCase.allowed)
			require.NoError(t, err)

			// ACT: Attempt to mint.
			_, err = k.SendRestrictionFn(ctx, types.ModuleAddress, user.Bytes, coins)

			// ASSERT: Send restriction correctly handled test case.
			if testCase.err != "" {
				require.ErrorContains(t, err, testCase.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSendRestrictionTransfer(t *testing.T) {
	alice, bob := utils.TestAccount(), utils.TestAccount()
	k, ctx := mocks.HaloKeeper(t)
	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, ONE))

	testCases := []struct {
		name             string
		paused           bool
		senderAllowed    bool
		recipientAllowed bool
		coins            sdk.Coins
		err              string
	}{
		{
			name:             "NonUSYCTransfer",
			paused:           true,
			senderAllowed:    true,
			recipientAllowed: true,
			coins:            sdk.NewCoins(sdk.NewCoin("uusdc", math.NewInt(1_000_000))),
			err:              "",
		},
		{
			name:             "PausedAndSenderAllowedAndRecipientAllowed",
			paused:           true,
			senderAllowed:    true,
			recipientAllowed: true,
			coins:            coins,
			err:              "cannot transfer",
		},
		{
			name:             "PausedAndSenderAllowedAndRecipientNotAllowed",
			paused:           true,
			senderAllowed:    true,
			recipientAllowed: false,
			coins:            coins,
			err:              "cannot transfer",
		},
		{
			name:             "PausedAndSenderNotAllowedAndRecipientAllowed",
			paused:           true,
			senderAllowed:    false,
			recipientAllowed: true,
			coins:            coins,
			err:              "cannot transfer",
		},
		{
			name:             "PausedAndSenderNotAllowedAndRecipientNotAllowed",
			paused:           true,
			senderAllowed:    false,
			recipientAllowed: false,
			coins:            coins,
			err:              "cannot transfer",
		},
		{
			name:             "UnpausedAndSenderAllowedAndRecipientAllowed",
			paused:           false,
			senderAllowed:    true,
			recipientAllowed: true,
			coins:            coins,
			err:              "",
		},
		{
			name:             "UnpausedAndSenderAllowedAndRecipientNotAllowed",
			paused:           false,
			senderAllowed:    true,
			recipientAllowed: false,
			coins:            coins,
			err:              "cannot transfer",
		},
		{
			name:             "UnpausedAndSenderNotAllowedAndRecipientAllowed",
			paused:           false,
			senderAllowed:    false,
			recipientAllowed: true,
			coins:            coins,
			err:              "cannot transfer",
		},
		{
			name:             "UnpausedAndSenderNotAllowedAndRecipientNotAllowed",
			paused:           false,
			senderAllowed:    false,
			recipientAllowed: false,
			coins:            coins,
			err:              "cannot transfer",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// ARRANGE: Set paused state.
			err := k.SetPaused(ctx, testCase.paused)
			require.NoError(t, err)
			// ARRANGE: Set sender allowed state.
			err = k.SetUserRole(ctx, alice.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, testCase.senderAllowed)
			require.NoError(t, err)
			// ARRANGE: Set recipient allowed state.
			err = k.SetUserRole(ctx, bob.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, testCase.recipientAllowed)
			require.NoError(t, err)

			// ACT: Attempt to transfer.
			_, err = k.SendRestrictionFn(ctx, alice.Bytes, bob.Bytes, testCase.coins)

			// ASSERT: Send restriction correctly handled test case.
			if testCase.err != "" {
				require.ErrorContains(t, err, testCase.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNewKeeper(t *testing.T) {
	// ARRANGE: Set the NoncePrefix to an already existing key
	types.NoncePrefix = types.OwnerKey

	// ACT: Attempt to create a new Keeper with overlapping prefixes
	require.Panics(t, func() {
		cfg := mocks.MakeTestEncodingConfig("noble")
		keeper.NewKeeper(cfg.Codec, mocks.FailingStore(mocks.Set, nil), "uusyc", "uusdc", mocks.AccountKeeper{}, mocks.BankKeeper{}, cfg.InterfaceRegistry)
	})
	// ASSERT: The function should've panicked.

	// ARRANGE: Restore the original NoncePrefix
	types.NoncePrefix = []byte("nonce/")
}
