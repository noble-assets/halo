package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k queryServer) Owner(ctx context.Context, req *types.QueryOwner) (*types.QueryOwnerResponse, error) {
	if req == nil {
		return nil, errorstypes.ErrInvalidRequest
	}

	return &types.QueryOwnerResponse{
		Owner: k.GetOwner(ctx),
	}, nil
}

func (k queryServer) Nonces(ctx context.Context, req *types.QueryNonces) (*types.QueryNoncesResponse, error) {
	if req == nil {
		return nil, errorstypes.ErrInvalidRequest
	}

	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(adapter, types.NoncePrefix)

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

func (k queryServer) Nonce(ctx context.Context, req *types.QueryNonce) (*types.QueryNonceResponse, error) {
	if req == nil {
		return nil, errorstypes.ErrInvalidRequest
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode address %s", req.Address)
	}

	return &types.QueryNonceResponse{
		Nonce: k.GetNonce(ctx, address),
	}, nil
}
