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


## RPC Endpoints test examples
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
### RPC Endpoint /construction/submit test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "signed_transaction": "{\"type\": \"0x0\",\"nonce\": \"'$nonce'\",\"gasPrice\": \"'$gas_price'\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\": \"'$gas'\",\"value\": \"'$value'\",\"input\": \"0x\",\"v\": \"'$v'\",\"r\": \"'$r'\",\"s\": \"'$s'\",\"to\": \"'$addr'\",\"hash\": \"'$hash'\"}"}' http://127.0.0.1/construction/submit
```
### RPC Endpoint /construction/metadata test examplec
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "options": {"from": "'$addr'"}}' http://127.0.0.1/construction/metadata
```
### RPC Endpoint /construction/combine test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "unsigned_transaction": "{\"from\": \"'$addr'\",\"to\": \"'$addr'\",\"value\": \"'$value'\",\"data\": \"0x\",\"nonce\": \"'$nonce'\",\"gas_price\": \"'$gas_price'\",\"gas\": \"'$gas'\",\"chain_id\": \"'$chain_id'\"}", "signatures":[{"hex_bytes": "'$signatures'", "signing_payload": {"address": "'$addr'", "hex_bytes": "'$signing_payload'", "account_identifier": {"address": "'$addr'"}, "signature_type": "ecdsa_recovery"}, "public_key": {"hex_bytes": "'$pub_key'", "curve_type": "secp256k1"}, "signature_type": "ecdsa_recovery"}]}' http://127.0.0.1/construction/combine
```
### RPC Endpoint /construction/derive test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "public_key": { "hex_bytes": "'$pub_key'", "curve_type": "secp256k1" } }' http://127.0.0.1/construction/derive
```
### RPC Endpoint /construction/hash test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "signed_transaction": "{\"type\": \"0x0\",\"nonce\": \"'$nonce'\",\"gasPrice\": \"'$gas_price'\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\": \"'$gas'\",\"value\": \"'$value'\",\"input\": \"0x\",\"v\": \"'$v'\",\"r\": \"'$r'\",\"s\": \"'$s'\",\"to\": \"'$to'\",\"hash\": \"'$hash'\"}"}' http://127.0.0.1/construction/hash
```
### RPC Endpoint /construction/parse test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "signed":true, "transaction": "{\"type\": \"0x0\",\"nonce\": \"'$nonce'\",\"gasPrice\": \"'$gas_price'\",\"maxPriorityFeePerGas\":null,\"maxFeePerGas\":null,\"gas\": \"'$gas'\",\"value\": \"'$value'\",\"input\": \"0x\",\"v\": \"'$v'\",\"r\": \"'$r'\",\"s\": \"'$s'\",\"to\": \"'$to'\",\"hash\": \"'$hash'\"}"}' http://127.0.0.1/construction/parse
```
### RPC Endpoint /construction/payloads test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "operations": [{"operation_identifier": {"index": '$index'}, "type": "CALL", "account": {"address": "'$addr'"}, "amount": {"value": "'$reduce_value'", "currency": {"symbol": "FRA", "decimals": 18}}},{"operation_identifier": {"index": '$index'}, "type": "CALL", "account": {"address": "'$addr'"}, "amount": {"value": "'$add_value'", "currency": {"symbol": "FRA", "decimals": 18}}}], "metadata": {"gas_price": "'$gas_price'", "nonce": "'$nonce'"}}' http://127.0.0.1/construction/payloads
```
### RPC Endpoint /construction/preprocess test example
```
curl -H 'Content-Type: application/json' --data '{"network_identifier": {"blockchain": "Findora", "network": "Prinet"}, "operations":[{"operation_identifier": {"index": '$index'}, "type": "CALL", "account": {"address": "'$addr'"}, "amount": {"value": "'$reduce_value'", "currency": {"symbol": "FRA", "decimals": 18}}},{"operation_identifier": {"index": '$index'}, "type": "CALL", "account": {"address": "'$addr'"}, "amount": {"value": "'$add_value'", "currency": {"symbol": "FRA", "decimals": 18}}}]}' http://127.0.0.1/construction/preprocess
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