package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble/x/cctp/types"
)

// DeletePendingOwner deletes the pending owner of the router module from state.
func (k *Keeper) DeletePendingOwner(ctx sdk.Context) {
	ctx.KVStore(k.storeKey).Delete(types.PendingOwnerKey)
}

// GetOwner returns the owner of the router module from state.
func (k *Keeper) GetOwner(ctx sdk.Context) (owner string) {
	bz := ctx.KVStore(k.storeKey).Get(types.OwnerKey)
	if bz == nil {
		panic("cctp owner not found in state")
	}

	return string(bz)
}

// GetPendingOwner returns the pending owner of the router module from state.
func (k *Keeper) GetPendingOwner(ctx sdk.Context) (pendingOwner string, found bool) {
	bz := ctx.KVStore(k.storeKey).Get(types.PendingOwnerKey)
	if bz == nil {
		return "", false
	}

	return string(bz), true
}

// SetOwner stores the owner of the router module in state.
func (k *Keeper) SetOwner(ctx sdk.Context, owner string) {
	bz := []byte(owner)
	ctx.KVStore(k.storeKey).Set(types.OwnerKey, bz)
}

// SetPendingOwner stores the pending owner of the router module in state.
func (k *Keeper) SetPendingOwner(ctx sdk.Context, pendingOwner string) {
	bz := []byte(pendingOwner)
	ctx.KVStore(k.storeKey).Set(types.PendingOwnerKey, bz)
}
