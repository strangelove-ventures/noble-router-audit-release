package types

const (
	// Indices of each field in message
	VersionIndex           = 0
	SourceDomainIndex      = 4
	DestinationDomainIndex = 8
	NonceIndex             = 12
	SenderIndex            = 20
	RecipientIndex         = 52
	DestinationCallerIndex = 84
	MessageBodyIndex       = 116

	// Indices of each field in BurnMessage
	BurnMsgVersionIndex = 0
	VersionLen          = 4
	BurnTokenIndex      = 4
	BurnTokenLen        = 32
	MintRecipientIndex  = 36
	MintRecipientLen    = 32
	AmountIndex         = 68
	AmountLen           = 32
	MsgSenderIndex      = 100
	MsgSenderLen        = 32
	// 4 byte version + 32 bytes burnToken + 32 bytes mintRecipient + 32 bytes amount + 32 bytes messageSender
	BurnMessageLen = 132

	NobleMessageVersion = 0
	MessageBodyVersion  = 0
	NobleDomainId       = 4

	DomainBytesLen  = 4
	UsedNonceLen    = 8
	NonceBytesLen   = 8
	AddressBytesLen = 32

	DomainBitLen                = 32
	NonceBitLen                 = 32
	DestinationCallerLen        = 32
	SignatureThresholdBitLength = 32
	BaseTen                     = 10

	SignatureLength = 65
)
