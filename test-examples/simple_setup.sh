#!/bin/bash

if [ -z $RPCURL ]; then
    export RPCURL=http://127.0.0.1:8545
fi

if [ -z $PORT ]; then
    export PORT=8080
else
    if [ -n "$(echo $PORT | sed 's/[0-9]//g')" ]; then
        echo "please input corret rosetta port"
        exit 1
    fi
fi

export MODE=ONLINE
export NETWORK=PRINET

./findora-rosetta run
