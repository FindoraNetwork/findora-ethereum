#!/bin/bash
if [ "$#" -ne 1 ]; then
    echo "please input correct test start block index for data api 1" >&2
    exit 1
fi

if [ -n "$(echo $1 | sed 's/[0-9]//g')" ]; then
    echo "please input correct test start block index for data api 2 " >&2
    exit 1
fi

rm cli_tests/rosetta-data -rf
mkdir -p cli_tests/rosetta-data
export ROSETTA_CONFIGURATION_FILE=./rosetta-cli-conf/prinet/config.json
rosetta-cli check:data --start-block $1
