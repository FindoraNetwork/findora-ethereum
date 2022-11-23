# Findora Rosetta Implement
Rosetta server implementation for Findora BlockChain.

## Findora Rosetta image build and run example
```bash
go get
go build
export PORT=8080
export MODE=ONLINE
export NETWORK=PRINET
export RPCURL=http://127.0.0.1:8545 
export GENEHIGH=1
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
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "account_identifier": { "address": "address" }, "block_identifier": { "index": index } }' http://localhost:8080/account/balance

curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "account_identifier": { "address": "address" }, "block_identifier": { "hash": "hash" } }' http://localhost:8080/account/balance
```
### RPC Endpoint /block test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "index": index } }' http://127.0.0.1:8080/block

curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "hash": "hash" } }' http://127.0.0.1:8080/block
```
### RPC Endpoint /block/transaction test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "block_identifier": { "index": index, "hash": "hash" }, "transaction_identifier": { "hash": "address" } }' http://127.0.0.1:8080/block/transaction
```
### RPC Endpoint /construction/submit test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet", }, "signed_transaction": "data" }' http://127.0.0.1:8080/construction/submit
```
### RPC Endpoint /construction/metadata test examplec
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "options": {}, "public_keys": [{ "hex_bytes": "string", "curve_type": "secp256k1" }] }' http://127.0.0.1:8080/construction/metadata
```
### RPC Endpoint /construction/combine test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "unsigned_transaction": "data", "signatures": [ { "signing_payload": { "hex_bytes": "data", }, "public_key": { "hex_bytes": "data", "curve_type": "secp256k1" }, "signature_type": "ecdsa", "hex_bytes": "data" } ] }' http://127.0.0.1:8080/construction/combine
```
### RPC Endpoint /construction/derive test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "public_key": { "hex_bytes": "data", "curve_type": "secp256k1" } }' http://127.0.0.1:8080/construction/derive
```
### RPC Endpoint /construction/hash test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "signed_transaction": "data" }' http://127.0.0.1:8080/construction/hash
```
### RPC Endpoint /construction/parse test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "signed": true, "transaction": "data" }' http://127.0.0.1:8080/construction/parse
```
### RPC Endpoint /construction/payloads test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "operations": [ { "operation_identifier": { "index": index, "network_index": index }, "type": "Transfer" } ] }' http://127.0.0.1:8080/construction/payloads
```
### RPC Endpoint /construction/preprocess test example
```
curl -H 'Content-Type: application/json' --data '{ "network_identifier": { "blockchain": "Findora", "network": "Prinet" }, "operations": [ { "operation_identifier": { "index": index, "network_index": index }, "type": "Transfer" } ] }' http://127.0.0.1:8080/construction/preprocess
```

## Rosetta Cli Tools test examples
### Data Api test example
```
export START_BLOCK=2
export ROSETTA_CONFIGURATION_FILE=./rosetta-cli-conf/prinet/config_data.json
rosetta-cli check:data --start-block $START_BLOCK
```

### Construction Api test example
export ROSETTA_CONFIGURATION_FILE=./rosetta-cli-conf/prinet/config_data.json
rosetta-cli check:construction