package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/strangelove-ventures/noble/x/router/types"
)

// msgServerRouterKeeper defines the router keeper methods required by the query server.
type queryServerRouterKeeper interface {
	GetParams(ctx sdk.Context) types.Params

	IsAllowedSourceDomainSender(ctx sdk.Context, domainID uint32, address []byte) bool
	GetAllAllowedSourceDomainSendersPaginated(ctx sdk.Context, pagination *query.PageRequest) ([]types.AllowedSourceDomainSender, *query.PageResponse, error)

	GetIBCForward(ctx sdk.Context, sourceDomain uint32, nonce uint64) (types.StoreIBCForwardMetadata, bool)
	GetAllIBCForwardsPaginated(ctx sdk.Context, pagination *query.PageRequest) ([]types.StoreIBCForwardMetadata, *query.PageResponse, error)

	GetInFlightPacket(ctx sdk.Context, channelID string, portID string, sequence uint64) (types.InFlightPacket, bool)
	GetAllInFlightPacketsPaginated(ctx sdk.Context, pagination *query.PageRequest) ([]types.InFlightPacket, *query.PageResponse, error)

	GetMint(ctx sdk.Context, sourceDomain uint32, nonce uint64) (types.Mint, bool)
	GetAllMintsPaginated(ctx sdk.Context, pagination *query.PageRequest) ([]types.Mint, *query.PageResponse, error)
}

var _ queryServerRouterKeeper = &Keeper{}

type QueryServer struct {
	keeper queryServerRouterKeeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
func NewQueryServer(keeper queryServerRouterKeeper) QueryServer {
	return QueryServer{keeper: keeper}
}

var _ types.QueryServer = QueryServer{}
