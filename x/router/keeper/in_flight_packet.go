package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/strangelove-ventures/noble/x/router/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetInFlightPacket sets a InFlightPacket in the store
func (k *Keeper) SetInFlightPacket(ctx sdk.Context, ifp types.InFlightPacket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InFlightPacketPrefix)
	b := k.cdc.MustMarshal(&ifp)
	store.Set(types.InFlightPacketKey(ifp.Channel, ifp.Port, ifp.Sequence), b)
}

// GetInFlightPacket returns InFlightPacket
func (k *Keeper) GetInFlightPacket(ctx sdk.Context, channelID string, portID string, sequence uint64) (val types.InFlightPacket, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InFlightPacketPrefix)

	b := store.Get(types.InFlightPacketKey(channelID, portID, sequence))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DeleteInFlightPacket removes a InFlightPacket from the store
func (k *Keeper) DeleteInFlightPacket(ctx sdk.Context, channelID string, portID string, sequence uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InFlightPacketPrefix)
	store.Delete(types.InFlightPacketKey(channelID, portID, sequence))
}

// GetAllInFlightPackets returns all InFlightPackets
func (k *Keeper) GetAllInFlightPackets(ctx sdk.Context) (list []types.InFlightPacket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InFlightPacketPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InFlightPacket
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k *Keeper) GetAllInFlightPacketsPaginated(ctx sdk.Context, pagination *query.PageRequest) ([]types.InFlightPacket, *query.PageResponse, error) {
	var inFlightPackets []types.InFlightPacket

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InFlightPacketPrefix)

	pageRes, err := query.Paginate(store, pagination, func(key []byte, value []byte) error {
		var InFlightPacket types.InFlightPacket
		if err := k.cdc.Unmarshal(value, &InFlightPacket); err != nil {
			return err
		}

		inFlightPackets = append(inFlightPackets, InFlightPacket)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return inFlightPackets, pageRes, nil
}
