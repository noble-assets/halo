package keeper

import "context"

//

func (k *Keeper) GetOwner(ctx context.Context) string {
	owner, _ := k.Owner.Get(ctx)
	return owner
}

func (k *Keeper) SetOwner(ctx context.Context, owner string) {
	_ = k.Owner.Set(ctx, owner)
}

//

func (k *Keeper) GetNonce(ctx context.Context, address []byte) uint64 {
	nonce, _ := k.Nonces.Get(ctx, address)
	return nonce
}

func (k *Keeper) GetNonces(ctx context.Context) map[string]uint64 {
	nonces := make(map[string]uint64)

	_ = k.Nonces.Walk(ctx, nil, func(bz []byte, nonce uint64) (stop bool, err error) {
		address, _ := k.accountKeeper.AddressCodec().BytesToString(bz)

		nonces[address] = nonce
		return false, nil
	})

	return nonces
}

func (k *Keeper) IncrementNonce(ctx context.Context, address []byte) uint64 {
	nonce := k.GetNonce(ctx, address)
	k.SetNonce(ctx, address, nonce+1)
	return nonce
}

func (k *Keeper) SetNonce(ctx context.Context, address []byte, nonce uint64) {
	_ = k.Nonces.Set(ctx, address, nonce)
}
