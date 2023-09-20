package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/strangelove-ventures/noble/x/router/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IsAllowedSourceDomainSender returns true if the source domain sender is allowed
func (k *Keeper) IsAllowedSourceDomainSender(ctx sdk.Context, domainID uint32, address []byte) (allowed bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AllowedSourceDomainSenderKeyPrefix)

	b := store.Get(types.SourceDomainSenderKey(domainID, address))
	return b != nil
}

// AddAllowedSourceDomainSender adds an allowed source domain sender
func (k *Keeper) AddAllowedSourceDomainSender(ctx sdk.Context, domainID uint32, address []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AllowedSourceDomainSenderKeyPrefix)
	store.Set(types.SourceDomainSenderKey(domainID, address), []byte{})
}

// DeleteAllowedSourceDomainSender removes an allowed source domain sender
func (k *Keeper) DeleteAllowedSourceDomainSender(ctx sdk.Context, domainID uint32, address []byte,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AllowedSourceDomainSenderKeyPrefix)
	store.Delete(types.SourceDomainSenderKey(domainID, address))
}

// GetAllowedSourceDomainSenders returns all allowed source domain senders
func (k *Keeper) GetAllowedSourceDomainSenders(ctx sdk.Context) (list []types.AllowedSourceDomainSender) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AllowedSourceDomainSenderKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		val := types.AllowedSourceDomainSender{
			DomainId: binary.BigEndian.Uint32(key[0:4]),
			Address:  key[4:36],
		}
		list = append(list, val)
	}

	return
}

func (k *Keeper) GetAllAllowedSourceDomainSendersPaginated(ctx sdk.Context, pagination *query.PageRequest) ([]types.AllowedSourceDomainSender, *query.PageResponse, error) {
	var allowedSourceDomainSenders []types.AllowedSourceDomainSender

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AllowedSourceDomainSenderKeyPrefix)

	pageRes, err := query.Paginate(store, pagination, func(key []byte, _ []byte) error {
		val := types.AllowedSourceDomainSender{
			DomainId: binary.BigEndian.Uint32(key[0:4]),
			Address:  key[4:36],
		}
		allowedSourceDomainSenders = append(allowedSourceDomainSenders, val)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return allowedSourceDomainSenders, pageRes, nil
}
