#!/bin/bash

if [ -z $RPCURL ]; then
    export RPCURL=https://prod-testnet.prod.findora.org:8545
fi

if [ -z $PORT ]; then
    export PORT=8080
else
    if [ -n "$(echo $PORT | sed 's/[0-9]//g')" ]; then
        echo "please input corret rosetta port"
        exit 1
    fi
fi

if [ "$#" -ne 1 ]; then
    echo "please input corret network" >&2
    exit 1
fi

if [ "$1" = "mainnet" ]; then
    export RPCURL=https://prod-mainnet.prod.findora.org:8545
elif [ "$1" = "anvil" ]; then
    export RPCURL=https://prod-testnet.prod.findora.org:8545
elif [ "$1" = "qa02" ]; then
    export RPCURL=https://dev-qa02.dev.findora.org:8545
elif [ "$1" = "prinet" ]; then
    export RPCURL=http://127.0.0.1:8545
else 
    echo "please input corret network"
    exit 1
fi


network=$(echo $1|sed 's/\b[a-z]*/\U&/g')


export MODE=ONLINE
export NETWORK=$network

export http_proxy=
export https_proxy=

./findora-rosetta run
