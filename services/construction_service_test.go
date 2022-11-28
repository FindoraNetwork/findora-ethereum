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

package services

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"testing"

	"github/findoranetwork/findora-rosetta/configuration"
	findora "github/findoranetwork/findora-rosetta/findora"
	mocks "github/findoranetwork/findora-rosetta/mocks/services"

	"github.com/ethereum/go-ethereum/common"
	"github.com/findoranetwork/rosetta-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func forceHexDecode(t *testing.T, s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		t.Fatalf("could not decode hex %s", s)
	}

	return b
}

func forceMarshalMap(t *testing.T, i interface{}) map[string]interface{} {
	m, err := marshalJSONMap(i)
	if err != nil {
		t.Fatalf("could not marshal map %s", types.PrintStruct(i))
	}

	return m
}

func TestConstructionService(t *testing.T) {
	networkIdentifier = &types.NetworkIdentifier{
		Network:    findora.AnvilNetwork,
		Blockchain: findora.Blockchain,
	}

	cfg := &configuration.Configuration{
		Mode:    configuration.Online,
		Network: networkIdentifier,
		Params:  findora.AnvilChainConfig,
	}

	mockClient := &mocks.Client{}
	servicer := NewConstructionAPIService(cfg, mockClient)
	ctx := context.Background()

	// Test Derive
	// curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "public_key": { "hex_bytes": "'$PUBKEY'", "curve_type": "secp256k1" } }' http://127.0.0.1:8080/construction/derive

	publicKey := &types.PublicKey{
		Bytes: forceHexDecode(
			t,
			"03d3d3358e7f69cbe45bde38d7d6f24660c7eeeaee5c5590cfab985c8839b21fd5",
		),
		CurveType: types.Secp256k1,
	}
	deriveResponse, err := servicer.ConstructionDerive(ctx, &types.ConstructionDeriveRequest{
		NetworkIdentifier: networkIdentifier,
		PublicKey:         publicKey,
	})
	assert.Nil(t, err)
	assert.Equal(t, &types.ConstructionDeriveResponse{
		AccountIdentifier: &types.AccountIdentifier{
			Address: "0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309",
		},
	}, deriveResponse)

	// Test Preprocess
	// curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"operations":[{"operation_identifier":{"index":0},"type":"CALL","account":{"address":"0xa46Ff39c35da34E91f308389b64f43226B579327"},"amount":{"value":"-42894881044106498","currency":{"symbol":"FRA","decimals":18}}},{"operation_identifier":{"index":1},"type":"CALL","account":{"address":"0x57B414a0332B5CaB885a451c2a28a07d1e9b8a8d"},"amount":{"value":"42894881044106498","currency":{"symbol":"FRA","decimals":18}}}]}' http://127.0.0.1:8080/construction/preprocess

	intent := `[{"operation_identifier":{"index":0},"type":"CALL","account":{"address":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"},"amount":{"value":"-42894881044106498","currency":{"symbol":"ETH","decimals":18}}},{"operation_identifier":{"index":1},"type":"CALL","account":{"address":"0x57B414a0332B5CaB885a451c2a28a07d1e9b8a8d"},"amount":{"value":"42894881044106498","currency":{"symbol":"ETH","decimals":18}}}]` // nolint
	var ops []*types.Operation
	assert.NoError(t, json.Unmarshal([]byte(intent), &ops))
	preprocessResponse, err := servicer.ConstructionPreprocess(
		ctx,
		&types.ConstructionPreprocessRequest{
			NetworkIdentifier: networkIdentifier,
			Operations:        ops,
		},
	)
	assert.Nil(t, err)
	optionsRaw := `{"from":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"}`
	var options options
	assert.NoError(t, json.Unmarshal([]byte(optionsRaw), &options))
	assert.Equal(t, &types.ConstructionPreprocessResponse{
		Options: forceMarshalMap(t, options),
	}, preprocessResponse)

	// Test Metadata
	// curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"options":{"from":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"}}' http://127.0.0.1:8080/construction/metadata

	metadata := &metadata{
		GasPrice: big.NewInt(1000000000),
		Nonce:    0,
	}

	mockClient.On(
		"SuggestGasPrice",
		ctx,
	).Return(
		big.NewInt(1000000000),
		nil,
	).Once()
	mockClient.On(
		"PendingNonceAt",
		ctx,
		common.HexToAddress("0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"),
	).Return(
		uint64(0),
		nil,
	).Once()
	metadataResponse, err := servicer.ConstructionMetadata(ctx, &types.ConstructionMetadataRequest{
		NetworkIdentifier: networkIdentifier,
		Options:           forceMarshalMap(t, options),
	})
	assert.Nil(t, err)
	assert.Equal(t, &types.ConstructionMetadataResponse{
		Metadata: forceMarshalMap(t, metadata),
		SuggestedFee: []*types.Amount{
			{
				Value:    "21000000000000",
				Currency: findora.Currency,
			},
		},
	}, metadataResponse)

	// Test Payloads
	// curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"operations":[{"operation_identifier":{"index":0},"type":"CALL","account":{"address":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"},"amount":{"value":"-42894881044106498","currency":{"symbol":"FRA","decimals":18}}},{"operation_identifier":{"index":1},"type":"CALL","account":{"address":"0x57B414a0332B5CaB885a451c2a28a07d1e9b8a8d"},"amount":{"value":"42894881044106498","currency":{"symbol":"FRA","decimals":18}}}],"metadata":{"gas_price":"0x2540BE400","nonce":"0x0"}}' http://127.0.0.1:8080/construction/payloads

	unsignedRaw := `{"from":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309","to":"0x57B414a0332B5CaB885a451c2a28a07d1e9b8a8d","value":"0x9864aac3510d02","data":"0x","nonce":"0x0","gas_price":"0x3b9aca00","gas":"0x5208","chain_id":"0x3"}` // nolint
	payloadsResponse, err := servicer.ConstructionPayloads(ctx, &types.ConstructionPayloadsRequest{
		NetworkIdentifier: networkIdentifier,
		Operations:        ops,
		Metadata:          forceMarshalMap(t, metadata),
	})
	assert.Nil(t, err)
	payloadsRaw := `[{"address":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309","hex_bytes":"b682f3e39c512ff57471f482eab264551487320cbd3b34485f4779a89e5612d1","account_identifier":{"address":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"},"signature_type":"ecdsa_recovery"}]` // nolint
	var payloads []*types.SigningPayload
	assert.NoError(t, json.Unmarshal([]byte(payloadsRaw), &payloads))
	assert.Equal(t, &types.ConstructionPayloadsResponse{
		UnsignedTransaction: unsignedRaw,
		Payloads:            payloads,
	}, payloadsResponse)

	// Test Parse Unsigned
	// curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"signed":true,"transaction":"{\"type\":\"0x0\",\"nonce\":\"0x2\",\"gasPrice\":\"0x2540be400\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\":\"0x5208\",\"value\":\"0x9864aac3510d02\",\"input\":\"0x\",\"v\":\"0x10f4\",\"r\":\"0x8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b2\",\"s\":\"0x5a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e8\",\"to\":\"0x57b414a0332b5cab885a451c2a28a07d1e9b8a8d\",\"hash\":\"0x3125474d60f757a35a9b98ca0e3817d62d0d8720a45f3a74ef84e6aa892bcc2a\"}"}' http://127.0.0.1:8080/construction/parse

	parseOpsRaw := `[{"operation_identifier":{"index":0},"type":"CALL","account":{"address":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"},"amount":{"value":"-42894881044106498","currency":{"symbol":"ETH","decimals":18}}},{"operation_identifier":{"index":1},"related_operations":[{"index":0}],"type":"CALL","account":{"address":"0x57B414a0332B5CaB885a451c2a28a07d1e9b8a8d"},"amount":{"value":"42894881044106498","currency":{"symbol":"ETH","decimals":18}}}]` // nolint
	var parseOps []*types.Operation
	assert.NoError(t, json.Unmarshal([]byte(parseOpsRaw), &parseOps))
	parseUnsignedResponse, err := servicer.ConstructionParse(ctx, &types.ConstructionParseRequest{
		NetworkIdentifier: networkIdentifier,
		Signed:            false,
		Transaction:       unsignedRaw,
	})
	assert.Nil(t, err)
	parseMetadata := &parseMetadata{
		Nonce:    metadata.Nonce,
		GasPrice: metadata.GasPrice,
		ChainID:  big.NewInt(3),
	}
	assert.Equal(t, &types.ConstructionParseResponse{
		Operations:               parseOps,
		AccountIdentifierSigners: []*types.AccountIdentifier{},
		Metadata:                 forceMarshalMap(t, parseMetadata),
	}, parseUnsignedResponse)

	// Test Combine
	// curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"unsigned_transaction":"{\"from\":\"0xa46Ff39c35da34E91f308389b64f43226B579327\",\"to\":\"0x57B414a0332B5CaB885a451c2a28a07d1e9b8a8d\",\"value\":\"0x9864aac3510d02\",\"data\":\"0x\",\"nonce\":\"0x0\",\"gas_price\":\"0x2540BE400\",\"gas\":\"0x5208\",\"chain_id\":\"0x868\"}","signatures":[{"hex_bytes":"8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b25a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e801","signing_payload":{"address":"0xA405BA2b64DC04466E0f23487FD1c4A084787326","hex_bytes":"b682f3e39c512ff57471f482eab264551487320cbd3b34485f4779a89e5612d1","account_identifier":{"address":"0xA405BA2b64DC04466E0f23487FD1c4A084787326"},"signature_type":"ecdsa_recovery"},"public_key":{"hex_bytes":"03d3d3358e7f69cbe45bde38d7d6f24660c7eeeaee5c5590cfab985c8839b21fd5","curve_type":"secp256k1"},"signature_type":"ecdsa_recovery"}]}' http://127.0.0.1:8080/construction/combine

	signaturesRaw := `[{"hex_bytes":"8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b25a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e801","signing_payload":{"address":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309","hex_bytes":"b682f3e39c512ff57471f482eab264551487320cbd3b34485f4779a89e5612d1","account_identifier":{"address":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"},"signature_type":"ecdsa_recovery"},"public_key":{"hex_bytes":"03d3d3358e7f69cbe45bde38d7d6f24660c7eeeaee5c5590cfab985c8839b21fd5","curve_type":"secp256k1"},"signature_type":"ecdsa_recovery"}]` // nolint
	var signatures []*types.Signature
	assert.NoError(t, json.Unmarshal([]byte(signaturesRaw), &signatures))
	signedRaw := `{"type":"0x0","nonce":"0x0","gasPrice":"0x3b9aca00","maxPriorityFeePerGas":null,"maxFeePerGas":null,"gas":"0x5208","value":"0x9864aac3510d02","input":"0x","v":"0x2a","r":"0x8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b2","s":"0x5a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e8","to":"0x57b414a0332b5cab885a451c2a28a07d1e9b8a8d","hash":"0x424969b1a98757bcd748c60bad2a7de9745cfb26bfefb4550e780a098feada42"}` // nolint
	combineResponse, err := servicer.ConstructionCombine(ctx, &types.ConstructionCombineRequest{
		NetworkIdentifier:   networkIdentifier,
		UnsignedTransaction: unsignedRaw,
		Signatures:          signatures,
	})
	assert.Nil(t, err)
	assert.Equal(t, &types.ConstructionCombineResponse{
		SignedTransaction: signedRaw,
	}, combineResponse)

	// Test Parse Signed
	parseSignedResponse, err := servicer.ConstructionParse(ctx, &types.ConstructionParseRequest{
		NetworkIdentifier: networkIdentifier,
		Signed:            true,
		Transaction:       signedRaw,
	})
	assert.Nil(t, err)
	assert.Equal(t, &types.ConstructionParseResponse{
		Operations: parseOps,
		AccountIdentifierSigners: []*types.AccountIdentifier{
			{Address: "0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"},
		},
		Metadata: forceMarshalMap(t, parseMetadata),
	}, parseSignedResponse)

	// Test Hash
	// curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"signed_transaction":"{\"type\":\"0x0\",\"nonce\":\"0x2\",\"gasPrice\":\"0x2540BE400\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\":\"0x5208\",\"value\":\"0x9864aac3510d02\",\"input\":\"0x\",\"v\":\"0x2a\",\"r\":\"0x8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b2\",\"s\":\"0x5a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e8\",\"to\":\"0x57b414a0332b5cab885a451c2a28a07d1e9b8a8d\",\"hash\":\"0x424969b1a98757bcd748c60bad2a7de9745cfb26bfefb4550e780a098feada42\"}"}' http::/127.0.0.1:8080/construction/hash

	transactionIdentifier := &types.TransactionIdentifier{
		Hash: "0x424969b1a98757bcd748c60bad2a7de9745cfb26bfefb4550e780a098feada42",
	}
	hashResponse, err := servicer.ConstructionHash(ctx, &types.ConstructionHashRequest{
		NetworkIdentifier: networkIdentifier,
		SignedTransaction: signedRaw,
	})
	assert.Nil(t, err)
	assert.Equal(t, &types.TransactionIdentifierResponse{
		TransactionIdentifier: transactionIdentifier,
	}, hashResponse)

	// Test Submit
	// curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"signed_transaction":"{\"type\":\"0x0\",\"nonce\":\"0x2\",\"gasPrice\":\"0x2540be400\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\":\"0x5208\",\"value\":\"0x9864aac3510d02\",\"input\":\"0x\",\"v\":\"0x10f4\",\"r\":\"0x8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b2\",\"s\":\"0x5a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e8\",\"to\":\"0x57b414a0332b5cab885a451c2a28a07d1e9b8a8d\",\"hash\":\"0xfe312a7788fd578725c1c1917b06d2df46fd50132189fd37aff261075e3d0502\"}"}' http::/127.0.0.1:8080//construction/

	mockClient.On(
		"SendTransaction",
		ctx,
		mock.Anything, // can't test ethTx here because it contains "time"
	).Return(
		nil,
	)
	submitResponse, err := servicer.ConstructionSubmit(ctx, &types.ConstructionSubmitRequest{
		NetworkIdentifier: networkIdentifier,
		SignedTransaction: signedRaw,
	})
	assert.Nil(t, err)
	assert.Equal(t, &types.TransactionIdentifierResponse{
		TransactionIdentifier: transactionIdentifier,
	}, submitResponse)

	mockClient.AssertExpectations(t)
}
