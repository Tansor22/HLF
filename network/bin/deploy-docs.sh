#!/bin/bash

dev-init.sh -s
echo    "Installing the chaincode docs"
.    set-env.sh    acme
set-chain-env.sh       -n docs  -v 1.0   -p  docs
chain.sh install -p

echo    "Instantiating..."
set-chain-env.sh        -c   '{"Args":["init"]}'
chain.sh  instantiate

echo "Done."