# Findora Rosetta Implement
Rosetta server implementation for Findora BlockChain.

## Findora Rosetta image build and set up example
### Findora Rosetta image build
```bash
go get
go build
```

### Findora Rosetta set up
```bash
export PORT=8080
export MODE=ONLINE
export NETWORK=PRINET
export RPCURL=http://127.0.0.1:8545 
./findora-rosetta run
```


## RPC Endpoints
List of all Findora Rosetta RPC server endpoints
| Method | Path                     | Status | Description
|--------|--------------------------|--------|----------------------------------
| POST   | /network/list            | Y      | Get List of Available Networks
| POST   | /network/status          | Y      | Get Network Status
| POST   | /network/options         | Y      | Get Network Options
| POST   | /block                   | Y      | Get a Block
| POST   | /block/transaction       | Y      | Get a Block Transaction
| POST   | /account/balance         | Y      | Get an Account Balance
| POST   | /construction/submit     | Y      | Submit a Signed Transaction
| POST   | /construction/metadata   | Y      | Get Transaction Construction Metadata
| POST   | /construction/combine    | Y      | Create Network Transaction from Signatures
| POST   | /construction/derive     | Y      | Derive an AccountIdentifier from a PublicKey
| POST   | /construction/hash       | Y      | Get the Hash of a Signed Transaction
| POST   | /construction/parse      | Y      | Parse a Transaction
| POST   | /construction/payloads   | Y      | Generate an Unsigned Transaction and Signing Payloads
| POST   | /construction/preprocess | Y      | Create a Request to Fetch Metadata
| POST   | /call                    | Y      | Perform a Blockchain Call


## RPC Endpoints data api test examples
### RPC Endpoint /network/list test example
```
curl -H 'Content-Type: application/json' --data '{ }' http://127.0.0.1:8080/network/list
```
### RPC Endpoint /network/options test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" } }' http://127.0.0.1:8080/network/options
```
### RPC Endpoint /network/status test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" } }' http://127.0.0.1:8080/network/status
```
### RPC Endpoint /account/balance test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "account_identifier": { "address": "'$addr'" }, "block_identifier": { "index": '$index' } }' http://127.0.0.1/account/balance

curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "account_identifier": { "address": "'$addr'" }, "block_identifier": { "hash": "'$hash'" } }' http://127.0.0.1/account/balance
```
### RPC Endpoint /block test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "index": '$index' } }' http://127.0.0.1/block

curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "hash": "'$hash'" } }' http://127.0.0.1/block
```
### RPC Endpoint /block/transaction test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "index": '$index', "hash": "'$hash'" }, "transaction_identifier": { "hash": "'$tx_hash'" } }' http://127.0.0.1/block/transaction
```

## RPC Endpoints construction api test examples
### RPC Endpoint /construction/derive test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "public_key": { "hex_bytes": "'035d99fa7bf1f9c23f054e7c5af4fe5e0e93cea0c42da7f9883c5da8e187e93644'", "curve_type": "secp256k1" } }' http://127.0.0.1:8080/construction/derive
```

### RPC Endpoint /construction/preprocess test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "operations":[{"operation_identifier": {"index": 0}, "type": "CALL", "account": {"address": "0xA405BA2b64DC04466E0f23487FD1c4A084787326"}, "amount": {"value": "-468962696390199077760007", "currency": {"symbol": "FRA", "decimals": 18}}},{"operation_identifier": {"index": 1}, "type": "CALL", "account": {"address": "0xE7b450C0aae53610e7f315fACf267f9062A99326"}, "amount": {"value": "468962696390199077760007", "currency": {"symbol": "FRA", "decimals": 18}}}]}' http://127.0.0.1:8080/construction/preprocess
```

### RPC Endpoint /construction/metadata test examplec
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "options": {"from": "0xA405BA2b64DC04466E0f23487FD1c4A084787326"}, "public_keys": []}' http://127.0.0.1:8080/construction/metadata
```

### RPC Endpoint /construction/payloads test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "operations": [{"operation_identifier": {"index": 0}, "type": "CALL", "account": {"address": "0xA405BA2b64DC04466E0f23487FD1c4A084787326"}, "amount": {"value": "-468962696390199077760007", "currency": {"symbol": "FRA", "decimals": 18}}},{"operation_identifier": {"index": 1}, "type": "CALL", "account": {"address": "0xE7b450C0aae53610e7f315fACf267f9062A99326"}, "amount": {"value": "468962696390199077760007", "currency": {"symbol": "FRA", "decimals": 18}}}], "metadata": {"gas_price": "0x2540BE400", "nonce": "0x7"}}' http://127.0.0.1:8080/construction/payloads
```

### RPC Endpoint /construction/parse test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "signed":false, "transaction": "{\"from\":\"0xA405BA2b64DC04466E0f23487FD1c4A084787326\",\"to\":\"0xE7b450C0aae53610e7f315fACf267f9062A99326\",\"value\":\"0x634e84ca50084b5fec07\",\"data\":\"0x\",\"nonce\":\"0x7\",\"gas_price\":\"0x2540be400\",\"gas\":\"0x5208\",\"chain_id\":\"0x868\"}"}' http://127.0.0.1:8080/construction/parse
```

### RPC Endpoint /construction/combine test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "unsigned_transaction": "{\"from\": \"0xA405BA2b64DC04466E0f23487FD1c4A084787326\",\"to\": \"0xE7b450C0aae53610e7f315fACf267f9062A99326\",\"value\": \"0x634e84ca50084b5fec07\",\"data\": \"0x\",\"nonce\": \"0x7\",\"gas_price\": \"0x2540be400\",\"gas\": \"0x5208\",\"chain_id\": \"0x868\"}", "signatures":[{"hex_bytes": "8cd208a3770138ce53aad19315dc126a75af39aafe3c41635ad57cd299a82d490ccf6b147fe7b2a68d484b3bac621853272338fea984ba46a088bed5048ed6db00", "signing_payload": {"address": "0xA405BA2b64DC04466E0f23487FD1c4A084787326", "hex_bytes": "a56603ab8e9759e9386381dc300239ba8e841e2a22eab7456ed3c1b3928956e6", "account_identifier": {"address": "0xA405BA2b64DC04466E0f23487FD1c4A084787326"}, "signature_type": "ecdsa_recovery"}, "public_key": {"hex_bytes": "028361d80e1e18ac7406b5233219e64f2ee3033692ab3fbf7e6c457a7d7534750a", "curve_type": "secp256k1"}, "signature_type": "ecdsa_recovery"}]}' http://127.0.0.1:8080/construction/combine
```

### RPC Endpoint /construction/parse test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "signed": true, "transaction": "{\"type\":\"0x0\",\"nonce\":\"0x7\",\"gasPrice\":\"0x2540be400\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\":\"0x5208\",\"value\":\"0x634e84ca50084b5fec07\",\"input\":\"0x\",\"v\":\"0x10f3\",\"r\":\"0x8cd208a3770138ce53aad19315dc126a75af39aafe3c41635ad57cd299a82d49\",\"s\":\"0xccf6b147fe7b2a68d484b3bac621853272338fea984ba46a088bed5048ed6db\",\"to\":\"0xe7b450c0aae53610e7f315facf267f9062a99326\",\"hash\":\"0xac2496ba651c81c757273cf671a599bfebce7715807cfc79ac7f276926527203\"}" }' http://127.0.0.1:8080/construction/parse
```

### RPC Endpoint /construction/hash test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "signed_transaction": "{\"type\": \"0x0\",\"nonce\": \"0x7\",\"gasPrice\": \"0x2540be400\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\": \"0x5208\",\"value\": \"0x634e84ca50084b5fec07\",\"input\": \"0x\",\"v\": \"0x10f3\",\"r\": \"0x8cd208a3770138ce53aad19315dc126a75af39aafe3c41635ad57cd299a82d49\",\"s\": \"0xccf6b147fe7b2a68d484b3bac621853272338fea984ba46a088bed5048ed6db\",\"to\": \"0xe7b450c0aae53610e7f315facf267f9062a99326\",\"hash\": \"0xac2496ba651c81c757273cf671a599bfebce7715807cfc79ac7f276926527203\"}"}' http://127.0.0.1:8080/construction/hash
```

### RPC Endpoint /construction/submit test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier":{"blockchain":"Findora","network":"Prinet"},"signed_transaction":"{\"type\":\"0x0\",\"nonce\":\"0x7\",\"gasPrice\":\"0x2540be400\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\":\"0x5208\",\"value\":\"0x6251e197ca22ee20c76\",\"input\":\"0x\",\"v\":\"0x10f3\",\"r\":\"0xab055b55a1bf8f59991a13cf483a100269fd504ec96913363b173f3e8a15dd43\",\"s\":\"0x7d247f949159cf9703b326424e2e4a6fb9ba1e96da2f7e64cc7ab06a53d3be9b\",\"to\":\"0x8fb9e7034cb6254eb3cb7673ed25181bad64259d\",\"hash\":\"0x328dcf2778e5e6b9e9a6e9ecc71622876716b66ecc19b221fca8af655a0ef064\"}"}' http://127.0.0.1:8080/construction/submit
```


## Rosetta Cli Tools test examples
### Data Api test example
```
rm -rf test-cli
mkdir -p test-cli/rosetta-data
export ROSETTA_CONFIGURATION_FILE=./rosetta-cli-conf/prinet/config.json
rosetta-cli check:data --start-block $start_block
```

### Construction Api test example
```
rm -rf test-cli
mkdir -p test-cli/rosetta-data
export ROSETTA_CONFIGURATION_FILE=./rosetta-cli-conf/prinet/config.json
rosetta-cli check:construction
```

## Findora Rosetta Docker setup
### Findora Rosetta Docker build and run
```
docker build . -t findora-rosetta
docker run -p 8080:8080 -p 8545:8545 -e MODE=OFFLINE -itd --name findora-rosetta --restart always --privileged=true findora-rosetta
```
### Findora Rosetta Docker test
```
git clone https://github.com/FindoraNetwork/rosetta-cli.git
cd rosetta-cli
go build
export ROSETTA_CONFIGURATION_FILE=./rosetta-cli-conf/prinet/config.json
./test-examples/cli_construction_api_tests.sh prinet
./test-examples/cli_data_api_tests.sh prinet 2
```