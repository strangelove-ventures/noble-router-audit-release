package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/strangelove-ventures/noble/x/router/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetIBCForward sets a IBCForward in the store
func (k *Keeper) SetIBCForward(ctx sdk.Context, forward types.StoreIBCForwardMetadata) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCForwardPrefix)
	b := k.cdc.MustMarshal(&forward)
	store.Set(types.LookupKey(forward.SourceDomain, forward.Metadata.Nonce), b)
}

// GetIBCForward returns IBCForward
func (k *Keeper) GetIBCForward(ctx sdk.Context, sourceDomain uint32, nonce uint64) (val types.StoreIBCForwardMetadata, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCForwardPrefix)

	b := store.Get(types.LookupKey(sourceDomain, nonce))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DeleteIBCForward removes a IBCForward from the store
func (k *Keeper) DeleteIBCForward(ctx sdk.Context, sourceDomain uint32, nonce uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCForwardPrefix)
	store.Delete(types.LookupKey(sourceDomain, nonce))
}

// GetAllIBCForwards returns all IBCForwards
func (k *Keeper) GetAllIBCForwards(ctx sdk.Context) (list []types.StoreIBCForwardMetadata) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCForwardPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoreIBCForwardMetadata
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k *Keeper) GetAllIBCForwardsPaginated(ctx sdk.Context, pagination *query.PageRequest) ([]types.StoreIBCForwardMetadata, *query.PageResponse, error) {
	var ibcForwards []types.StoreIBCForwardMetadata

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCForwardPrefix)

	pageRes, err := query.Paginate(store, pagination, func(key []byte, value []byte) error {
		var ibcForward types.StoreIBCForwardMetadata
		if err := k.cdc.Unmarshal(value, &ibcForward); err != nil {
			return err
		}

		ibcForwards = append(ibcForwards, ibcForward)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return ibcForwards, pageRes, nil
}
