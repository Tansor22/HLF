dev-init.sh -s
(coachDb - http://localhost:5984/_utils/)
. set-env.sh astu

set-chain-env.sh -p docs -n docs -v 1.0 -c '{"Args":["init"]}' -C docschannel

chain.sh install -p

chain.sh instantiate

!!UPGRADE!!!
chain.sh upgrade-auto


vagrant up
vagrant ssh

deploy-docs.sh