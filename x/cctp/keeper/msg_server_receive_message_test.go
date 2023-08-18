package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/sample"
	"github.com/strangelove-ventures/noble/x/cctp/keeper"
	"github.com/strangelove-ventures/noble/x/cctp/types"
	"github.com/stretchr/testify/require"
)

/*
 * Happy path
 * Happy path with destination caller
 * Sending and receiving messages paused
 * No attesters found
 * Signature threshold not found
 * Unable to verify signatures
 * Invalid message length
 * Incorrect destination domain
 * Incorrect destination caller
 * Invalid message version
 * Fails when nonce already used
 * Invalid message body version
 * Token pair not found
 */
func TestReceiveMessageHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	burnMessage := types.BurnMessage{
		Version:       0,
		BurnToken:     []byte("02345678901234567890123456789012"),
		MintRecipient: []byte("message sender567890123456789012"),
		Amount:        math.NewInt(9876),
		MessageSender: []byte("message sender567890123456789012"),
	}

	tokenPair := types.TokenPair{
		RemoteDomain: 0,
		RemoteToken:  burnMessage.BurnToken,
		LocalToken:   string(crypto.Keccak256([]byte("uusdc"))),
	}
	testkeeper.SetTokenPair(ctx, tokenPair)

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	message := types.Message{
		Version:           0,
		SourceDomain:      0,
		DestinationDomain: 4,
		Nonce:             0,
		Sender:            []byte("01234567890123456789012345678912"),
		Recipient:         crypto.Keccak256([]byte("cctp/TokenMessenger")),
		DestinationCaller: make([]byte, types.DestinationCallerLen),
		MessageBody:       burnMessageBytes,
	}
	messageBytes, err := message.Bytes()
	require.Nil(t, err)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	attestation := generateAttestation(messageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReceiveMessage{
		From:        "random address",
		Message:     messageBytes,
		Attestation: attestation,
	}

	resp, err := server.ReceiveMessage(sdk.WrapSDKContext(ctx), &msg)
	require.Nil(t, err)
	require.True(t, resp.Success)
}

func TestReceiveMessageWithDestinationCallerHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	burnMessage := types.BurnMessage{
		Version:       0,
		BurnToken:     []byte("02345678901234567890123456789012"),
		MintRecipient: []byte("message sender567890123456789012"),
		Amount:        math.NewInt(9876),
		MessageSender: []byte("message sender567890123456789012"),
	}

	tokenPair := types.TokenPair{
		RemoteDomain: 0,
		RemoteToken:  burnMessage.BurnToken,
		LocalToken:   string(crypto.Keccak256([]byte("uusdc"))),
	}
	testkeeper.SetTokenPair(ctx, tokenPair)

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	destinationCaller := sample.TestAccount()

	message := types.Message{
		Version:           0,
		SourceDomain:      0,
		DestinationDomain: 4,
		Nonce:             0,
		Sender:            []byte("01234567890123456789012345678912"),
		Recipient:         crypto.Keccak256([]byte("cctp/TokenMessenger")),
		DestinationCaller: append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, destinationCaller.AddressBz...),
		MessageBody:       burnMessageBytes,
	}
	messageBytes, err := message.Bytes()
	require.Nil(t, err)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	attestation := generateAttestation(messageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReceiveMessage{
		From:        destinationCaller.Address,
		Message:     messageBytes,
		Attestation: attestation,
	}

	resp, err := server.ReceiveMessage(sdk.WrapSDKContext(ctx), &msg)
	require.Nil(t, err)
	require.True(t, resp.Success)
}

func TestReceiveMessageSendingAndReceivingMessagesPaused(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: true}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	_, err := server.ReceiveMessage(sdk.WrapSDKContext(ctx), &types.MsgReceiveMessage{})
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "sending and receiving messages are paused")
}

func TestReceiveMessageNoAttestersFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	_, err := server.ReceiveMessage(sdk.WrapSDKContext(ctx), &types.MsgReceiveMessage{})
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "no attesters found")
}

func TestReceiveMessageSignatureThresholdNotFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}

	_, err := server.ReceiveMessage(sdk.WrapSDKContext(ctx), &types.MsgReceiveMessage{})
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "signature threshold not found")
}

func TestReceiveMessageUnableToVerifySignatures(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	_, err := server.ReceiveMessage(sdk.WrapSDKContext(ctx), &types.MsgReceiveMessage{})
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "unable to verify signatures")
}

func TestReceiveMessageInvalidMessageLength(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	messageBytes := []byte("too short")

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	attestation := generateAttestation(messageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReceiveMessage{
		From:        "random address",
		Message:     messageBytes,
		Attestation: attestation,
	}

	_, err := server.ReceiveMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrParsingMessage, err)
	require.Contains(t, err.Error(), "cctp message must be at least 116 bytes, got 9: error while parsing message into bytes")
}

func TestReceiveMessageIncorrectDestinationDomain(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	burnMessage := types.BurnMessage{
		Version:       0,
		BurnToken:     []byte("02345678901234567890123456789012"),
		MintRecipient: []byte("message sender567890123456789012"),
		Amount:        math.NewInt(9876),
		MessageSender: []byte("message sender567890123456789012"),
	}

	tokenPair := types.TokenPair{
		RemoteDomain: 0,
		RemoteToken:  burnMessage.BurnToken,
		LocalToken:   string(crypto.Keccak256([]byte("uusdc"))),
	}
	testkeeper.SetTokenPair(ctx, tokenPair)

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	message := types.Message{
		Version:           0,
		SourceDomain:      0,
		DestinationDomain: 11, // not noble
		Nonce:             0,
		Sender:            []byte("01234567890123456789012345678912"),
		Recipient:         crypto.Keccak256([]byte("cctp/TokenMessenger")),
		DestinationCaller: make([]byte, types.DestinationCallerLen),
		MessageBody:       burnMessageBytes,
	}
	messageBytes, err := message.Bytes()
	require.Nil(t, err)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	attestation := generateAttestation(messageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReceiveMessage{
		From:        "random address",
		Message:     messageBytes,
		Attestation: attestation,
	}

	_, err = server.ReceiveMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "incorrect destination domain")
}

func TestReceiveMessageIncorrectDestinationCaller(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	burnMessage := types.BurnMessage{
		Version:       0,
		BurnToken:     []byte("02345678901234567890123456789012"),
		MintRecipient: []byte("message sender567890123456789012"),
		Amount:        math.NewInt(9876),
		MessageSender: []byte("message sender567890123456789012"),
	}

	tokenPair := types.TokenPair{
		RemoteDomain: 0,
		RemoteToken:  burnMessage.BurnToken,
		LocalToken:   string(crypto.Keccak256([]byte("uusdc"))),
	}
	testkeeper.SetTokenPair(ctx, tokenPair)

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	message := types.Message{
		Version:           0,
		SourceDomain:      0,
		DestinationDomain: 4,
		Nonce:             0,
		Sender:            []byte("01234567890123456789012345678912"),
		Recipient:         crypto.Keccak256([]byte("cctp/TokenMessenger")),
		DestinationCaller: []byte("01234567890123456789012345678912"),
		MessageBody:       burnMessageBytes,
	}
	messageBytes, err := message.Bytes()
	require.Nil(t, err)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	attestation := generateAttestation(messageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReceiveMessage{
		From:        "not the destination caller",
		Message:     messageBytes,
		Attestation: attestation,
	}

	_, err = server.ReceiveMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "incorrect destination caller")
}

func TestReceiveMessageInvalidMessageVersion(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	burnMessage := types.BurnMessage{
		Version:       0,
		BurnToken:     []byte("02345678901234567890123456789012"),
		MintRecipient: []byte("message sender567890123456789012"),
		Amount:        math.NewInt(9876),
		MessageSender: []byte("message sender567890123456789012"),
	}

	tokenPair := types.TokenPair{
		RemoteDomain: 0,
		RemoteToken:  burnMessage.BurnToken,
		LocalToken:   string(crypto.Keccak256([]byte("uusdc"))),
	}
	testkeeper.SetTokenPair(ctx, tokenPair)

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	message := types.Message{
		Version:           4, // not the current version
		SourceDomain:      0,
		DestinationDomain: 4,
		Nonce:             0,
		Sender:            []byte("01234567890123456789012345678912"),
		Recipient:         crypto.Keccak256([]byte("cctp/TokenMessenger")),
		DestinationCaller: make([]byte, types.DestinationCallerLen),
		MessageBody:       burnMessageBytes,
	}
	messageBytes, err := message.Bytes()
	require.Nil(t, err)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	attestation := generateAttestation(messageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReceiveMessage{
		From:        "random address",
		Message:     messageBytes,
		Attestation: attestation,
	}

	_, err = server.ReceiveMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "incorrect message version")
}

func TestReceiveMessageNonceAlreadyUsed(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	burnMessage := types.BurnMessage{
		Version:       0,
		BurnToken:     []byte("02345678901234567890123456789012"),
		MintRecipient: []byte("message sender567890123456789012"),
		Amount:        math.NewInt(9876),
		MessageSender: []byte("message sender567890123456789012"),
	}

	tokenPair := types.TokenPair{
		RemoteDomain: 0,
		RemoteToken:  burnMessage.BurnToken,
		LocalToken:   string(crypto.Keccak256([]byte("uusdc"))),
	}
	testkeeper.SetTokenPair(ctx, tokenPair)

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	message := types.Message{
		Version:           0,
		SourceDomain:      5,
		DestinationDomain: 4,
		Nonce:             18,
		Sender:            []byte("01234567890123456789012345678912"),
		Recipient:         crypto.Keccak256([]byte("cctp/TokenMessenger")),
		DestinationCaller: make([]byte, types.DestinationCallerLen),
		MessageBody:       burnMessageBytes,
	}
	messageBytes, err := message.Bytes()
	require.Nil(t, err)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	attestation := generateAttestation(messageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	usedNonce := types.Nonce{
		SourceDomain: 5,
		Nonce:        18,
	}
	testkeeper.SetUsedNonce(ctx, usedNonce)

	msg := types.MsgReceiveMessage{
		From:        "random address",
		Message:     messageBytes,
		Attestation: attestation,
	}

	_, err = server.ReceiveMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "nonce already used")
}

func TestReceiveMessageInvalidMessageBodyVersion(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	burnMessage := types.BurnMessage{
		Version:       13, // not the current version
		BurnToken:     []byte("02345678901234567890123456789012"),
		MintRecipient: []byte("message sender567890123456789012"),
		Amount:        math.NewInt(9876),
		MessageSender: []byte("message sender567890123456789012"),
	}

	tokenPair := types.TokenPair{
		RemoteDomain: 0,
		RemoteToken:  burnMessage.BurnToken,
		LocalToken:   string(crypto.Keccak256([]byte("uusdc"))),
	}
	testkeeper.SetTokenPair(ctx, tokenPair)

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	message := types.Message{
		Version:           0,
		SourceDomain:      0,
		DestinationDomain: 4,
		Nonce:             5,
		Sender:            []byte("01234567890123456789012345678912"),
		Recipient:         crypto.Keccak256([]byte("cctp/TokenMessenger")),
		DestinationCaller: make([]byte, types.DestinationCallerLen),
		MessageBody:       burnMessageBytes,
	}
	messageBytes, err := message.Bytes()
	require.Nil(t, err)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	attestation := generateAttestation(messageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReceiveMessage{
		From:        "random address",
		Message:     messageBytes,
		Attestation: attestation,
	}

	_, err = server.ReceiveMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "invalid message body version")
}

func TestReceiveMessageTokenPairNotFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	burnMessage := types.BurnMessage{
		Version:       0,
		BurnToken:     []byte("02345678901234567890123456789012"),
		MintRecipient: []byte("message sender567890123456789012"),
		Amount:        math.NewInt(9876),
		MessageSender: []byte("message sender567890123456789012"),
	}

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	message := types.Message{
		Version:           0,
		SourceDomain:      0,
		DestinationDomain: 4,
		Nonce:             0,
		Sender:            []byte("01234567890123456789012345678912"),
		Recipient:         crypto.Keccak256([]byte("cctp/TokenMessenger")),
		DestinationCaller: make([]byte, types.DestinationCallerLen),
		MessageBody:       burnMessageBytes,
	}
	messageBytes, err := message.Bytes()
	require.Nil(t, err)

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	attestation := generateAttestation(messageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReceiveMessage{
		From:        "random address",
		Message:     messageBytes,
		Attestation: attestation,
	}

	_, err = server.ReceiveMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrReceiveMessage, err)
	require.Contains(t, err.Error(), "corresponding noble mint token not found")
}
