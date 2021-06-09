#!/bin/bash

docker-compose down

# REMOVE the dev- container images also - TBD
docker rm $(docker ps -a -q)            &> /dev/null
docker rmi $(docker images dev-* -q)    &> /dev/null
sudo rm -rf $HOME/ledgers/ca &> /dev/null

docker-compose up -d

SLEEP_TIME=3s
echo    '========= Submitting txn for channel creation as AstuAdmin ============'
export CHANNEL_TX_FILE=./config/docs-channel.tx
export ORDERER_ADDRESS=orderer.astu.com:7050
# export FABRIC_LOGGING_SPEC=DEBUG
export CORE_PEER_LOCALMSPID=AstuMSP
export CORE_PEER_MSPCONFIGPATH=$PWD/client/astu/admin/msp
export CORE_PEER_ADDRESS=astu-admin-peer1.astu.com:7051
peer channel create -o $ORDERER_ADDRESS -c docschannel -f ./config/docschannel.tx

echo    '========= Joining the astu-peer1 to Docs channel ============'
AIRLINE_CHANNEL_BLOCK=./docschannel.block
export CORE_PEER_ADDRESS=astu-admin-peer1.astu.com:7051
peer channel join -o $ORDERER_ADDRESS -b $AIRLINE_CHANNEL_BLOCK
# Update anchor peer on channel for astu
# sleep  3s
sleep $SLEEP_TIME
ANCHOR_UPDATE_TX=./config/docs-anchor-update-astu.tx
peer channel update -o $ORDERER_ADDRESS -c docschannel -f $ANCHOR_UPDATE_TX

echo    '========= Joining the astu-service-peer1 to Docs channel ============'
# peer channel fetch config $AIRLINE_CHANNEL_BLOCK -o $ORDERER_ADDRESS -c docschannel
export CORE_PEER_LOCALMSPID=Astu-ServiceMSP
ORG_NAME=astu-service.com
export CORE_PEER_ADDRESS=astu-service-peer1.astu.com:8051
export CORE_PEER_MSPCONFIGPATH=$PWD/client/astu-service/admin/msp
peer channel join -o $ORDERER_ADDRESS -b $AIRLINE_CHANNEL_BLOCK
# Update anchor peer on channel for astu-service
sleep  $SLEEP_TIME
ANCHOR_UPDATE_TX=./config/docs-anchor-update-astu-service.tx
peer channel update -o $ORDERER_ADDRESS -c docschannel -f $ANCHOR_UPDATE_TX

