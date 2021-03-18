dev-init.sh -s
(coachDb - http://localhost:5984/_utils/)
. set-env.sh acme

set-chain-env.sh -p docs -n docs -v 1.0 -c '{"Args":["init"]}' -C airlinechannel

chain.sh install -p

chain.sh instantiate

!!UPGRADE!!!
chain.sh upgrade-auto