// Copyright 2020 Findora, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/findoranetwork/rosetta-sdk-go/types"
)

const (
	// NodeVersion is the version of findora we are using.
	NodeVersion = "1.9.24"

	// Blockchain is Findora.
	Blockchain string = "Findora"

	// MainnetNetwork is the value of the network
	// in MainnetNetworkIdentifier.
	MainnetNetwork string = "Mainnet"

	// AnvilNetwork is the value of the network
	// in AnvilNetworkIdentifier.
	AnvilNetwork string = "Anvil"

	// Qa02Network is the value of the network
	// in Qa02NetworkIdentifier.
	Qa02Network string = "Qa02"

	// PrinetNetwork is the value of the network
	// in PrinetNetworkNetworkIdentifier.
	PrinetNetwork string = "Prinet"

	// Symbol is the symbol value
	// used in Currency.
	Symbol = "FRA"

	// Decimals is the decimals value
	// used in Currency.
	Decimals = 18

	// MinerRewardOpType is used to describe
	// a miner block reward.
	MinerRewardOpType = "MINER_REWARD"

	// UncleRewardOpType is used to describe
	// an uncle block reward.
	UncleRewardOpType = "UNCLE_REWARD"

	// FeeOpType is used to represent fee operations.
	FeeOpType = "FEE"

	// CallOpType is used to represent CALL trace operations.
	CallOpType = "CALL"

	// CreateOpType is used to represent CREATE trace operations.
	CreateOpType = "CREATE"

	// Create2OpType is used to represent CREATE2 trace operations.
	Create2OpType = "CREATE2"

	// SelfDestructOpType is used to represent SELFDESTRUCT trace operations.
	SelfDestructOpType = "SELFDESTRUCT"

	// CallCodeOpType is used to represent CALLCODE trace operations.
	CallCodeOpType = "CALLCODE"

	// DelegateCallOpType is used to represent DELEGATECALL trace operations.
	DelegateCallOpType = "DELEGATECALL"

	// StaticCallOpType is used to represent STATICCALL trace operations.
	StaticCallOpType = "STATICCALL"

	// DestructOpType is a synthetic operation used to represent the
	// deletion of suicided accounts that still have funds at the end
	// of a transaction.
	DestructOpType = "DESTRUCT"

	// SuccessStatus is the status of any
	// Findora operation considered successful.
	SuccessStatus = "SUCCESS"

	// FailureStatus is the status of any
	// Findora operation considered unsuccessful.
	FailureStatus = "FAILURE"

	// HistoricalBalanceSupported is whether
	// historical balance is supported.
	HistoricalBalanceSupported = true

	// UnclesRewardMultiplier is the uncle reward
	// multiplier.
	UnclesRewardMultiplier = 32

	// MaxUncleDepth is the maximum depth for
	// an uncle to be rewarded.
	MaxUncleDepth = 8

	// GenesisBlockIndex is the index of the
	// genesis block.
	GenesisBlockIndex = int64(0)

	// TransferGasLimit is the gas limit
	// of a transfer.
	TransferGasLimit = int64(21000) //nolint:gomnd

	// IncludeMempoolCoins does not apply to findora-rosetta as it is not UTXO-based.
	IncludeMempoolCoins = false
)

var (
	// MainnetCommandArguments are the arguments to start a mainnet findroa instance.
	MainnetCommandArguments = ""

	// AnvilCommandArguments are the arguments to start a anvil findroa instance.
	AnvilCommandArguments = ""

	// Qa02CommandArguments are the arguments to start a qa02 findroa instance.
	Qa02CommandArguments = ""

	// PrinetCommandArguments are the arguments to start a prinet findroa instance.
	PrinetCommandArguments = ""

	MainnetGenesisHash = common.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")
	AnvilGenesisHash   = common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d")
	Qa02GenesisHash    = common.HexToHash("0xbf7e331f7f7c1dd2e05159666b3bf8bc7a8a3a9eb1d518969eab529dd9b88c1a")
	PrinetGenesisHash  = common.HexToHash("0xbf7e331f7f7c1dd2e05159666b3bf8bc7a8a3a9eb1d518969eab529dd9b88c1a")

	// MainnetGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the mainnet genesis block.
	MainnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  MainnetGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// AnvilGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the Anvil genesis block.
	AnvilGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  AnvilGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// Qa02GenesisBlockIdentifier is the *types.BlockIdentifier
	// of the Qa02 genesis block.
	Qa02GenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  Qa02GenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// PrinetGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the Prinet genesis block.
	PrinetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  PrinetGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// Currency is the *types.Currency for all
	// Findora networks.
	Currency = &types.Currency{
		Symbol:   Symbol,
		Decimals: Decimals,
	}

	// OperationTypes are all suppoorted operation types.
	OperationTypes = []string{
		MinerRewardOpType,
		UncleRewardOpType,
		FeeOpType,
		CallOpType,
		CreateOpType,
		Create2OpType,
		SelfDestructOpType,
		CallCodeOpType,
		DelegateCallOpType,
		StaticCallOpType,
		DestructOpType,
	}

	// OperationStatuses are all supported operation statuses.
	OperationStatuses = []*types.OperationStatus{
		{
			Status:     SuccessStatus,
			Successful: true,
		},
		{
			Status:     FailureStatus,
			Successful: false,
		},
	}

	// CallMethods are all supported call methods.
	CallMethods = []string{
		"eth_getBlockByNumber",
		"eth_getTransactionReceipt",
		"eth_call",
		"eth_estimateGas",
	}
)

var (
	MainnetChainConfig = &params.ChainConfig{
		ChainID:                 big.NewInt(2152),
		HomesteadBlock:          big.NewInt(0),
		DAOForkBlock:            nil,
		DAOForkSupport:          true,
		EIP150Block:             big.NewInt(0),
		EIP150Hash:              common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
		EIP155Block:             big.NewInt(10),
		EIP158Block:             big.NewInt(10),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     big.NewInt(4_230_000),
		PetersburgBlock:         big.NewInt(4_939_394),
		IstanbulBlock:           big.NewInt(6_485_846),
		MuirGlacierBlock:        big.NewInt(7_117_117),
		BerlinBlock:             big.NewInt(9_812_189),
		LondonBlock:             big.NewInt(10_499_401),
		TerminalTotalDifficulty: new(big.Int).SetUint64(50000000000000000),
		Ethash:                  new(params.EthashConfig),
	}

	AnvilChainConfig = &params.ChainConfig{
		ChainID:                 big.NewInt(2153),
		HomesteadBlock:          big.NewInt(0),
		DAOForkBlock:            nil,
		DAOForkSupport:          true,
		EIP150Block:             big.NewInt(0),
		EIP150Hash:              common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
		EIP155Block:             big.NewInt(10),
		EIP158Block:             big.NewInt(10),
		ByzantiumBlock:          big.NewInt(1_700_000),
		ConstantinopleBlock:     big.NewInt(4_230_000),
		PetersburgBlock:         big.NewInt(4_939_394),
		IstanbulBlock:           big.NewInt(6_485_846),
		MuirGlacierBlock:        big.NewInt(7_117_117),
		BerlinBlock:             big.NewInt(9_812_189),
		LondonBlock:             big.NewInt(10_499_401),
		TerminalTotalDifficulty: new(big.Int).SetUint64(50000000000000000),
		Ethash:                  new(params.EthashConfig),
	}

	Qa02ChainConfig = &params.ChainConfig{
		ChainID:                 big.NewInt(1111),
		HomesteadBlock:          big.NewInt(0),
		DAOForkBlock:            nil,
		DAOForkSupport:          true,
		EIP150Block:             big.NewInt(0),
		EIP150Hash:              common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
		EIP155Block:             big.NewInt(10),
		EIP158Block:             big.NewInt(10),
		ByzantiumBlock:          big.NewInt(1_700_000),
		ConstantinopleBlock:     big.NewInt(4_230_000),
		PetersburgBlock:         big.NewInt(4_939_394),
		IstanbulBlock:           big.NewInt(6_485_846),
		MuirGlacierBlock:        big.NewInt(7_117_117),
		BerlinBlock:             big.NewInt(9_812_189),
		LondonBlock:             big.NewInt(10_499_401),
		TerminalTotalDifficulty: new(big.Int).SetUint64(50000000000000000),
		Ethash:                  new(params.EthashConfig),
	}

	PrinetPChainConfig = &params.ChainConfig{
		ChainID:                 big.NewInt(2152),
		HomesteadBlock:          big.NewInt(0),
		DAOForkBlock:            nil,
		DAOForkSupport:          true,
		EIP150Block:             big.NewInt(0),
		EIP150Hash:              common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
		EIP155Block:             big.NewInt(10),
		EIP158Block:             big.NewInt(10),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     big.NewInt(0),
		PetersburgBlock:         big.NewInt(0),
		IstanbulBlock:           big.NewInt(0),
		MuirGlacierBlock:        big.NewInt(0),
		BerlinBlock:             big.NewInt(0),
		LondonBlock:             big.NewInt(0),
		TerminalTotalDifficulty: new(big.Int).SetUint64(50000000000000000),
		Ethash:                  new(params.EthashConfig),
	}
)

// JSONRPC is the interface for accessing go-ethereum's JSON RPC endpoint.
type JSONRPC interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	Close()
}

// GraphQL is the interface for accessing go-ethereum's GraphQL endpoint.
type GraphQL interface {
	Query(ctx context.Context, input string) (string, error)
}

// CallType returns a boolean indicating
// if the provided trace type is a call type.
func CallType(t string) bool {
	callTypes := []string{
		CallOpType,
		CallCodeOpType,
		DelegateCallOpType,
		StaticCallOpType,
	}

	for _, callType := range callTypes {
		if callType == t {
			return true
		}
	}

	return false
}

// CreateType returns a boolean indicating
// if the provided trace type is a create type.
func CreateType(t string) bool {
	createTypes := []string{
		CreateOpType,
		Create2OpType,
	}

	for _, createType := range createTypes {
		if createType == t {
			return true
		}
	}

	return false
}
