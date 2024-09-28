// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"context"
	"strings"

	"cosmossdk.io/errors"
	"github.com/noble-assets/halo/v2/types/entitlements"
)

var _ entitlements.MsgServer = &entitlementsMsgServer{}

type entitlementsMsgServer struct {
	*Keeper
}

func NewEntitlementsMsgServer(keeper *Keeper) entitlements.MsgServer {
	return &entitlementsMsgServer{Keeper: keeper}
}

func (k entitlementsMsgServer) SetPublicCapability(ctx context.Context, msg *entitlements.MsgSetPublicCapability) (*entitlements.MsgSetPublicCapabilityResponse, error) {
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	resolved, _ := k.interfaceRegistry.Resolve(msg.Method)
	if !(msg.Method == "transfer" || (resolved != nil && strings.HasPrefix(msg.Method, "/halo"))) {
		return nil, errors.Wrapf(entitlements.ErrInvalidMethod, "method %s does not exist or is not allowed", msg.Method)
	}

	if err = k.Keeper.SetPublicCapability(ctx, msg.Method, msg.Enabled); err != nil {
		return nil, err
	}

	return &entitlements.MsgSetPublicCapabilityResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &entitlements.PublicCapabilityUpdated{
		Method:  msg.Method,
		Enabled: msg.Enabled,
	})
}

func (k entitlementsMsgServer) SetRoleCapability(ctx context.Context, msg *entitlements.MsgSetRoleCapability) (*entitlements.MsgSetRoleCapabilityResponse, error) {
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	_, roleExists := entitlements.Role_value[msg.Role.String()]
	if !roleExists {
		return nil, errors.Wrapf(entitlements.ErrInvalidRole, "role %s does not exist", msg.Role)
	}

	resolved, _ := k.interfaceRegistry.Resolve(msg.Method)
	if !(msg.Method == "transfer" || (resolved != nil && strings.HasPrefix(msg.Method, "/halo"))) {
		return nil, errors.Wrapf(entitlements.ErrInvalidMethod, "method %s does not exist or is not allowed", msg.Method)
	}

	if err = k.Keeper.SetRoleCapability(ctx, msg.Method, msg.Role, msg.Enabled); err != nil {
		return nil, err
	}

	return &entitlements.MsgSetRoleCapabilityResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &entitlements.RoleCapabilityUpdated{
		Role:    msg.Role,
		Method:  msg.Method,
		Enabled: msg.Enabled,
	})
}

func (k entitlementsMsgServer) SetUserRole(ctx context.Context, msg *entitlements.MsgSetUserRole) (*entitlements.MsgSetUserRoleResponse, error) {
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	user, err := k.addressCodec.StringToBytes(msg.User)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode user address %s", msg.User)
	}

	_, roleExists := entitlements.Role_value[msg.Role.String()]
	if !roleExists {
		return nil, errors.Wrapf(entitlements.ErrInvalidRole, "role %s does not exist", msg.Role)
	}

	if err = k.Keeper.SetUserRole(ctx, user, msg.Role, msg.Enabled); err != nil {
		return nil, err
	}

	return &entitlements.MsgSetUserRoleResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &entitlements.UserRoleUpdated{
		User:    msg.User,
		Role:    msg.Role,
		Enabled: msg.Enabled,
	})
}

func (k entitlementsMsgServer) Pause(ctx context.Context, msg *entitlements.MsgPause) (*entitlements.MsgPauseResponse, error) {
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if err = k.SetPaused(ctx, true); err != nil {
		return nil, err
	}

	return &entitlements.MsgPauseResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &entitlements.Paused{
		Account: msg.Signer,
	})
}

func (k entitlementsMsgServer) Unpause(ctx context.Context, msg *entitlements.MsgUnpause) (*entitlements.MsgUnpauseResponse, error) {
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if err = k.SetPaused(ctx, false); err != nil {
		return nil, err
	}

	return &entitlements.MsgUnpauseResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &entitlements.Unpaused{
		Account: msg.Signer,
	})
}

func (k entitlementsMsgServer) TransferOwnership(ctx context.Context, msg *entitlements.MsgTransferOwnership) (*entitlements.MsgTransferOwnershipResponse, error) {
	owner, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if msg.NewOwner == owner {
		return nil, entitlements.ErrSameOwner
	}

	if err = k.SetEntitlementsOwner(ctx, msg.NewOwner); err != nil {
		return nil, err
	}

	return &entitlements.MsgTransferOwnershipResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &entitlements.OwnershipTransferred{
		PreviousOwner: owner,
		NewOwner:      msg.NewOwner,
	})
}

//

func (k entitlementsMsgServer) EnsureOwner(ctx context.Context, signer string) (string, error) {
	owner := k.GetEntitlementsOwner(ctx)
	if owner == "" {
		return "", entitlements.ErrNoOwner
	}
	if signer != owner {
		return "", errors.Wrapf(entitlements.ErrInvalidOwner, "expected %s, got %s", owner, signer)
	}
	return owner, nil
}
