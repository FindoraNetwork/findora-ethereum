#!/bin/bash
rm -rf test-cli
mkdir -p test-cli/rosetta-data
export ROSETTA_CONFIGURATION_FILE=./rosetta-cli-conf/prinet/config.json
rosetta-cli check:construction
