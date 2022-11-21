#!/bin/bash
export RPC_URL=http://127.0.0.1:8080
export ADDR=0x872071Fe160d99387041f8cA8a41Cc7307955A2d
export INDEX=68
export HASH=
export TXHASH=0xefac09b729a3d882ef78caac4d73554eb491a59c8478399b9fb4e0e8833ca145
export PUBKEY=028361d80e1e18ac7406b5233219e64f2ee3033692ab3fbf7e6c457a7d7534750a
export NONCE=0x2

echo "RPC Endpoint /network/list test example"
curl -H 'Content-Type: application/json' --data '{ }' $RPC_URL/network/list

echo "RPC Endpoint /network/options test example"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" } }' $RPC_URL/network/options

echo "RPC Endpoint /network/status test example"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" } }' $RPC_URL/network/status

echo "RPC Endpoint /account/balance test example by index"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "account_identifier": { "address": "'$ADDR'" }, "block_identifier": { "index": '$INDEX' } }' $RPC_URL/account/balance

echo "RPC Endpoint /account/balance test example by hash"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "account_identifier": { "address": "address" }, "block_identifier": { "hash": "hash" } }' $RPC_URL/account/balance

echo "RPC Endpoint /block test example by index"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "index": '$INDEX' } }' $RPC_URL/block

echo "RPC Endpoint /block test example by hash"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "hash": "'$HASH'" } }' $RPC_URL/block

echo "RPC Endpoint /block/transaction test example"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "index": '$INDEX', "hash": "'$HASH'" }, "transaction_identifier": { "hash": "'$TXHASH'" } }' $RPC_URL/block/transaction

echo "RPC Endpoint /construction/derive test example"
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "public_key": { "hex_bytes": "'$PUBKEY'", "curve_type": "secp256k1" } }' $RPC_URL/construction/derive

echo "RPC Endpoint /construction/payloads test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"operations":[{"operation_identifier":{"index":0},"type":"CALL","account":{"address":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"},"amount":{"value":"-42894881044106498","currency":{"symbol":"FRA","decimals":18}}},{"operation_identifier":{"index":1},"type":"CALL","account":{"address":"0x57B414a0332B5CaB885a451c2a28a07d1e9b8a8d"},"amount":{"value":"42894881044106498","currency":{"symbol":"FRA","decimals":18}}}],"metadata":{"gas_price":"0x2540BE400","nonce":"0x0"}}' $RPC_URL/construction/payloads

echo "RPC Endpoint /construction/preprocess test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"operations":[{"operation_identifier":{"index":0},"type":"CALL","account":{"address":"0xa46Ff39c35da34E91f308389b64f43226B579327"},"amount":{"value":"-42894881044106498","currency":{"symbol":"FRA","decimals":18}}},{"operation_identifier":{"index":1},"type":"CALL","account":{"address":"0x57B414a0332B5CaB885a451c2a28a07d1e9b8a8d"},"amount":{"value":"42894881044106498","currency":{"symbol":"FRA","decimals":18}}}]}' $RPC_URL/construction/preprocess

echo "RPC Endpoint /construction/hash test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"signed_transaction":"{\"type\":\"0x0\",\"nonce\":\"0x2\",\"gasPrice\":\"0x2540BE400\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\":\"0x5208\",\"value\":\"0x9864aac3510d02\",\"input\":\"0x\",\"v\":\"0x2a\",\"r\":\"0x8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b2\",\"s\":\"0x5a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e8\",\"to\":\"0x57b414a0332b5cab885a451c2a28a07d1e9b8a8d\",\"hash\":\"0x424969b1a98757bcd748c60bad2a7de9745cfb26bfefb4550e780a098feada42\"}"}' $RPC_URL/construction/hash

echo "RPC Endpoint /construction/parse test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"signed":true,"transaction":"{\"type\":\"0x0\",\"nonce\":\"0x2\",\"gasPrice\":\"0x2540be400\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\":\"0x5208\",\"value\":\"0x9864aac3510d02\",\"input\":\"0x\",\"v\":\"0x10f4\",\"r\":\"0x8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b2\",\"s\":\"0x5a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e8\",\"to\":\"0x57b414a0332b5cab885a451c2a28a07d1e9b8a8d\",\"hash\":\"0x3125474d60f757a35a9b98ca0e3817d62d0d8720a45f3a74ef84e6aa892bcc2a\"}"}' $RPC_URL/construction/parse

echo "RPC Endpoint /construction/combine test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"unsigned_transaction":"{\"from\":\"0xa46Ff39c35da34E91f308389b64f43226B579327\",\"to\":\"0x57B414a0332B5CaB885a451c2a28a07d1e9b8a8d\",\"value\":\"0x9864aac3510d02\",\"data\":\"0x\",\"nonce\":\"0x0\",\"gas_price\":\"0x2540BE400\",\"gas\":\"0x5208\",\"chain_id\":\"0x868\"}","signatures":[{"hex_bytes":"8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b25a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e801","signing_payload":{"address":"0xA405BA2b64DC04466E0f23487FD1c4A084787326","hex_bytes":"b682f3e39c512ff57471f482eab264551487320cbd3b34485f4779a89e5612d1","account_identifier":{"address":"0xA405BA2b64DC04466E0f23487FD1c4A084787326"},"signature_type":"ecdsa_recovery"},"public_key":{"hex_bytes":"03d3d3358e7f69cbe45bde38d7d6f24660c7eeeaee5c5590cfab985c8839b21fd5","curve_type":"secp256k1"},"signature_type":"ecdsa_recovery"}]}' $RPC_URL/construction/combine

echo "RPC Endpoint /construction/metadata test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"options":{"from":"0xe3a5B4d7f79d64088C8d4ef153A7DDe2B2d47309"}}' $RPC_URL/construction/metadata

echo "RPC Endpoint /construction/submit test example"
curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"signed_transaction":"{\"type\":\"0x0\",\"nonce\":\"0x2\",\"gasPrice\":\"0x2540be400\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\":\"0x5208\",\"value\":\"0x9864aac3510d02\",\"input\":\"0x\",\"v\":\"0x10f4\",\"r\":\"0x8c712c64bc65c4a88707fa93ecd090144dffb1bf133805a10a51d354c2f9f2b2\",\"s\":\"0x5a63cea6989f4c58372c41f31164036a6b25dce1d5c05e1d31c16c0590c176e8\",\"to\":\"0x57b414a0332b5cab885a451c2a28a07d1e9b8a8d\",\"hash\":\"0xfe312a7788fd578725c1c1917b06d2df46fd50132189fd37aff261075e3d0502\"}"}' $RPC_URL/construction/submit
