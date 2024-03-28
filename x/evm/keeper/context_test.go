package keeper_test

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"

	"github.com/initia-labs/minievm/x/evm/contracts/counter"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/contracts/i_cosmos"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"

	"github.com/stretchr/testify/require"
)

func Test_Create(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hex.DecodeString(strings.TrimPrefix(counter.CounterBin, "0x"))
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)
}

func Test_Call(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hex.DecodeString(strings.TrimPrefix(counter.CounterBin, "0x"))
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	queryInputBz, err := parsed.Pack("count")
	require.NoError(t, err)

	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(0).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	res, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz)
	require.NoError(t, err)
	require.Empty(t, res)
	require.NotEmpty(t, logs)

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	// calling not existing function
	erc20ABI, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	queryInputBz, err = erc20ABI.Pack("balanceOf", caller)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz)
	require.ErrorContains(t, err, vm.ErrExecutionReverted.Error())
}

func Test_ExecuteCosmosMessage(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// mint native coin
	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	require.NoError(t, err)

	abi, err := i_cosmos.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("execute_cosmos_message", fmt.Sprintf(`
		{
			"@type": "/cosmos.bank.v1beta1.MsgSend",
			"from_address": "%s",
			"to_address": "%s",
			"amount": [
				{
					"denom": "bar",
					"amount": "100"
				}
			]
		}
	`, addr, addr2))
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, evmAddr, types.CosmosPrecompileAddress, inputBz)
	require.NoError(t, err)

	balance := input.BankKeeper.GetBalance(ctx, addr2, "bar")
	require.Equal(t, math.NewInt(100), balance.Amount)
}
