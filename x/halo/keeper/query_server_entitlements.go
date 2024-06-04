package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
)

var _ entitlements.QueryServer = &entitlementsQueryServer{}

type entitlementsQueryServer struct {
	*Keeper
}

func NewEntitlementsQueryServer(keeper *Keeper) entitlements.QueryServer {
	return &entitlementsQueryServer{Keeper: keeper}
}

func (k entitlementsQueryServer) Owner(goCtx context.Context, req *entitlements.QueryOwner) (*entitlements.QueryOwnerResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &entitlements.QueryOwnerResponse{
		Owner: k.GetEntitlementsOwner(ctx),
	}, nil
}

func (k entitlementsQueryServer) Paused(goCtx context.Context, req *entitlements.QueryPaused) (*entitlements.QueryPausedResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &entitlements.QueryPausedResponse{
		Paused: k.GetPaused(ctx),
	}, nil
}

func (k entitlementsQueryServer) PublicCapabilities(goCtx context.Context, req *entitlements.QueryPublicCapabilities) (*entitlements.QueryPublicCapabilitiesResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), entitlements.PublicPrefix)

	publicCapabilities := make(map[string]bool)
	pagination, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		if len(value) == 1 && value[0] == 1 {
			publicCapabilities[string(key)] = true
		} else {
			publicCapabilities[string(key)] = false
		}
		return nil
	})

	return &entitlements.QueryPublicCapabilitiesResponse{
		PublicCapabilities: publicCapabilities,
		Pagination:         pagination,
	}, err
}

func (k entitlementsQueryServer) PublicCapability(goCtx context.Context, req *entitlements.QueryPublicCapability) (*entitlements.QueryPublicCapabilityResponse, error) {
	if req == nil || req.Method == "" {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &entitlements.QueryPublicCapabilityResponse{
		Enabled: k.IsPublicCapability(ctx, req.Method),
	}, nil
}
