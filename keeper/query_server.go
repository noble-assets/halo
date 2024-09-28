package keeper

import (
	"context"

	"cosmossdk.io/errors"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/noble-assets/halo/v2/types"
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

	nonces := make(map[string]uint64)
	_, pagination, err := query.CollectionPaginate(ctx, k.Keeper.Nonces, req.Pagination, func(key []byte, nonce uint64) (string, error) {
		address, err := k.addressCodec.BytesToString(key)
		if err != nil {
			return "", err
		}

		nonces[address] = nonce
		return "", nil
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

	address, err := k.addressCodec.StringToBytes(req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode address %s", req.Address)
	}

	return &types.QueryNonceResponse{
		Nonce: k.GetNonce(ctx, address),
	}, nil
}
