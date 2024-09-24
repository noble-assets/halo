package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/halo/v2/types/entitlements"
)

var _ entitlements.QueryServer = &entitlementsQueryServer{}

type entitlementsQueryServer struct {
	*Keeper
}

func NewEntitlementsQueryServer(keeper *Keeper) entitlements.QueryServer {
	return &entitlementsQueryServer{Keeper: keeper}
}

func (k entitlementsQueryServer) Owner(ctx context.Context, req *entitlements.QueryOwner) (*entitlements.QueryOwnerResponse, error) {
	if req == nil {
		return nil, errorstypes.ErrInvalidRequest
	}

	return &entitlements.QueryOwnerResponse{
		Owner: k.GetEntitlementsOwner(ctx),
	}, nil
}

func (k entitlementsQueryServer) Paused(ctx context.Context, req *entitlements.QueryPaused) (*entitlements.QueryPausedResponse, error) {
	if req == nil {
		return nil, errorstypes.ErrInvalidRequest
	}

	return &entitlements.QueryPausedResponse{
		Paused: k.GetPaused(ctx),
	}, nil
}

func (k entitlementsQueryServer) PublicCapability(ctx context.Context, req *entitlements.QueryPublicCapability) (*entitlements.QueryPublicCapabilityResponse, error) {
	if req == nil || req.Method == "" {
		return nil, errorstypes.ErrInvalidRequest
	}

	return &entitlements.QueryPublicCapabilityResponse{
		Enabled: k.IsPublicCapability(ctx, req.Method),
	}, nil
}

func (k entitlementsQueryServer) RoleCapability(ctx context.Context, req *entitlements.QueryRoleCapability) (*entitlements.QueryRoleCapabilityResponse, error) {
	if req == nil || req.Method == "" {
		return nil, errorstypes.ErrInvalidRequest
	}

	return &entitlements.QueryRoleCapabilityResponse{
		Roles: k.GetCapabilityRoles(ctx, req.Method),
	}, nil
}

func (k entitlementsQueryServer) UserCapability(ctx context.Context, req *entitlements.QueryUserCapability) (*entitlements.QueryUserCapabilityResponse, error) {
	if req == nil {
		return nil, errorstypes.ErrInvalidRequest
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode address %s", req.Address)
	}

	return &entitlements.QueryUserCapabilityResponse{
		Roles: k.GetUserRoles(ctx, address),
	}, nil
}
