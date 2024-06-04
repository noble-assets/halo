package keeper

import (
	"context"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/noble-assets/halo/x/halo/types"
)

var _ types.QueryServer = &queryServer{}

type queryServer struct {
	*Keeper
}

func NewQueryServer(keeper *Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

func (k queryServer) Owner(goCtx context.Context, req *types.QueryOwner) (*types.QueryOwnerResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryOwnerResponse{
		Owner: k.GetOwner(ctx),
	}, nil
}

func (k queryServer) Nonces(goCtx context.Context, req *types.QueryNonces) (*types.QueryNoncesResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.NoncePrefix)

	nonces := make(map[string]uint64)
	pagination, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		address := sdk.AccAddress(key).String()
		nonces[address] = binary.BigEndian.Uint64(value)
		return nil
	})

	return &types.QueryNoncesResponse{
		Nonces:     nonces,
		Pagination: pagination,
	}, err
}

func (k queryServer) Nonce(goCtx context.Context, req *types.QueryNonce) (*types.QueryNonceResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode address %s", req.Address)
	}

	return &types.QueryNonceResponse{
		Nonce: k.GetNonce(ctx, address),
	}, nil
}
