// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/noble-assets/halo/v2/keeper"
	"github.com/noble-assets/halo/v2/types"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/noble-assets/halo/v2/types/entitlements"
	"github.com/noble-assets/halo/v2/utils"
	"github.com/noble-assets/halo/v2/utils/mocks"
	"github.com/stretchr/testify/require"
)

var ONE = math.NewInt(1_000_000)

func TestDeposit(t *testing.T) {
	// This test is based off of a real action on Ethereum.
	// https://etherscan.io/tx/0x9fc3dc5ec218eea1e84bed8f2a0f2aaf243b355c560ddd30820cf3ea0358fec0
	amount, expected := math.NewInt(202400000), math.NewInt(193549970)

	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate a user account.
	user := utils.TestAccount()

	// ACT: Attempt to deposit with an invalid signer address.
	_, err := server.Deposit(ctx, &types.MsgDeposit{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to deposit without required permissions.
	_, err = server.Deposit(ctx, &types.MsgDeposit{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgDeposit")

	// ARRANGE: Assign the international feeder role to user.
	err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)
	// ARRANGE: Report Ethereum Round #229.
	// https://etherscan.io/tx/0xcff68ffc6f79afadf835f559f8a51ed7092bc679d2a4f34cd153ef321d6bc8ec
	err = k.Rounds.Set(ctx, 229, aggregator.RoundData{
		Answer:    math.NewInt(104572478),
		UpdatedAt: 1717153499,
	})
	require.NoError(t, err)
	err = k.SetLastRoundId(ctx, 229)
	require.NoError(t, err)

	// ACT: Attempt to deposit with insufficient funds.
	_, err = server.Deposit(ctx, &types.MsgDeposit{
		Signer: user.Address,
		Amount: amount,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Give user 202.40 $USDC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Underlying, amount))

	// ACT: Attempt to deposit for with a negative amount.
	_, err = server.Deposit(ctx, &types.MsgDeposit{
		Signer: user.Address,
		Amount: math.NewInt(-20000),
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "invalid amount")

	// ACT: Attempt to deposit.
	_, err = server.Deposit(ctx, &types.MsgDeposit{
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
	amount, expected := math.NewInt(202400000), math.NewInt(193549970)

	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate user and recipient accounts.
	user, recipient := utils.TestAccount(), utils.TestAccount()

	// ACT: Attempt to deposit for with an invalid signer address.
	_, err := server.DepositFor(ctx, &types.MsgDepositFor{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to deposit for without required permissions.
	_, err = server.DepositFor(ctx, &types.MsgDepositFor{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgDepositFor")

	// ARRANGE: Assign the international feeder role to user.
	err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)

	// ACT: Attempt to deposit for with an invalid recipient address.
	_, err = server.DepositFor(ctx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Invalid,
	})
	// ASSERT: The action should've failed due to invalid recipient address.
	require.ErrorContains(t, err, "unable to decode recipient address")

	// ACT: Attempt to deposit for without required recipient permissions.
	_, err = server.DepositFor(ctx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Address,
	})
	// ASSERT: The action should've failed due to invalid recipient permissions.
	require.ErrorContains(t, err, "cannot receive uusyc")

	// ARRANGE: Assign the international feeder role to recipient.
	err = k.SetUserRole(ctx, recipient.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)

	// ACT: Attempt to deposit for with non-existing round.
	_, err = server.DepositFor(ctx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
	})
	// ASSERT: The action should've failed due to non-existing round.
	require.ErrorContains(t, err, "round 0 not found")

	// ARRANGE: Report Ethereum Round #229.
	// https://etherscan.io/tx/0xcff68ffc6f79afadf835f559f8a51ed7092bc679d2a4f34cd153ef321d6bc8ec
	err = k.Rounds.Set(ctx, 229, aggregator.RoundData{
		Answer:    math.NewInt(104572478),
		UpdatedAt: 1717153499,
	})
	require.NoError(t, err)
	err = k.SetLastRoundId(ctx, 229)
	require.NoError(t, err)

	// ACT: Attempt to deposit for with insufficient funds.
	_, err = server.DepositFor(ctx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ACT: Attempt to deposit for with a negative amount.
	_, err = server.DepositFor(ctx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    math.NewInt(-202400000),
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "invalid amount")

	// ARRANGE: Give user 202.40 $USDC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Underlying, amount))
	bank.Balances[recipient.Address] = sdk.Coins{}

	// ACT: Attempt to deposit for.
	_, err = server.DepositFor(ctx, &types.MsgDepositFor{
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

func TestDepositForWithRestrictions(t *testing.T) {
	amount := math.NewInt(202400000)
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.FailingSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate user and recipient accounts.
	user, recipient := utils.TestAccount(), utils.TestAccount()

	// ARRANGE: Assign the international feeder role to user.
	err := k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)
	// ARRANGE: Assign the international feeder role to recipient.
	err = k.SetUserRole(ctx, recipient.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)

	// ARRANGE: Report Ethereum Round #229.
	// https://etherscan.io/tx/0xcff68ffc6f79afadf835f559f8a51ed7092bc679d2a4f34cd153ef321d6bc8ec
	err = k.Rounds.Set(ctx, 229, aggregator.RoundData{
		Answer:    math.NewInt(104572478),
		UpdatedAt: 1717153499,
	})
	require.NoError(t, err)
	err = k.SetLastRoundId(ctx, 229)
	require.NoError(t, err)

	// ARRANGE: Give user 202.40 $USDC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Underlying, amount))
	bank.Balances[recipient.Address] = sdk.Coins{}

	// ACT: Attempt to deposit for.
	_, err = server.DepositFor(ctx, &types.MsgDepositFor{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
	})
	// ASSERT: The action should fail due to restrictions.
	require.ErrorContains(t, err, "unable to transfer from module to account")
}

func TestWithdraw(t *testing.T) {
	// This test is based off of a real action on Ethereum.
	// https://etherscan.io/tx/0x325e6d83a4f1067db2a872e8e2a10a1bff79a2a8047db49a6ce080733b6a1159
	amount, expected := math.NewInt(150634259038), math.NewInt(154924310000)

	account := mocks.AccountKeeper{
		Accounts: make(map[string]sdk.AccountI),
	}
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, account, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate a user account.
	user := utils.TestAccount()

	// ACT: Attempt to withdraw with an invalid signer address.
	_, err := server.Withdraw(ctx, &types.MsgWithdraw{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to withdraw without required permissions.
	_, err = server.Withdraw(ctx, &types.MsgWithdraw{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgWithdraw")

	// ARRANGE: Assign the international feeder role to user.
	err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)
	// ARRANGE: Generate an owner account.
	owner := utils.TestAccount()
	err = k.SetOwner(ctx, owner.Address)
	require.NoError(t, err)
	// ARRANGE: Generate a withdrawal signature
	signature, err := owner.Key.Sign([]byte(fmt.Sprintf(
		"{\"halo_withdraw\":{\"recipient\":\"%s\",\"amount\":\"%s\",\"nonce\":%d}}",
		base64.StdEncoding.EncodeToString(user.Bytes),
		amount.String(),
		10,
	)))
	require.NoError(t, err)

	// ACT: Attempt to withdraw without owner public key in state.
	_, err = server.Withdraw(ctx, &types.MsgWithdraw{
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
	_, err = server.Withdraw(ctx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to invalid signature.
	require.ErrorContains(t, err, types.ErrInvalidSignature.Error())

	// ARRANGE: Set user withdrawal nonce in state.
	err = k.SetNonce(ctx, user.Bytes, 10)
	require.NoError(t, err)
	// ARRANGE: Report Ethereum Round #139.
	// https://etherscan.io/tx/0x9095266d81856a28b80c4500228ab994197652fc4ad1c05cd4345d1454fccfd7
	err = k.Rounds.Set(ctx, 139, aggregator.RoundData{
		Answer:    math.NewInt(102847997),
		UpdatedAt: 1706011979,
	})
	require.NoError(t, err)
	err = k.SetLastRoundId(ctx, 139)
	require.NoError(t, err)

	// ACT: Attempt to withdraw with insufficient funds.
	_, err = server.Withdraw(ctx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Set user withdrawal nonce in state.
	err = k.SetNonce(ctx, user.Bytes, 10)
	require.NoError(t, err)
	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))

	// ACT: Attempt to withdraw with insufficient module funds.
	_, err = server.Withdraw(ctx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to insufficient module funds.
	require.ErrorContains(t, err, "unable to transfer from module to account")

	// ARRANGE: Set user withdrawal nonce in state.
	err = k.SetNonce(ctx, user.Bytes, 10)
	require.NoError(t, err)
	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))
	// ARRANGE: Give module 154,924.31 $USDC.
	bank.Balances[types.ModuleAddress.String()] = sdk.NewCoins(sdk.NewCoin(k.Underlying, expected))

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmp := k.Nonces
	k.Nonces = collections.NewMap(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		types.NoncePrefix, "nonces", collections.BytesKey, collections.Uint64Value,
	)

	// ACT: Attempt to withdraw with failing Nonces collection store.
	_, err = server.Withdraw(ctx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.Nonces = tmp

	// ACT: Attempt to withdraw.
	_, err = server.Withdraw(ctx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.True(t, bank.Balances[user.Address].AmountOf(k.Denom).IsZero())
	require.Equal(t, expected, bank.Balances[user.Address].AmountOf(k.Underlying))

	// ACT: Attempt to withdraw with a negative amount.
	negativeSignature, _ := owner.Key.Sign([]byte(fmt.Sprintf(
		"{\"halo_withdraw\":{\"recipient\":\"%s\",\"amount\":\"%s\",\"nonce\":%d}}",
		base64.StdEncoding.EncodeToString(user.Bytes),
		math.NewInt(-200).String(),
		11,
	)))
	_, err = server.Withdraw(ctx, &types.MsgWithdraw{
		Signer:    user.Address,
		Amount:    math.NewInt(-200),
		Signature: negativeSignature,
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "invalid amount")
}

func TestWithdrawTo(t *testing.T) {
	// This test is based off of a real action on Ethereum.
	// https://etherscan.io/tx/0x325e6d83a4f1067db2a872e8e2a10a1bff79a2a8047db49a6ce080733b6a1159
	amount, expected := math.NewInt(150634259038), math.NewInt(154924310000)

	account := mocks.AccountKeeper{
		Accounts: make(map[string]sdk.AccountI),
	}
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, account, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate user and recipient accounts.
	user, recipient := utils.TestAccount(), utils.TestAccount()

	// ACT: Attempt to withdraw to with an invalid signer address.
	_, err := server.WithdrawTo(ctx, &types.MsgWithdrawTo{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to withdraw to without required permissions.
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgWithdrawTo")

	// ARRANGE: Assign the international feeder role to user.
	err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)

	// ACT: Attempt to withdraw to with an invalid recipient address.
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Invalid,
	})
	// ASSERT: The action should've failed due to invalid recipient address.
	require.ErrorContains(t, err, "unable to decode recipient address")

	// ACT: Attempt to withdraw to without required recipient permissions.
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
	})
	// ASSERT: The action should've failed due to invalid recipient permissions.
	require.ErrorContains(t, err, "cannot receive uusyc")

	// ARRANGE: Assign the international feeder role to recipient.
	err = k.SetUserRole(ctx, recipient.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)
	// ARRANGE: Generate an owner account.
	owner := utils.TestAccount()
	err = k.SetOwner(ctx, owner.Address)
	require.NoError(t, err)
	// ARRANGE: Generate a withdrawal signature
	signature, err := owner.Key.Sign([]byte(fmt.Sprintf(
		"{\"halo_withdraw\":{\"recipient\":\"%s\",\"amount\":\"%s\",\"nonce\":%d}}",
		base64.StdEncoding.EncodeToString(recipient.Bytes),
		amount.String(),
		10,
	)))
	require.NoError(t, err)

	// ACT: Attempt to withdraw to without owner public key in state.
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
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
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to invalid signature.
	require.ErrorContains(t, err, types.ErrInvalidSignature.Error())

	// ARRANGE: Set user withdrawal nonce in state.
	err = k.SetNonce(ctx, recipient.Bytes, 10)
	require.NoError(t, err)

	// ACT: Attempt to withdraw to with non-existing last round.
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to non-existing last round.
	require.ErrorContains(t, err, "round 0 not found")

	// ARRANGE: Report Ethereum Round #139.
	// https://etherscan.io/tx/0x9095266d81856a28b80c4500228ab994197652fc4ad1c05cd4345d1454fccfd7
	err = k.Rounds.Set(ctx, 139, aggregator.RoundData{
		Answer:    math.NewInt(102847997),
		UpdatedAt: 1706011979,
	})
	require.NoError(t, err)
	err = k.SetLastRoundId(ctx, 139)
	require.NoError(t, err)

	// ARRANGE: Set user withdrawal nonce in state.
	err = k.SetNonce(ctx, recipient.Bytes, 10)
	require.NoError(t, err)

	// ACT: Attempt to withdraw to with insufficient funds.
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Set user withdrawal nonce in state.
	err = k.SetNonce(ctx, recipient.Bytes, 10)
	require.NoError(t, err)
	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))

	// ACT: Attempt to withdraw to with insufficient module funds.
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    amount,
		Signature: signature,
	})
	// ASSERT: The action should've failed due to insufficient module funds.
	require.ErrorContains(t, err, "unable to transfer from module to account")

	// ARRANGE: Set user withdrawal nonce in state.
	err = k.SetNonce(ctx, recipient.Bytes, 10)
	require.NoError(t, err)
	// ARRANGE: Give user 150,634.259038 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, amount))
	// ARRANGE: Give module 154,924.31 $USDC.
	bank.Balances[types.ModuleAddress.String()] = sdk.NewCoins(sdk.NewCoin(k.Underlying, expected))

	// ACT: Attempt to withdraw to.
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
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

	// ACT: Attempt to withdraw to with a negative amount.
	negativeSignature, _ := owner.Key.Sign([]byte(fmt.Sprintf(
		"{\"halo_withdraw\":{\"recipient\":\"%s\",\"amount\":\"%s\",\"nonce\":%d}}",
		base64.StdEncoding.EncodeToString(recipient.Bytes),
		math.NewInt(-200).String(),
		11,
	)))
	_, err = server.WithdrawTo(ctx, &types.MsgWithdrawTo{
		Signer:    user.Address,
		Recipient: recipient.Address,
		Amount:    math.NewInt(-200),
		Signature: negativeSignature,
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "invalid amount")
}

func TestWithdrawToAdmin(t *testing.T) {
	// This test is based off of a real action on Ethereum.
	// https://etherscan.io/tx/0x325e6d83a4f1067db2a872e8e2a10a1bff79a2a8047db49a6ce080733b6a1159
	amount, expected := math.NewInt(150634259038), math.NewInt(154924310000)

	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate admin, user and recipient accounts.
	admin, user, recipient := utils.TestAccount(), utils.TestAccount(), utils.TestAccount()
	err := k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	require.NoError(t, err)
	err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)
	err = k.SetUserRole(ctx, recipient.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)

	// ACT: Attempt to withdraw to admin with an invalid signer address.
	_, err = server.WithdrawToAdmin(ctx, &types.MsgWithdrawToAdmin{
		Signer: utils.TestAccount().Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to withdraw to admin with an invalid signer.
	_, err = server.WithdrawToAdmin(ctx, &types.MsgWithdrawToAdmin{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ACT: Attempt to withdraw to admin with an invalid from address.
	_, err = server.WithdrawToAdmin(ctx, &types.MsgWithdrawToAdmin{
		Signer: admin.Address,
		From:   user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid from address.
	require.ErrorContains(t, err, "unable to decode from address")

	// ACT: Attempt to withdraw to admin with an invalid recipient address.
	_, err = server.WithdrawToAdmin(ctx, &types.MsgWithdrawToAdmin{
		Signer:    admin.Address,
		From:      user.Address,
		Recipient: recipient.Invalid,
	})
	// ASSERT: The action should've failed due to invalid recipient address.
	require.ErrorContains(t, err, "unable to decode recipient address")

	// ARRANGE: Report Ethereum Round #139.
	// https://etherscan.io/tx/0x9095266d81856a28b80c4500228ab994197652fc4ad1c05cd4345d1454fccfd7
	err = k.Rounds.Set(ctx, 139, aggregator.RoundData{
		Answer:    math.NewInt(102847997),
		UpdatedAt: 1706011979,
	})
	require.NoError(t, err)
	err = k.SetLastRoundId(ctx, 139)
	require.NoError(t, err)

	// ACT: Attempt to withdraw to admin with insufficient funds.
	_, err = server.WithdrawToAdmin(ctx, &types.MsgWithdrawToAdmin{
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
	_, err = server.WithdrawToAdmin(ctx, &types.MsgWithdrawToAdmin{
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
	_, err = server.WithdrawToAdmin(ctx, &types.MsgWithdrawToAdmin{
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
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate a user account.
	user := utils.TestAccount()

	// ACT: Attempt to burn with an invalid signer address.
	_, err := server.Burn(ctx, &types.MsgBurn{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to burn without required permissions.
	_, err = server.Burn(ctx, &types.MsgBurn{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgBurn")

	// ARRANGE: Assign the international feeder role to user.
	err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)

	// ACT: Attempt to burn with insufficient funds.
	_, err = server.Burn(ctx, &types.MsgBurn{
		Signer: user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ACT: Attempt to burn with negative amount.
	_, err = server.Burn(ctx, &types.MsgBurn{
		Signer: user.Address,
		Amount: math.NewInt(-1_000_000),
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "invalid amount")

	// ARRANGE: Give user 1 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, ONE))

	// ACT: Attempt to burn.
	_, err = server.Burn(ctx, &types.MsgBurn{
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
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate admin and user accounts.
	admin, user := utils.TestAccount(), utils.TestAccount()
	err := k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	require.NoError(t, err)
	err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)

	// ACT: Attempt to burn for with an invalid signer address.
	_, err = server.BurnFor(ctx, &types.MsgBurnFor{
		Signer: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to burn for with an invalid signer.
	_, err = server.BurnFor(ctx, &types.MsgBurnFor{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ACT: Attempt to burn for with an invalid from address.
	_, err = server.BurnFor(ctx, &types.MsgBurnFor{
		Signer: admin.Address,
		From:   user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid from address.
	require.ErrorContains(t, err, "unable to decode from address")

	// ACT: Attempt to burn for with insufficient funds.
	_, err = server.BurnFor(ctx, &types.MsgBurnFor{
		Signer: admin.Address,
		From:   user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Give user 1 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, ONE))

	// ACT: Attempt to burn for with negative amount.
	_, err = server.BurnFor(ctx, &types.MsgBurnFor{
		Signer: admin.Address,
		From:   user.Address,
		Amount: math.NewInt(-1_000_000),
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "invalid amount")

	// ACT: Attempt to burn.
	_, err = server.BurnFor(ctx, &types.MsgBurnFor{
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
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate admin and user accounts.
	admin, user := utils.TestAccount(), utils.TestAccount()
	err := k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	require.NoError(t, err)
	err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)

	// ACT: Attempt to mint with an invalid signer address.
	_, err = server.Mint(ctx, &types.MsgMint{
		Signer: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to mint with an invalid signer.
	_, err = server.Mint(ctx, &types.MsgMint{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ACT: Attempt to mint with an invalid to address.
	_, err = server.Mint(ctx, &types.MsgMint{
		Signer: admin.Address,
		To:     user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid to address.
	require.ErrorContains(t, err, "unable to decode to address")

	// ACT: Attempt to mint without required permissions.
	_, err = server.Mint(ctx, &types.MsgMint{
		Signer: admin.Address,
		To:     utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot transfer")

	// ACT: Attempt to mint with negative amount.
	_, err = server.Mint(ctx, &types.MsgMint{
		Signer: admin.Address,
		To:     user.Address,
		Amount: math.NewInt(-1_000_000),
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "invalid amount")

	// ACT: Attempt to mint.
	_, err = server.Mint(ctx, &types.MsgMint{
		Signer: admin.Address,
		To:     user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.Equal(t, ONE, bank.Balances[user.Address].AmountOf(k.Denom))
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestMintWithRestrictions(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.FailingSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate admin and user accounts.
	admin, user := utils.TestAccount(), utils.TestAccount()
	err := k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	require.NoError(t, err)
	err = k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)
	require.NoError(t, err)
	// ACT: Attempt to mint with send restrictions.
	_, err = server.Mint(ctx, &types.MsgMint{
		Signer: admin.Address,
		To:     user.Address,
		Amount: math.NewInt(1_000_000),
	})
	// ASSERT: The action should've failed due to restrictions.
	require.ErrorContains(t, err, "unable to transfer from module to account")
}

func TestTradeToFiat(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate an admin and recipient account.
	admin, recipient := utils.TestAccount(), utils.TestAccount()

	// ACT: Attempt to trade to fiat with an invalid signer address.
	_, err := server.TradeToFiat(ctx, &types.MsgTradeToFiat{
		Signer: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to trade to fiat with an invalid signer.
	_, err = server.TradeToFiat(ctx, &types.MsgTradeToFiat{
		Signer: admin.Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ARRANGE: Set fund admin in state.
	err = k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	require.NoError(t, err)

	// ACT: Attempt to trade to fiat with an invalid recipient address.
	_, err = server.TradeToFiat(ctx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Recipient: recipient.Invalid,
	})
	// ASSERT: The action should've failed due to invalid recipient address.
	require.ErrorContains(t, err, "unable to decode recipient address")

	// ACT: Attempt to trade to fiat with invalid recipient permissions.
	_, err = server.TradeToFiat(ctx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Recipient: recipient.Address,
	})
	// ASSERT: The action should've failed due to invalid recipient permissions.
	require.ErrorContains(t, err, types.ErrInvalidLiquidityProvider.Error())

	// ARRANGE: Set liquidity provider in state.
	err = k.SetUserRole(ctx, recipient.Bytes, entitlements.ROLE_LIQUIDITY_PROVIDER, true)
	require.NoError(t, err)

	// ACT: Attempt to trade to fiat with insufficient funds.
	_, err = server.TradeToFiat(ctx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Amount:    ONE,
		Recipient: recipient.Address,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "insufficient funds")

	// ARRANGE: Give the module 1 $USDC.
	bank.Balances[types.ModuleAddress.String()] = sdk.NewCoins(sdk.NewCoin(k.Underlying, ONE))

	// ACT: Attempt to trade to fiat with negative amount.
	_, err = server.TradeToFiat(ctx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Amount:    math.NewInt(-1_000_000),
		Recipient: recipient.Address,
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "invalid amount")

	// ACT: Attempt to trade to fiat.
	_, err = server.TradeToFiat(ctx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Amount:    ONE,
		Recipient: recipient.Address,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.Equal(t, ONE, bank.Balances[recipient.Address].AmountOf(k.Underlying))
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestTransferOwnership(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewMsgServer(k)

	// ACT: Attempt to transfer ownership with no owner set.
	_, err := server.TransferOwnership(ctx, &types.MsgTransferOwnership{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, types.ErrNoOwner.Error())

	// ARRANGE: Set owner in state.
	owner := utils.TestAccount()
	err = k.SetOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt to transfer ownership with invalid signer.
	_, err = server.TransferOwnership(ctx, &types.MsgTransferOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidOwner.Error())

	// ACT: Attempt to transfer ownership to same address.
	_, err = server.TransferOwnership(ctx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: owner.Address,
	})
	// ASSERT: The action should've failed due to same address.
	require.ErrorContains(t, err, types.ErrSameOwner.Error())

	// ARRANGE: Generate a new owner account.
	newOwner := utils.TestAccount()

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmp := k.Owner
	k.Owner = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		types.OwnerKey, "owner", collections.StringValue,
	)

	// ACT: Attempt to transfer ownership with failing Owner collection store.
	_, err = server.TransferOwnership(ctx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.Owner = tmp

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(ctx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've succeeded, and set owner in state.
	require.NoError(t, err)
	require.Equal(t, newOwner.Address, k.GetOwner(ctx))
}
