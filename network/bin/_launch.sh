#PEER_MODE=net
#Command=dev-init.sh -e -s 
#Generated: Wed Jun  9 03:37:04 UTC 2021 
docker-compose  -f ./compose/docker-compose.base.yaml     -f ./compose/docker-compose.couchdb.yaml   -f ./compose/docker-compose.explorer.yaml    up -d --remove-orphans
