package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/utils"
	"github.com/noble-assets/halo/utils/mocks"
	"github.com/noble-assets/halo/x/halo/types"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
	"github.com/stretchr/testify/require"
)

func TestSendRestrictionBurn(t *testing.T) {
	user := utils.TestAccount()
	keeper, ctx := mocks.HaloKeeper(t)
	coins := sdk.NewCoins(sdk.NewCoin(keeper.Denom, ONE))

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
			keeper.SetPaused(ctx, testCase.paused)
			// ARRANGE: Set allowed state.
			keeper.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, testCase.allowed)

			// ACT: Attempt to burn.
			_, err := keeper.SendRestrictionFn(ctx, user.Bytes, types.ModuleAddress, coins)

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
	keeper, ctx := mocks.HaloKeeper(t)
	coins := sdk.NewCoins(sdk.NewCoin(keeper.Denom, ONE))

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
			keeper.SetPaused(ctx, testCase.paused)
			// ARRANGE: Set allowed state.
			keeper.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, testCase.allowed)

			// ACT: Attempt to mint.
			_, err := keeper.SendRestrictionFn(ctx, types.ModuleAddress, user.Bytes, coins)

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
	keeper, ctx := mocks.HaloKeeper(t)
	coins := sdk.NewCoins(sdk.NewCoin(keeper.Denom, ONE))

	testCases := []struct {
		name             string
		paused           bool
		ftfPaused        bool
		ftfBlacklist     [][]byte
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
			coins:            sdk.NewCoins(sdk.NewCoin("uusdc", sdk.NewInt(1_000_000))),
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
		{
			name:             "FtfPaused",
			paused:           false,
			senderAllowed:    true,
			recipientAllowed: true,
			coins:            sdk.NewCoins(sdk.NewCoin("uusdc", ONE)),
			ftfPaused:        true,
			ftfBlacklist:     make([][]byte, 0),
			err:              "transfers are paused",
		},
		{
			name:             "FtfSenderBlacklisted",
			paused:           false,
			senderAllowed:    true,
			recipientAllowed: true,
			coins:            sdk.NewCoins(sdk.NewCoin("uusdc", ONE)),
			ftfPaused:        false,
			ftfBlacklist:     [][]byte{alice.Bytes},
			err:              "is blocked from sending",
		},
		{
			name:             "FtfReceiverBlacklisted",
			paused:           false,
			senderAllowed:    true,
			recipientAllowed: true,
			coins:            sdk.NewCoins(sdk.NewCoin("uusdc", ONE)),
			ftfPaused:        false,
			ftfBlacklist:     [][]byte{bob.Bytes},
			err:              "is blocked from receiving",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// ARRANGE: Set paused state.
			keeper.SetPaused(ctx, testCase.paused)
			// ARRANGE: Set sender allowed state.
			keeper.SetUserRole(ctx, alice.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, testCase.senderAllowed)
			// ARRANGE: Set recipient allowed state.
			keeper.SetUserRole(ctx, bob.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, testCase.recipientAllowed)
			// ARRANGE: Set FiatTokenFactory Keeper
			keeper.SetFTFKeeper(mocks.FTFKeeper{Paused: testCase.ftfPaused, Blacklist: testCase.ftfBlacklist})

			// ACT: Attempt to transfer.
			_, err := keeper.SendRestrictionFn(ctx, alice.Bytes, bob.Bytes, testCase.coins)

			// ASSERT: Send restriction correctly handled test case.
			if testCase.err != "" {
				require.ErrorContains(t, err, testCase.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
