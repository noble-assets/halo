package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
)

var _ entitlements.MsgServer = &entitlementsMsgServer{}

type entitlementsMsgServer struct {
	*Keeper
}

func NewEntitlementsMsgServer(keeper *Keeper) entitlements.MsgServer {
	return &entitlementsMsgServer{Keeper: keeper}
}

func (k entitlementsMsgServer) SetPublicCapability(goCtx context.Context, msg *entitlements.MsgSetPublicCapability) (*entitlements.MsgSetPublicCapabilityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	k.Keeper.SetPublicCapability(ctx, msg.Method, msg.Enabled)

	return &entitlements.MsgSetPublicCapabilityResponse{}, ctx.EventManager().EmitTypedEvent(&entitlements.PublicCapabilityUpdated{
		Method:  msg.Method,
		Enabled: msg.Enabled,
	})
}

func (k entitlementsMsgServer) SetRoleCapability(goCtx context.Context, msg *entitlements.MsgSetRoleCapability) (*entitlements.MsgSetRoleCapabilityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	_, roleExists := entitlements.Role_value[msg.Role.String()]
	if !roleExists {
		return nil, errors.Wrapf(entitlements.ErrInvalidRole, "role %s does not exist", msg.Role)
	}

	k.Keeper.SetRoleCapability(ctx, msg.Method, msg.Role, msg.Enabled)

	return &entitlements.MsgSetRoleCapabilityResponse{}, ctx.EventManager().EmitTypedEvent(&entitlements.RoleCapabilityUpdated{
		Role:    msg.Role,
		Method:  msg.Method,
		Enabled: msg.Enabled,
	})
}

func (k entitlementsMsgServer) SetUserRole(goCtx context.Context, msg *entitlements.MsgSetUserRole) (*entitlements.MsgSetUserRoleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	user, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode user address %s", msg.User)
	}

	_, roleExists := entitlements.Role_value[msg.Role.String()]
	if !roleExists {
		return nil, errors.Wrapf(entitlements.ErrInvalidRole, "role %s does not exist", msg.Role)
	}

	k.Keeper.SetUserRole(ctx, user, msg.Role, msg.Enabled)

	return &entitlements.MsgSetUserRoleResponse{}, ctx.EventManager().EmitTypedEvent(&entitlements.UserRoleUpdated{
		User:    msg.User,
		Role:    msg.Role,
		Enabled: msg.Enabled,
	})
}

func (k entitlementsMsgServer) Pause(goCtx context.Context, msg *entitlements.MsgPause) (*entitlements.MsgPauseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	k.SetPaused(ctx, true)

	return &entitlements.MsgPauseResponse{}, ctx.EventManager().EmitTypedEvent(&entitlements.Paused{
		Account: msg.Signer,
	})
}

func (k entitlementsMsgServer) Unpause(goCtx context.Context, msg *entitlements.MsgUnpause) (*entitlements.MsgUnpauseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	k.SetPaused(ctx, false)

	return &entitlements.MsgUnpauseResponse{}, ctx.EventManager().EmitTypedEvent(&entitlements.Unpaused{
		Account: msg.Signer,
	})
}

func (k entitlementsMsgServer) TransferOwnership(goCtx context.Context, msg *entitlements.MsgTransferOwnership) (*entitlements.MsgTransferOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	owner, err := k.EnsureOwner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	if msg.NewOwner == owner {
		return nil, entitlements.ErrSameOwner
	}

	k.SetEntitlementsOwner(ctx, msg.NewOwner)

	return &entitlements.MsgTransferOwnershipResponse{}, ctx.EventManager().EmitTypedEvent(&entitlements.OwnershipTransferred{
		PreviousOwner: owner,
		NewOwner:      msg.NewOwner,
	})
}

//

func (k entitlementsMsgServer) EnsureOwner(ctx sdk.Context, signer string) (string, error) {
	owner := k.GetEntitlementsOwner(ctx)
	if owner == "" {
		return "", entitlements.ErrNoOwner
	}
	if signer != owner {
		return "", errors.Wrapf(entitlements.ErrInvalidOwner, "expected %s, got %s", owner, signer)
	}
	return owner, nil
}
