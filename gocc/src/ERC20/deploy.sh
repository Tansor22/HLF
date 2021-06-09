#!/bin/bash

echo    "Installing the chaincode ERC20"
.    set-env.sh    astu
set-chain-env.sh       -n erc20  -v 1.0   -p  ERC20
chain.sh install -p

echo    "Instantiating..."
set-chain-env.sh        -c   '{"Args":["init","ACFT","1000", "A Cloud Fan Token!!!","john"]}'
chain.sh  instantiate

echo "Done."