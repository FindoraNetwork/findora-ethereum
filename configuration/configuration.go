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

package configuration

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	findora "github/findoranetwork/findora-rosetta/findora"

	"github.com/ethereum/go-ethereum/params"
	"github.com/findoranetwork/rosetta-sdk-go/types"
)

// Mode is the setting that determines if
// the implementation is "online" or "offline".
type Mode string

const (
	// Online is when the implementation is permitted
	// to make outbound connections.
	Online Mode = "ONLINE"

	// Offline is when the implementation is not permitted
	// to make outbound connections.
	Offline Mode = "OFFLINE"

	// Mainnet is the findora Mainnet.
	Mainnet string = "MAINNET"

	// Testnet defaults to `Anvil` for backwards compatibility.
	Testnet string = "TESTNET"

	// Anvil is the findora Anvil testnet.
	Anvil string = "ANVIL"

	// Qa02 is the findora Qa02 testnet.
	Qa02 string = "QA02"

	// Prinet is the findora Prinet testnet.
	Prinet string = "PRINET"

	// DataDirectory is the default location for all
	// persistent data.
	DataDirectory = "/data"

	// ModeEnv is the environment variable read
	// to determine mode.
	ModeEnv = "MODE"

	// NetworkEnv is the environment variable
	// read to determine network.
	NetworkEnv = "NETWORK"

	// PortEnv is the environment variable
	// read to determine the port for the Rosetta
	// implementation.
	PortEnv = "PORT"

	// RpcEnv is an optional environment variable
	// used to connect findora-rosetta to an already
	// running findora node.
	RpcEnv = "RPCURL"

	// DefaultRpcURL is the default URL for
	// a running findora node. This is used
	// when RpcEnv is not populated.
	DefaultRpcURL = "http://127.0.0.1:8545"

	// SkipFindoraAdminEnv is an optional environment variable
	// to skip findora `admin` calls which are typically not supported
	// by hosted node services. When not set, defaults to false.
	SkipFindoraAdminEnv = "SKIP_FINDORA_ADMIN"

	// MiddlewareVersion is the version of findora-rosetta.
	MiddlewareVersion = "0.0.4"
)

// Configuration determines how
type Configuration struct {
	Mode                   Mode
	Network                *types.NetworkIdentifier
	GenesisBlockIdentifier *types.BlockIdentifier
	RpcURL                 string
	RemoteRpc              bool
	Port                   int
	FindoraArguments       string
	SkipFindoraAdmin       bool

	// Block Reward Data
	Params *params.ChainConfig
}

// LoadConfiguration attempts to create a new Configuration
// using the ENVs in the environment.
func LoadConfiguration() (*Configuration, error) {
	config := &Configuration{}

	modeValue := Mode(os.Getenv(ModeEnv))
	switch modeValue {
	case Online:
		config.Mode = Online
	case Offline:
		config.Mode = Offline
	case "":
		return nil, errors.New("MODE must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid mode", modeValue)
	}

	networkValue := os.Getenv(NetworkEnv)
	switch networkValue {
	case Mainnet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: findora.Blockchain,
			Network:    findora.MainnetNetwork,
		}
		config.GenesisBlockIdentifier = findora.MainnetGenesisBlockIdentifier
		config.Params = findora.MainnetChainConfig
		config.FindoraArguments = findora.MainnetCommandArguments
	case Testnet, Anvil:
		config.Network = &types.NetworkIdentifier{
			Blockchain: findora.Blockchain,
			Network:    findora.AnvilNetwork,
		}
		config.GenesisBlockIdentifier = findora.AnvilGenesisBlockIdentifier
		config.Params = findora.AnvilChainConfig
		config.FindoraArguments = findora.AnvilCommandArguments
	case Qa02:
		config.Network = &types.NetworkIdentifier{
			Blockchain: findora.Blockchain,
			Network:    findora.Qa02Network,
		}
		config.GenesisBlockIdentifier = findora.Qa02GenesisBlockIdentifier
		config.Params = findora.Qa02ChainConfig
		config.FindoraArguments = findora.Qa02CommandArguments
	case Prinet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: findora.Blockchain,
			Network:    findora.PrinetNetwork,
		}
		config.GenesisBlockIdentifier = findora.PrinetGenesisBlockIdentifier
		config.Params = findora.PrinetPChainConfig
		config.FindoraArguments = findora.PrinetCommandArguments
	case "":
		return nil, errors.New("NETWORK must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid network", networkValue)
	}

	config.RpcURL = DefaultRpcURL
	envRpcURL := os.Getenv(RpcEnv)
	if len(envRpcURL) > 0 {
		config.RemoteRpc = true
		config.RpcURL = envRpcURL
	}

	config.SkipFindoraAdmin = false
	envSkipFindoraAdmin := os.Getenv(SkipFindoraAdminEnv)
	if len(envSkipFindoraAdmin) > 0 {
		val, err := strconv.ParseBool(envSkipFindoraAdmin)
		if err != nil {
			return nil, fmt.Errorf("%w: unable to parse SKIP_FINDORA_ADMIN %s", err, envSkipFindoraAdmin)
		}
		config.SkipFindoraAdmin = val
	}

	portValue := os.Getenv(PortEnv)
	if len(portValue) == 0 {
		return nil, errors.New("PORT must be populated")
	}

	port, err := strconv.Atoi(portValue)
	if err != nil || len(portValue) == 0 || port <= 0 {
		return nil, fmt.Errorf("%w: unable to parse port %s", err, portValue)
	}
	config.Port = port

	return config, nil
}
