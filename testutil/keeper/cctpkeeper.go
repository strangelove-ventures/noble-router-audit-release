package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/strangelove-ventures/noble/x/cctp/keeper"
	"github.com/strangelove-ventures/noble/x/cctp/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func CctpKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		codec.NewLegacyAmino(),
		storeKey,
		nil,
		"CctpParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		paramsSubspace,
		MockBankKeeper{},
		MockFiatTokenfactoryKeeper{},
		MockRouterKeeper{},
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, ctx
}

type MockCctpKeeper struct{}

func (k MockCctpKeeper) GetTokenPair(ctx sdk.Context, remoteDomain uint32, remoteToken []byte) (val types.TokenPair, found bool) {
	return types.TokenPair{}, true
}

func (MockCctpKeeper) GetAuthority(ctx sdk.Context) (val string, found bool) {
	return "", true
}

// ErrCctpKeeper is used for wrapping a MockErrFiatTokenfactoryKeeper, which fails on mint/burn
func ErrCctpKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		codec.NewLegacyAmino(),
		storeKey,
		nil,
		"CctpParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		paramsSubspace,
		MockBankKeeper{},
		MockErrFiatTokenfactoryKeeper{},
		MockRouterKeeper{},
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, ctx
}

type MockErrCctpKeeper struct{}
