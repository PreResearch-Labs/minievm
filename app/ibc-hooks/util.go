package evm_hooks

import (
	"encoding/json"
	"fmt"
	"strings"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

const senderPrefix = "ibc-evm-hook-intermediary"

// deriveIntermediateSender compute intermediate sender address
// Bech32(Hash(Hash("ibc-hook-intermediary") + channelID/sender)[12:])
//
// @dev: use 20bytes of address as intermediate sender address
//
// TODO - make this as module account to check address collision
func deriveIntermediateSender(channel, originalSender string) string {
	senderStr := fmt.Sprintf("%s/%s", channel, originalSender)
	senderAddr := sdk.AccAddress(address.Hash(senderPrefix, []byte(senderStr))[12:])
	return senderAddr.String()
}

func isIcs20Packet(packetData []byte) (isIcs20 bool, ics20data transfertypes.FungibleTokenPacketData) {
	var data transfertypes.FungibleTokenPacketData
	decoder := json.NewDecoder(strings.NewReader(string(packetData)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&data); err != nil {
		return false, data
	}
	return true, data
}

// func isIcs721Packet(packetData []byte) (isIcs721 bool, ics721data nfttransfertypes.NonFungibleTokenPacketData) {
// 	// Use evm port prefix to ack like normal evm chain.
// 	//
// 	// initia l1 is handling encoding and decoding depends on port id,
// 	// so minievm should ack like normal evm chain.
// 	if data, err := nfttransfertypes.DecodePacketData(packetData, evmPortPrefix); err != nil {
// 		return false, data
// 	} else {
// 		return true, data
// 	}
// }

func validateAndParseMemo(memo string) (
	isEVMRouted bool,
	hookData HookData,
	err error,
) {
	isEVMRouted, metadata := jsonStringHasKey(memo, evmHookMemoKey)
	if !isEVMRouted {
		return
	}

	evmHookRaw := metadata[evmHookMemoKey]

	// parse evm raw bytes to execute message
	bz, err := json.Marshal(evmHookRaw)
	if err != nil {
		err = errors.Wrap(channeltypes.ErrInvalidPacket, err.Error())
		return
	}

	err = json.Unmarshal(bz, &hookData)
	if err != nil {
		err = errors.Wrap(channeltypes.ErrInvalidPacket, err.Error())
		return
	}

	return
}

func validateReceiver(msg *evmtypes.MsgCall, receiver string) error {
	if receiver != msg.ContractAddr {
		return errors.Wrapf(channeltypes.ErrInvalidPacket, "receiver is not properly set")
	}

	return nil
}

// jsonStringHasKey parses the memo as a json object and checks if it contains the key.
func jsonStringHasKey(memo, key string) (found bool, jsonObject map[string]interface{}) {
	jsonObject = make(map[string]interface{})

	// If there is no memo, the packet was either sent with an earlier version of IBC, or the memo was
	// intentionally left blank. Nothing to do here. Ignore the packet and pass it down the stack.
	if len(memo) == 0 {
		return false, jsonObject
	}

	// the jsonObject must be a valid JSON object
	err := json.Unmarshal([]byte(memo), &jsonObject)
	if err != nil {
		return false, jsonObject
	}

	// If the key doesn't exist, there's nothing to do on this hook. Continue by passing the packet
	// down the stack
	_, ok := jsonObject[key]
	if !ok {
		return false, jsonObject
	}

	return true, jsonObject
}

// newEmitErrorAcknowledgement creates a new error acknowledgement after having emitted an event with the
// details of the error.
func newEmitErrorAcknowledgement(err error) channeltypes.Acknowledgement {
	return channeltypes.Acknowledgement{
		Response: &channeltypes.Acknowledgement_Error{
			Error: fmt.Sprintf("ibc evm hook error: %s", err.Error()),
		},
	}
}

// isAckError checks an IBC acknowledgement to see if it's an error.
func isAckError(appCodec codec.Codec, acknowledgement []byte) bool {
	var ack channeltypes.Acknowledgement
	if err := appCodec.UnmarshalJSON(acknowledgement, &ack); err == nil && !ack.Success() {
		return true
	}

	return false
}
