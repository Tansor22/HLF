#!/bin/bash

dev-init.sh -s -e
# -e means explorer
# OR
# exp-start.sh
# - exp-stop.sh
echo    "Installing the chaincode docs"
.    set-env.sh    astu
set-chain-env.sh       -n docs  -v 1.0   -p  docs
chain.sh install -p

echo    "Instantiating..."
set-chain-env.sh        -c   '{"Args":["init"]}'
chain.sh  instantiate

echo "Done."