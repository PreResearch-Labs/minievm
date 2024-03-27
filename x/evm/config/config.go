package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

// DefaultContractQueryGasLimit - default max query gas for external query
const DefaultContractQueryGasLimit = uint64(3_000_000)

// DefaultContractSimulationGasLimit - default max simulation gas
const DefaultContractSimulationGasLimit = uint64(3_000_000)

const (
	flagContractSimulationGasLimit = "evm.contract-simulation-gas-limit"
	flagContractQueryGasLimit      = "evm.contract-query-gas-limit"
	flagTracingEnableMemory        = "evm.tracing-enable-memory"
	flagTracingEnableStack         = "evm.tracing-enable-stack"
	flagTracingEnableStorage       = "evm.tracing-enable-storage"
	flagTracingEnableReturnData    = "evm.tracing-enable-return-data"
)

// EVMConfig is the extra config required for evm
type EVMConfig struct {
	ContractSimulationGasLimit uint64 `mapstructure:"contract-simulation-gas-limit"`
	ContractQueryGasLimit      uint64 `mapstructure:"contract-query-gas-limit"`
	TracingEnableMemory        bool   `mapstructure:"tracing-enable-memory"`
	TracingEnableStack         bool   `mapstructure:"tracing-enable-stack"`
	TracingEnableStorage       bool   `mapstructure:"tracing-enable-storage"`
	TracingEnableReturnData    bool   `mapstructure:"tracing-enable-return-data"`
}

// DefaultEVMConfig returns the default settings for EVMConfig
func DefaultEVMConfig() EVMConfig {
	return EVMConfig{
		ContractSimulationGasLimit: DefaultContractSimulationGasLimit,
		ContractQueryGasLimit:      DefaultContractQueryGasLimit,
		TracingEnableMemory:        false,
		TracingEnableStack:         false,
		TracingEnableStorage:       false,
		TracingEnableReturnData:    false,
	}
}

// GetConfig load config values from the app options
func GetConfig(appOpts servertypes.AppOptions) EVMConfig {
	return EVMConfig{
		ContractSimulationGasLimit: cast.ToUint64(appOpts.Get(flagContractSimulationGasLimit)),
		ContractQueryGasLimit:      cast.ToUint64(appOpts.Get(flagContractQueryGasLimit)),
		TracingEnableMemory:        cast.ToBool(appOpts.Get(flagTracingEnableMemory)),
		TracingEnableStack:         cast.ToBool(appOpts.Get(flagTracingEnableStack)),
		TracingEnableStorage:       cast.ToBool(appOpts.Get(flagTracingEnableStorage)),
		TracingEnableReturnData:    cast.ToBool(appOpts.Get(flagTracingEnableReturnData)),
	}
}

// AddConfigFlags implements servertypes.EVMConfigFlags interface.
func AddConfigFlags(startCmd *cobra.Command) {
	startCmd.Flags().Uint64(flagContractSimulationGasLimit, DefaultContractSimulationGasLimit, "Set the max simulation gas for evm contract execution")
	startCmd.Flags().Uint64(flagContractQueryGasLimit, DefaultContractQueryGasLimit, "Set the max gas that can be spent on executing a query with a Move contract")
	startCmd.Flags().Bool(flagTracingEnableMemory, false, "Enable memory tracing at query/Call")
	startCmd.Flags().Bool(flagTracingEnableStack, false, "Enable stack tracing at query/Call")
	startCmd.Flags().Bool(flagTracingEnableStorage, false, "Enable storage tracing at query/Call")
	startCmd.Flags().Bool(flagTracingEnableReturnData, false, "Enable return data tracing at query/Call")
}

// DefaultConfigTemplate default config template for evm
const DefaultConfigTemplate = `
###############################################################################
###                         EVM                                             ###
###############################################################################

[evm]

# The maximum gas amount can be used in a tx simulation call.
contract-simulation-gas-limit = "{{ .EVMConfig.ContractSimulationGasLimit }}"

# The maximum gas amount can be spent for contract query.
# The contract query will invoke contract execution vm,
# so we need to restrict the max usage to prevent DoS attack
contract-query-gas-limit = "{{ .EVMConfig.ContractQueryGasLimit }}"

# Enable memory tracing at query/Call
tracing-enable-memory = {{ .EVMConfig.TracingEnableMemory }}

# Enable stack tracing at query/Call
tracing-enable-stack = {{ .EVMConfig.TracingEnableStack }}

# Enable storage tracing at query/Call
tracing-enable-storage = {{ .EVMConfig.TracingEnableStorage }}

# Enable return data tracing at query/Call
tracing-enable-return-data = {{ .EVMConfig.TracingEnableReturnData }}
`
