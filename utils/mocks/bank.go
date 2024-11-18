// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package mocks

import (
	"context"
	"fmt"

	"github.com/noble-assets/halo/v3/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var _ types.BankKeeper = BankKeeper{}

type BankKeeper struct {
	Balances    map[string]sdk.Coins
	Restriction SendRestrictionFn
}

func (k BankKeeper) BurnCoins(_ context.Context, moduleName string, amt sdk.Coins) error {
	address := authtypes.NewModuleAddress(moduleName).String()

	balance := k.Balances[address]
	newBalance, negative := balance.SafeSub(amt...)
	if negative {
		return sdkerrors.Wrapf(errors.ErrInsufficientFunds, "%s is smaller than %s", balance, amt)
	}

	k.Balances[address] = newBalance

	return nil
}

func (k BankKeeper) MintCoins(_ context.Context, moduleName string, amt sdk.Coins) error {
	address := authtypes.NewModuleAddress(moduleName).String()
	k.Balances[address] = k.Balances[address].Add(amt...)

	return nil
}

func (k BankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	recipientAddr := authtypes.NewModuleAddress(recipientModule)

	return k.SendCoins(ctx, senderAddr, recipientAddr, amt)
}

func (k BankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	senderAddr := authtypes.NewModuleAddress(senderModule)

	return k.SendCoins(ctx, senderAddr, recipientAddr, amt)
}

//

type SendRestrictionFn func(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) (newToAddr sdk.AccAddress, err error)

func NoOpSendRestrictionFn(_ context.Context, _, toAddr sdk.AccAddress, _ sdk.Coins) (sdk.AccAddress, error) {
	return toAddr, nil
}

func FailingSendRestrictionFn(_ context.Context, _, toAddr sdk.AccAddress, _ sdk.Coins) (sdk.AccAddress, error) {
	return nil, fmt.Errorf("%s is blocked from sending/receiving", toAddr.String())
}

func (k BankKeeper) WithSendCoinsRestriction(check SendRestrictionFn) BankKeeper {
	oldRestriction := k.Restriction
	k.Restriction = func(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) (newToAddr sdk.AccAddress, err error) {
		newToAddr, err = check(ctx, fromAddr, toAddr, amt)
		if err != nil {
			return newToAddr, err
		}
		return oldRestriction(ctx, fromAddr, toAddr, amt)
	}
	return k
}

func (k BankKeeper) SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	toAddr, err := k.Restriction(ctx, fromAddr, toAddr, amt)
	if err != nil {
		return err
	}

	balance := k.Balances[fromAddr.String()]
	newBalance, negative := balance.SafeSub(amt...)
	if negative {
		return sdkerrors.Wrapf(errors.ErrInsufficientFunds, "%s is smaller than %s", balance, amt)
	}

	k.Balances[fromAddr.String()] = newBalance
	k.Balances[toAddr.String()] = k.Balances[toAddr.String()].Add(amt...)

	return nil
}

//

func init() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("noble", "noblepub")
	config.Seal()
}
