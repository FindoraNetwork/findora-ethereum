#!/bin/bash
export RPC_URL=http://127.0.0.1:8080

echo "RPC Endpoint /network/list test example"
curl -H 'Content-Type: application/json' --data '{ }' $RPC_URL/network/list

echo "RPC Endpoint /network/options test example"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" } }' $RPC_URL/network/options

echo "RPC Endpoint /network/status test example"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" } }' $RPC_URL/network/status

echo "RPC Endpoint /account/balance test example by index"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "account_identifier": { "address": "'$addr'" }, "block_identifier": { "index": '$index' } }' $RPC_URL/account/balance

echo "RPC Endpoint /account/balance test example by hash"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "account_identifier": { "address": "'$addr'" }, "block_identifier": { "hash": "'$hash'" } }' $RPC_URL/account/balance

echo "RPC Endpoint /block test example by index"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "index": '$index' } }' $RPC_URL/block

echo "RPC Endpoint /block test example by hash"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "hash": "'$hash'" } }' $RPC_URL/block

echo "RPC Endpoint /block/transaction test example"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "index": '$index', "hash": "'$hash'" }, "transaction_identifier": { "hash": "'$tx_hash'" } }' $RPC_URL/block/transaction

echo "RPC Endpoint /construction/derive test example"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "public_key": { "hex_bytes": "'$pub_key'", "curve_type": "secp256k1" } }' $RPC_URL/construction/derive

echo "RPC Endpoint /construction/payloads test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "operations": [{"operation_identifier": {"index": '$index'}, "type": "CALL", "account": {"address": "'$addr'"}, "amount": {"value": "'$reduce_value'", "currency": {"symbol": "FRA", "decimals": 18}}},{"operation_identifier": {"index": '$index'}, "type": "CALL", "account": {"address": "'$addr'"}, "amount": {"value": "'$add_value'", "currency": {"symbol": "FRA", "decimals": 18}}}], "metadata": {"gas_price": "'$gas_price'", "nonce": "'$nonce'"}}' $RPC_URL/construction/payloads

echo "RPC Endpoint /construction/preprocess test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "operations":[{"operation_identifier": {"index": '$index'}, "type": "CALL", "account": {"address": "'$addr'"}, "amount": {"value": "'$reduce_value'", "currency": {"symbol": "FRA", "decimals": 18}}},{"operation_identifier": {"index": '$index'}, "type": "CALL", "account": {"address": "'$addr'"}, "amount": {"value": "'$add_value'", "currency": {"symbol": "FRA", "decimals": 18}}}]}' $RPC_URL/construction/preprocess

echo "RPC Endpoint /construction/hash test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "signed_transaction": "{\"type\": \"0x0\",\"nonce\": \"'$nonce'\",\"gasPrice\": \"'$gas_price'\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\": \"'$gas'\",\"value\": \"'$value'\",\"input\": \"0x\",\"v\": \"'$v'\",\"r\": \"'$r'\",\"s\": \"'$s'\",\"to\": \"'$to'\",\"hash\": \"'$hash'\"}"}' $RPC_URL/construction/hash

echo "RPC Endpoint /construction/parse test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "signed":true, "transaction": "{\"type\": \"0x0\",\"nonce\": \"'$nonce'\",\"gasPrice\": \"'$gas_price'\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\": \"'$gas'\",\"value\": \"'$value'\",\"input\": \"0x\",\"v\": \"'$v'\",\"r\": \"'$r'\",\"s\": \"'$s'\",\"to\": \"'$to'\",\"hash\": \"'$hash'\"}"}' $RPC_URL/construction/parse

echo "RPC Endpoint /construction/combine test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "unsigned_transaction": "{\"from\": \"'$addr'\",\"to\": \"'$addr'\",\"value\": \"'$value'\",\"data\": \"0x\",\"nonce\": \"'$nonce'\",\"gas_price\": \"'$gas_price'\",\"gas\": \"'$gas'\",\"chain_id\": \"'$chain_id'\"}", "signatures":[{"hex_bytes": "'$signatures'", "signing_payload": {"address": "'$addr'", "hex_bytes": "'$signing_payload'", "account_identifier": {"address": "'$addr'"}, "signature_type": "ecdsa_recovery"}, "public_key": {"hex_bytes": "'$pub_key'", "curve_type": "secp256k1"}, "signature_type": "ecdsa_recovery"}]}' $RPC_URL/construction/combine

echo "RPC Endpoint /construction/metadata test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "options": {"from": "'$addr'"}}' $RPC_URL/construction/metadata

echo "RPC Endpoint /construction/submit test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "signed_transaction": "{\"type\": \"0x0\",\"nonce\": \"'$nonce'\",\"gasPrice\": \"'$gas_price'\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\": \"'$gas'\",\"value\": \"'$value'\",\"input\": \"0x\",\"v\": \"'$v'\",\"r\": \"'$r'\",\"s\": \"'$s'\",\"to\": \"'$addr'\",\"hash\": \"'$hash'\"}"}' $RPC_URL/construction/submit
