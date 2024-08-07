package keeper_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/noble-assets/halo/utils"
	"github.com/noble-assets/halo/utils/mocks"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
	"github.com/stretchr/testify/require"
)

var ONE = sdk.NewInt(1_000_000)

func TestDeposit(t *testing.T) {
	// This test is based off of a real action on Ethereum.
	// https://etherscan.io/tx/0x9fc3dc5ec218eea1e84bed8f2a0f2aaf243b355c560ddd30820cf3ea0358fec0
	amount, expected := sdk.NewInt(202400000), sdk.NewInt(193549970)

	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	ftf := mocks.FTFKeeper{
		Paused: false,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank, ftf)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate a user account.
	user := utils.TestAccount()

	// ACT: Attempt to deposit with an invalid signer address.
	_, err := server.Deposit(goCtx, &types.MsgDeposit{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to deposit without required permissions.
	_, err = server.Deposit(goCtx, &types.MsgDeposit{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgDeposit")

	// ARRANGE: Assign the international feeder role to user.
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	// ARRANGE: Report Ethereum Round #229.
	// https://etherscan.io/tx/0xcff68ffc6f79afadf835f559f8a51ed7092bc679d2a4f34cd153ef321d6bc8ec
	k.SetRound(ctx, 229, aggregator.RoundData{
		Answer:    sdk.NewInt(104572478),
		Balance:   sdk.NewInt(7016169453),
		Interest:  sdk.NewInt(1005815),
		Supply:    sdk.NewInt(67093843285741),
		UpdatedAt: 1717153499,
	})
	k.SetLastRoundId(ctx, 229)

	// ACT: Attempt to deposit with insufficient funds.
	_, err = server.Deposit(goCtx, &types.MsgDeposit{
		Signer: user.Address,
		Amount: amount,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Give user 202.40 $USDC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Underlying, amount))

	// ACT: Attempt to deposit.
	_, err = server.Deposit(goCtx, &types.MsgDeposit{
		Signer: user.Address,
		Amount: amount,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.Equal(t, expected, bank.Balances[user.Address].AmountOf(k.Denom))
	require.True(t, bank.Balances[user.Address].AmountOf(k.Underlying).IsZero())
}

func TestDepositFor(t *testing.T) {
	// This test is based off of a real action on Ethereum.
	// https://etherscan.io/tx/0x9fc3dc5ec218eea1e84bed8f2a0f2aaf243b355c560ddd30820cf3ea0358fec0
	amount, expected := sdk.NewInt(202400000), sdk.NewInt(193549970)

	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	ftf := mocks.FTFKeeper{
		Paused: false,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank, ftf)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate user and recipient accounts.
	user, recipient := utils.TestAccount(), utils.TestAccount()

	// ACT: Attempt to deposit for with an invalid signer address.
	_, err := server.DepositFor(goCtx, &types.MsgDepositFor{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to deposit for without required permissions.
	_, err = server.DepositFor(goCtx, &types.MsgDepositFor{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgDepositFor")

	// ARRANGE: Assign the international feeder role to user.
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to deposit for with an invalid recipient address.
	_, err = server.DepositFor(goCtx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Invalid,
	})
	// ASSERT: The action should've failed due to invalid recipient address.
	require.ErrorContains(t, err, "unable to decode recipient address")

	// ACT: Attempt to deposit for without required recipient permissions.
	_, err = server.DepositFor(goCtx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Address,
	})
	// ASSERT: The action should've failed due to invalid recipient permissions.
	require.ErrorContains(t, err, "cannot receive uusyc")

	// ARRANGE: Assign the international feeder role to recipient.
	k.SetUserRole(ctx, recipient.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to deposit for with non-existing round.
	_, err = server.DepositFor(goCtx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
	})
	// ASSERT: The action should've failed due to non-existing round.
	require.ErrorContains(t, err, "round 0 not found")

	// ARRANGE: Report Ethereum Round #229.
	// https://etherscan.io/tx/0xcff68ffc6f79afadf835f559f8a51ed7092bc679d2a4f34cd153ef321d6bc8ec
	k.SetRound(ctx, 229, aggregator.RoundData{
		Answer:    sdk.NewInt(104572478),
		Balance:   sdk.NewInt(7016169453),
		Interest:  sdk.NewInt(1005815),
		Supply:    sdk.NewInt(67093843285741),
		UpdatedAt: 1717153499,
	})
	k.SetLastRoundId(ctx, 229)

	// ACT: Attempt to deposit for with insufficient funds.
	_, err = server.DepositFor(goCtx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Give user 202.40 $USDC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Underlying, amount))
	bank.Balances[recipient.Address] = sdk.Coins{}

	// ACT: Attempt to deposit for.
	_, err = server.DepositFor(goCtx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.True(t, bank.Balances[user.Address].IsZero())
	require.Equal(t, expected, bank.Balances[recipient.Address].AmountOf(k.Denom))
	require.True(t, bank.Balances[recipient.Address].AmountOf(k.Underlying).IsZero())
}

func TestWithdraw(t *testing.T) {
	// This test is based off of a real action on Ethereum.
	// https://etherscan.io/tx/0x325e6d83a4f1067db2a872e8e2a10a1bff79a2a8047db49a6ce080733b6a1159
	amount, expected := sdk.NewInt(150634259038), sdk.NewInt(154924310000)

	account := mocks.AccountKeeper{
		Accounts: make(map[string]authtypes.AccountI),
	}
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	ftf := mocks.FTFKeeper{
		Paused: false,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, account, bank, ftf)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate a user account.
	user := utils.TestAccount()

	// ACT: Attempt to withdraw with an invalid signer address.
	_, err := server.Withdraw(goCtx, &types.MsgWithdraw{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to withdraw without required permissions.
	_, err = server.Withdraw(goCtx, &types.MsgWithdraw{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgWithdraw")

	// ARRANGE: Assign the international feeder role to user.
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	// ARRANGE: Generate an owner account.
	owner := utils.TestAccount()
	k.SetOwner(ctx, owner.Address)
	// ARRANGE: Generate a withdrawal signature
	signature, err := owner.Key.Sign([]byte(fmt.Sprintf(
		"{\"halo_withdraw\":{\"recipient\":\"%s\",\"amount\":\"%s\",\"nonce\":%d}}",
		base64.StdEncoding.EncodeToString(user.Bytes),
		amount.String(),
		10,
	)))
	require.NoError(t, err)

	// ACT: Attempt to withdraw without owner public key in state.
	_, err = server.Withdraw(goCtx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to invalid signature.
	require.ErrorContains(t, err, types.ErrInvalidSignature.Error())

	// ARRANGE: Set owner public key in state.
	account.Accounts[owner.Address] = &authtypes.BaseAccount{
		PubKey: owner.PubKey,
	}

	// ACT: Attempt to withdraw with invalid signature (nonce mismatch).
	_, err = server.Withdraw(goCtx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to invalid signature.
	require.ErrorContains(t, err, types.ErrInvalidSignature.Error())

	// ARRANGE: Set user withdrawal nonce in state.
	k.SetNonce(ctx, user.Bytes, 10)
	// ARRANGE: Report Ethereum Round #139.
	// https://etherscan.io/tx/0x9095266d81856a28b80c4500228ab994197652fc4ad1c05cd4345d1454fccfd7
	k.SetRound(ctx, 139, aggregator.RoundData{
		Answer:    sdk.NewInt(102847997),
		Balance:   sdk.NewInt(4986480452),
		Interest:  sdk.NewInt(708258),
		Supply:    sdk.NewInt(48483293336746),
		UpdatedAt: 1706011979,
	})
	k.SetLastRoundId(ctx, 139)

	// ACT: Attempt to withdraw with insufficient funds.
	_, err = server.Withdraw(goCtx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Set user withdrawal nonce in state.
	k.SetNonce(ctx, user.Bytes, 10)
	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))

	// ACT: Attempt to withdraw with insufficient module funds.
	_, err = server.Withdraw(goCtx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to insufficient module funds.
	require.ErrorContains(t, err, "unable to transfer from module to account")

	// ARRANGE: Set user withdrawal nonce in state.
	k.SetNonce(ctx, user.Bytes, 10)
	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))
	// ARRANGE: Give module 154,924.31 $USDC.
	bank.Balances[types.ModuleAddress.String()] = sdk.NewCoins(sdk.NewCoin(k.Underlying, expected))

	// ACT: Attempt to withdraw.
	_, err = server.Withdraw(goCtx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.True(t, bank.Balances[user.Address].AmountOf(k.Denom).IsZero())
	require.Equal(t, expected, bank.Balances[user.Address].AmountOf(k.Underlying))
}

func TestWithdrawTo(t *testing.T) {
	// This test is based off of a real action on Ethereum.
	// https://etherscan.io/tx/0x325e6d83a4f1067db2a872e8e2a10a1bff79a2a8047db49a6ce080733b6a1159
	amount, expected := sdk.NewInt(150634259038), sdk.NewInt(154924310000)

	account := mocks.AccountKeeper{
		Accounts: make(map[string]authtypes.AccountI),
	}
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	ftf := mocks.FTFKeeper{
		Paused: false,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, account, bank, ftf)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate user and recipient accounts.
	user, recipient := utils.TestAccount(), utils.TestAccount()

	// ACT: Attempt to withdraw to with an invalid signer address.
	_, err := server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to withdraw to without required permissions.
	_, err = server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgWithdrawTo")

	// ARRANGE: Assign the international feeder role to user.
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to withdraw to with an invalid recipient address.
	_, err = server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Invalid,
	})
	// ASSERT: The action should've failed due to invalid recipient address.
	require.ErrorContains(t, err, "unable to decode recipient address")

	// ACT: Attempt to withdraw to without required recipient permissions.
	_, err = server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
	})
	// ASSERT: The action should've failed due to invalid recipient permissions.
	require.ErrorContains(t, err, "cannot receive uusyc")

	// ARRANGE: Assign the international feeder role to recipient.
	k.SetUserRole(ctx, recipient.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	// ARRANGE: Generate an owner account.
	owner := utils.TestAccount()
	k.SetOwner(ctx, owner.Address)
	// ARRANGE: Generate a withdrawal signature
	signature, err := owner.Key.Sign([]byte(fmt.Sprintf(
		"{\"halo_withdraw\":{\"recipient\":\"%s\",\"amount\":\"%s\",\"nonce\":%d}}",
		base64.StdEncoding.EncodeToString(recipient.Bytes),
		amount.String(),
		10,
	)))
	require.NoError(t, err)

	// ACT: Attempt to withdraw to without owner public key in state.
	_, err = server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to invalid signature.
	require.ErrorContains(t, err, types.ErrInvalidSignature.Error())

	// ARRANGE: Set owner public key in state.
	account.Accounts[owner.Address] = &authtypes.BaseAccount{
		PubKey: owner.PubKey,
	}

	// ACT: Attempt to withdraw to with invalid signature (nonce mismatch).
	_, err = server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to invalid signature.
	require.ErrorContains(t, err, types.ErrInvalidSignature.Error())

	// ARRANGE: Set user withdrawal nonce in state.
	k.SetNonce(ctx, recipient.Bytes, 10)

	// ACT: Attempt to withdraw to with non-existing last round.
	_, err = server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to non-existing last round.
	require.ErrorContains(t, err, "round 0 not found")

	// ARRANGE: Report Ethereum Round #139.
	// https://etherscan.io/tx/0x9095266d81856a28b80c4500228ab994197652fc4ad1c05cd4345d1454fccfd7
	k.SetRound(ctx, 139, aggregator.RoundData{
		Answer:    sdk.NewInt(102847997),
		Balance:   sdk.NewInt(4986480452),
		Interest:  sdk.NewInt(708258),
		Supply:    sdk.NewInt(48483293336746),
		UpdatedAt: 1706011979,
	})
	k.SetLastRoundId(ctx, 139)

	// ARRANGE: Set user withdrawal nonce in state.
	k.SetNonce(ctx, recipient.Bytes, 10)

	// ACT: Attempt to withdraw to with insufficient funds.
	_, err = server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Set user withdrawal nonce in state.
	k.SetNonce(ctx, recipient.Bytes, 10)
	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))

	// ACT: Attempt to withdraw to with insufficient module funds.
	_, err = server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to insufficient module funds.
	require.ErrorContains(t, err, "unable to transfer from module to account")

	// ARRANGE: Set user withdrawal nonce in state.
	k.SetNonce(ctx, recipient.Bytes, 10)
	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))
	// ARRANGE: Give module 154,924.31 $USDC.
	bank.Balances[types.ModuleAddress.String()] = sdk.NewCoins(sdk.NewCoin(k.Underlying, expected))

	// ACT: Attempt to withdraw to.
	_, err = server.WithdrawTo(goCtx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.True(t, bank.Balances[user.Address].IsZero())
	require.True(t, bank.Balances[recipient.Address].AmountOf(k.Denom).IsZero())
	require.Equal(t, expected, bank.Balances[recipient.Address].AmountOf(k.Underlying))
}

func TestWithdrawToAdmin(t *testing.T) {
	// This test is based off of a real action on Ethereum.
	// https://etherscan.io/tx/0x325e6d83a4f1067db2a872e8e2a10a1bff79a2a8047db49a6ce080733b6a1159
	amount, expected := sdk.NewInt(150634259038), sdk.NewInt(154924310000)

	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	ftf := mocks.FTFKeeper{
		Paused: false,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank, ftf)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate admin, user and recipient accounts.
	admin, user, recipient := utils.TestAccount(), utils.TestAccount(), utils.TestAccount()
	k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	k.SetUserRole(ctx, recipient.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to withdraw to admin with an invalid signer address.
	_, err := server.WithdrawToAdmin(goCtx, &types.MsgWithdrawToAdmin{
		Signer: utils.TestAccount().Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to withdraw to admin with an invalid signer.
	_, err = server.WithdrawToAdmin(goCtx, &types.MsgWithdrawToAdmin{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ACT: Attempt to withdraw to admin with an invalid from address.
	_, err = server.WithdrawToAdmin(goCtx, &types.MsgWithdrawToAdmin{
		Signer: admin.Address,
		From:   user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid from address.
	require.ErrorContains(t, err, "unable to decode from address")

	// ACT: Attempt to withdraw to admin with an invalid recipient address.
	_, err = server.WithdrawToAdmin(goCtx, &types.MsgWithdrawToAdmin{
		Signer:    admin.Address,
		From:      user.Address,
		Recipient: recipient.Invalid,
	})
	// ASSERT: The action should've failed due to invalid recipient address.
	require.ErrorContains(t, err, "unable to decode recipient address")

	// ARRANGE: Report Ethereum Round #139.
	// https://etherscan.io/tx/0x9095266d81856a28b80c4500228ab994197652fc4ad1c05cd4345d1454fccfd7
	k.SetRound(ctx, 139, aggregator.RoundData{
		Answer:    sdk.NewInt(102847997),
		Balance:   sdk.NewInt(4986480452),
		Interest:  sdk.NewInt(708258),
		Supply:    sdk.NewInt(48483293336746),
		UpdatedAt: 1706011979,
	})
	k.SetLastRoundId(ctx, 139)

	// ACT: Attempt to withdraw to admin with insufficient funds.
	_, err = server.WithdrawToAdmin(goCtx, &types.MsgWithdrawToAdmin{
		Signer:    admin.Address,
		From:      user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))

	// ACT: Attempt to withdraw to admin with insufficient module funds.
	_, err = server.WithdrawToAdmin(goCtx, &types.MsgWithdrawToAdmin{
		Signer:    admin.Address,
		From:      user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
	})
	// ASSERT: The action should've failed due to insufficient module funds.
	require.ErrorContains(t, err, "unable to transfer from module to account")

	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))
	// ARRANGE: Give module 154,924.31 $USDC.
	bank.Balances[types.ModuleAddress.String()] = sdk.NewCoins(sdk.NewCoin(k.Underlying, expected))

	// ACT: Attempt to withdraw to admin.
	_, err = server.WithdrawToAdmin(goCtx, &types.MsgWithdrawToAdmin{
		Signer:    admin.Address,
		From:      user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.True(t, bank.Balances[user.Address].IsZero())
	require.True(t, bank.Balances[recipient.Address].AmountOf(k.Denom).IsZero())
	require.Equal(t, expected, bank.Balances[recipient.Address].AmountOf(k.Underlying))
}

func TestBurn(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank, mocks.FTFKeeper{})
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate a user account.
	user := utils.TestAccount()

	// ACT: Attempt to burn with an invalid signer address.
	_, err := server.Burn(goCtx, &types.MsgBurn{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to burn without required permissions.
	_, err = server.Burn(goCtx, &types.MsgBurn{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgBurn")

	// ARRANGE: Assign the international feeder role to user.
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to burn with insufficient funds.
	_, err = server.Burn(goCtx, &types.MsgBurn{
		Signer: user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Give user 1 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, ONE))

	// ACT: Attempt to burn.
	_, err = server.Burn(goCtx, &types.MsgBurn{
		Signer: user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.True(t, bank.Balances[user.Address].IsZero())
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestBurnFor(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank, mocks.FTFKeeper{})
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate admin and user accounts.
	admin, user := utils.TestAccount(), utils.TestAccount()
	k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to burn for with an invalid signer address.
	_, err := server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to burn for with an invalid signer.
	_, err = server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ACT: Attempt to burn for with an invalid from address.
	_, err = server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: admin.Address,
		From:   user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid from address.
	require.ErrorContains(t, err, "unable to decode from address")

	// ACT: Attempt to burn for with insufficient funds.
	_, err = server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: admin.Address,
		From:   user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Give user 1 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, ONE))

	// ACT: Attempt to burn.
	_, err = server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: admin.Address,
		From:   user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.True(t, bank.Balances[user.Address].IsZero())
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestMint(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank, mocks.FTFKeeper{})
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate admin and user accounts.
	admin, user := utils.TestAccount(), utils.TestAccount()
	k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to mint with an invalid signer address.
	_, err := server.Mint(goCtx, &types.MsgMint{
		Signer: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to mint with an invalid signer.
	_, err = server.Mint(goCtx, &types.MsgMint{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ACT: Attempt to mint with an invalid to address.
	_, err = server.Mint(goCtx, &types.MsgMint{
		Signer: admin.Address,
		To:     user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid to address.
	require.ErrorContains(t, err, "unable to decode to address")

	// ACT: Attempt to mint without required permissions.
	_, err = server.Mint(goCtx, &types.MsgMint{
		Signer: admin.Address,
		To:     utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot transfer")

	// ACT: Attempt to mint.
	_, err = server.Mint(goCtx, &types.MsgMint{
		Signer: admin.Address,
		To:     user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.Equal(t, ONE, bank.Balances[user.Address].AmountOf(k.Denom))
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestTradeToFiat(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	ftf := mocks.FTFKeeper{
		Paused: false,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank, ftf)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate an admin account.
	admin := utils.TestAccount()

	// ACT: Attempt to trade to fiat with an invalid signer address.
	_, err := server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to trade to fiat with an invalid signer.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer: admin.Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ARRANGE: Set fund admin in state.
	k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)

	// ACT: Attempt to trade to fiat with an invalid recipient address.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Recipient: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid recipient address.
	require.ErrorContains(t, err, "unable to decode recipient address")

	// ACT: Attempt to trade to fiat with invalid recipient permissions.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Recipient: admin.Address,
	})
	// ASSERT: The action should've failed due to invalid recipient permissions.
	require.ErrorContains(t, err, types.ErrInvalidLiquidityProvider.Error())

	// ARRANGE: Set liquidity provider in state.
	k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_LIQUIDITY_PROVIDER, true)

	// ACT: Attempt to trade to fiat with insufficient funds.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Amount:    ONE,
		Recipient: admin.Address,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "insufficient funds")

	// ARRANGE: Give the module 1 $USDC.
	bank.Balances[types.ModuleAddress.String()] = sdk.NewCoins(sdk.NewCoin(k.Underlying, ONE))

	// ACT: Attempt to trade to fiat.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Amount:    ONE,
		Recipient: admin.Address,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.Equal(t, ONE, bank.Balances[admin.Address].AmountOf(k.Underlying))
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestTransferOwnership(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ACT: Attempt to transfer ownership with no owner set.
	_, err := server.TransferOwnership(goCtx, &types.MsgTransferOwnership{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, types.ErrNoOwner.Error())

	// ARRANGE: Set owner in state.
	owner := utils.TestAccount()
	k.SetOwner(ctx, owner.Address)

	// ACT: Attempt to transfer ownership with invalid signer.
	_, err = server.TransferOwnership(goCtx, &types.MsgTransferOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidOwner.Error())

	// ACT: Attempt to transfer ownership to same address.
	_, err = server.TransferOwnership(goCtx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: owner.Address,
	})
	// ASSERT: The action should've failed due to same address.
	require.ErrorContains(t, err, types.ErrSameOwner.Error())

	// ARRANGE: Generate a new owner account.
	newOwner := utils.TestAccount()

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(goCtx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've succeeded, and set owner in state.
	require.NoError(t, err)
	require.Equal(t, newOwner.Address, k.GetOwner(ctx))
}
