#!/bin/bash
rm cli_tests/rosetta-data -rf
mkdir -p cli_tests/rosetta-data
export ROSETTA_CONFIGURATION_FILE=./rosetta-cli-conf/prinet/config.json
rosetta-cli check:construction
