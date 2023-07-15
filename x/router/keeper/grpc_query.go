package keeper

import (
	"github.com/strangelove-ventures/noble-router/x/router/types"
)

var _ types.QueryServer = Keeper{}
