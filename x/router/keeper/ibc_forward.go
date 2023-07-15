package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/strangelove-ventures/noble-router/x/router/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetIBCForward sets a IBCForward in the store
func (k Keeper) SetIBCForward(ctx sdk.Context, key types.StoreIBCForwardMetadata) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&key)
	store.Set(types.LookupKey(key.SourceDomainSender, key.Nonce), b)
}

// GetIBCForward returns IBCForward
func (k Keeper) GetIBCForward(ctx sdk.Context, sourceContractAddress string, nonce uint64) (val types.StoreIBCForwardMetadata, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCForwardPrefix(types.IBCForwardKeyPrefix))

	b := store.Get(types.LookupKey(sourceContractAddress, nonce))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DeleteIBCForward removes a IBCForward from the store
func (k Keeper) DeleteIBCForward(ctx sdk.Context, sourceContractAddress string, nonce uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCForwardPrefix(types.IBCForwardKeyPrefix))
	store.Delete(types.LookupKey(sourceContractAddress, nonce))
}

// GetAllIBCForwards returns all IBCForwards
func (k Keeper) GetAllIBCForwards(ctx sdk.Context) (list []types.StoreIBCForwardMetadata) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCForwardPrefix(types.IBCForwardKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoreIBCForwardMetadata
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
