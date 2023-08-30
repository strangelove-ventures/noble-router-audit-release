package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/router module sentinel errors
var (
	ErrHandleMessage               = sdkerrors.Register(ModuleName, 2, "err during handle message")
	ErrDecodingMessage             = sdkerrors.Register(ModuleName, 3, "err decoding message")
	ErrDecodingBurnMessage         = sdkerrors.Register(ModuleName, 4, "err decoding burn message")
	ErrDecodingIBCForward          = sdkerrors.Register(ModuleName, 5, "err decoding ibc forward")
	ErrInvalidMint                 = sdkerrors.Register(ModuleName, 6, "validation failed due to malformed Mint")
	ErrInvalidInFlightPacket       = sdkerrors.Register(ModuleName, 7, "validation failed due to malformed InFlightPacket")
	ErrInvalidStoreForwardMetadata = sdkerrors.Register(ModuleName, 8, "validation failed due to malformed StoreIBCForwardMetadata")
)
