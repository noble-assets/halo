package mocks

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
)

func HaloKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	return HaloKeeperWithKeepers(
		t,
		AccountKeeper{},
		BankKeeper{
			Restriction: NoOpSendRestrictionFn,
		},
	)
}

func HaloKeeperWithKeepers(_ testing.TB, account AccountKeeper, bank BankKeeper) (*keeper.Keeper, sdk.Context) {
	key := storetypes.NewKVStoreKey(types.ModuleName)
	tkey := storetypes.NewTransientStoreKey("transient_halo")
	ctx := testutil.DefaultContext(key, tkey)

	registry := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	k := keeper.NewKeeper(
		cdc,
		key,
		"uusyc",
		"uusdc",
		account,
		nil,
	)

	bank = bank.WithSendCoinsRestriction(k.SendRestrictionFn)
	k.SetBankKeeper(bank)

	for i := 1; i <= int(entitlements.ROLE_INTERNATIONAL_SDYF); i++ {
		role := entitlements.Role(i)

		k.SetRoleCapability(ctx, sdk.MsgTypeURL(&types.MsgBurn{}), role, true)
		k.SetRoleCapability(ctx, sdk.MsgTypeURL(&types.MsgDeposit{}), role, true)
		k.SetRoleCapability(ctx, sdk.MsgTypeURL(&types.MsgDepositFor{}), role, true)
		k.SetRoleCapability(ctx, sdk.MsgTypeURL(&types.MsgWithdraw{}), role, true)
		k.SetRoleCapability(ctx, sdk.MsgTypeURL(&types.MsgWithdrawTo{}), role, true)
		k.SetRoleCapability(ctx, "transfer", role, true)
	}

	return k, ctx
}
