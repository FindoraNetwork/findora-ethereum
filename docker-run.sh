#!/bin/bash
cd /root/
rm -rf ./.tendermint
mkdir -p /root/.____fn_config____
mkdir -p /root/.findora

echo "zoo nerve assault talk depend approve mercy surge bicycle ridge dismiss satoshi boring opera next fat cinnamon valley office actor above spray alcohol giant" > /root/.findora/mnenomic.key
echo "/root/.findora/mnenomic.key" > /root/.____fn_config____/mnemonic
echo "http://localhost" > /root/.____fn_config____/serv_addr

findorad init
nohup findorad node -q --enable-eth-api-service > /root/findorad.log 2>&1 &
sleep 5

setup -S 'http://localhost'
stt init -s
fn contract-deposit -a 0xA405BA2b64DC04466E0f23487FD1c4A084787326 -n 1000000000000000
fn contract-deposit -a 0xA405BA2b64DC04466E0f23487FD1c4A084787326 -n 1000000000000000
fn contract-deposit -a 0xA405BA2b64DC04466E0f23487FD1c4A084787326 -n 1000000000000000
fn contract-deposit -a 0xA405BA2b64DC04466E0f23487FD1c4A084787326 -n 1000000000000000
fn contract-deposit -a 0xA405BA2b64DC04466E0f23487FD1c4A084787326 -n 1000000000000000

export PORT=8080
export RPCURL=http://127.0.0.1:8545
export NETWORK=PRINET

if [ ! -n "$MODE" ]; then
 export MODE=ONLINE
fi


nohup /root/findora-rosetta run  > /root/findora-rosetta.log 2>&1 &
/bin/bash -c "while true;do echo hello;sleep 50000;done"
