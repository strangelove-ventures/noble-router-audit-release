package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/router module sentinel errors
var (
	ErrHandleMessage           = sdkerrors.Register(ModuleName, 2, "err during handle message")
)
