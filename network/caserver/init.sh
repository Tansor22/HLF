docker-compose -f docker-compose-ca.yaml down
rm -rf ./server/*
rm -rf ./client/*
cp fabric-ca-server-config.yaml ./server
docker-compose -f docker-compose-ca.yaml up -d

sleep 3s

# Bootstrap enrollment
export FABRIC_CA_CLIENT_HOME=$PWD/client/caserver/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:7054


######################
# Admin registration #
######################
echo "Registering: astu-admin"
ATTRIBUTES='"hf.Registrar.Roles=peer,user,client","hf.AffiliationMgr=true","hf.Revoker=true","hf.Registrar.Attributes=*"'
fabric-ca-client register --id.type client --id.name astu-admin --id.secret adminpw --id.affiliation astu --id.attrs $ATTRIBUTES

# 3. Register astu-service-admin
echo "Registering: astu-service-admin"
ATTRIBUTES='"hf.Registrar.Roles=peer,user,client","hf.AffiliationMgr=true","hf.Revoker=true","hf.Registrar.Attributes=*"'
fabric-ca-client register --id.type client --id.name astu-service-admin --id.secret adminpw --id.affiliation astu-service --id.attrs $ATTRIBUTES

# 4. Register orderer-admin
echo "Registering: orderer-admin"
ATTRIBUTES='"hf.Registrar.Roles=orderer"'
fabric-ca-client register --id.type client --id.name orderer-admin --id.secret adminpw --id.affiliation orderer --id.attrs $ATTRIBUTES


####################
# Admin Enrollment #
####################
export FABRIC_CA_CLIENT_HOME=$PWD/client/astu/admin
fabric-ca-client enroll -u http://astu-admin:adminpw@localhost:7054
mkdir -p $FABRIC_CA_CLIENT_HOME/msp/admincerts
cp $FABRIC_CA_CLIENT_HOME/../../caserver/admin/msp/signcerts/*  $FABRIC_CA_CLIENT_HOME/msp/admincerts

export FABRIC_CA_CLIENT_HOME=$PWD/client/astu-service/admin
fabric-ca-client enroll -u http://astu-service-admin:adminpw@localhost:7054
mkdir -p $FABRIC_CA_CLIENT_HOME/msp/admincerts
cp $FABRIC_CA_CLIENT_HOME/../../caserver/admin/msp/signcerts/*  $FABRIC_CA_CLIENT_HOME/msp/admincerts

export FABRIC_CA_CLIENT_HOME=$PWD/client/orderer/admin
fabric-ca-client enroll -u http://orderer-admin:adminpw@localhost:7054
mkdir -p $FABRIC_CA_CLIENT_HOME/msp/admincerts
cp $FABRIC_CA_CLIENT_HOME/../../caserver/admin/msp/signcerts/*  $FABRIC_CA_CLIENT_HOME/msp/admincerts

#################
# Org MSP Setup #
#################
# Path to the CA certificate
ROOT_CA_CERTIFICATE=./server/ca-cert.pem
mkdir -p ./client/orderer/msp/admincerts
mkdir ./client/orderer/msp/cacerts
mkdir ./client/orderer/msp/keystore
cp $ROOT_CA_CERTIFICATE ./client/orderer/msp/cacerts
cp ./client/orderer/admin/msp/signcerts/* ./client/orderer/msp/admincerts   

mkdir -p ./client/astu/msp/admincerts
mkdir ./client/astu/msp/cacerts
mkdir ./client/astu/msp/keystore
cp $ROOT_CA_CERTIFICATE ./client/astu/msp/cacerts
cp ./client/astu/admin/msp/signcerts/* ./client/astu/msp/admincerts

mkdir -p ./client/astu-service/msp/admincerts
mkdir ./client/astu-service/msp/cacerts
mkdir ./client/astu-service/msp/keystore
cp $ROOT_CA_CERTIFICATE ./client/astu-service/msp/cacerts
cp ./client/astu-service/admin/msp/signcerts/* ./client/astu-service/msp/admincerts

######################
# Orderer Enrollment #
######################
export FABRIC_CA_CLIENT_HOME=$PWD/client/orderer/admin
fabric-ca-client register --id.type orderer --id.name orderer --id.secret adminpw --id.affiliation orderer 
export FABRIC_CA_CLIENT_HOME=$PWD/client/orderer/orderer
fabric-ca-client enroll -u http://orderer:adminpw@localhost:7054
cp -a $PWD/client/orderer/admin/msp/signcerts  $FABRIC_CA_CLIENT_HOME/msp/admincerts

####################
# Peer Enrollments #
####################
export FABRIC_CA_CLIENT_HOME=$PWD/client/astu/admin
fabric-ca-client register --id.type peer --id.name astu-peer1 --id.secret adminpw --id.affiliation astu
export FABRIC_CA_CLIENT_HOME=$PWD/client/astu/peer1
fabric-ca-client enroll -u http://astu-peer1:adminpw@localhost:7054
cp -a $PWD/client/astu/admin/msp/signcerts  $FABRIC_CA_CLIENT_HOME/msp/admincerts

export FABRIC_CA_CLIENT_HOME=$PWD/client/astu-service/admin
fabric-ca-client register --id.type peer --id.name astu-service-peer1 --id.secret adminpw --id.affiliation astu-service
export FABRIC_CA_CLIENT_HOME=$PWD/client/astu-service/peer1
fabric-ca-client enroll -u http://astu-service-peer1:adminpw@localhost:7054
cp -a $PWD/client/astu-service/admin/msp/signcerts  $FABRIC_CA_CLIENT_HOME/msp/admincerts


##############################
# User Enrollments Astu only #
##############################
export FABRIC_CA_CLIENT_HOME=$PWD/client/astu/admin
ATTRIBUTES='"hf.AffiliationMgr=false:ecert","hf.Revoker=false:ecert","app.accounting.role=manager:ecert","department=accounting:ecert"'
fabric-ca-client register --id.type user --id.name mary --id.secret pw --id.affiliation astu --id.attrs $ATTRIBUTES
export FABRIC_CA_CLIENT_HOME=$PWD/client/astu/mary
fabric-ca-client enroll -u http://mary:pw@localhost:7054
cp -a $PWD/client/astu/admin/msp/signcerts  $FABRIC_CA_CLIENT_HOME/msp/admincerts

export FABRIC_CA_CLIENT_HOME=$PWD/client/astu/admin
ATTRIBUTES='"hf.AffiliationMgr=false:ecert","hf.Revoker=false:ecert","app.accounting.role=accountant:ecert","department=accounting:ecert"'
fabric-ca-client register --id.type user --id.name john --id.secret pw --id.affiliation astu --id.attrs $ATTRIBUTES
export FABRIC_CA_CLIENT_HOME=$PWD/client/astu/john
fabric-ca-client enroll -u http://john:pw@localhost:7054
cp -a $PWD/client/astu/admin/msp/signcerts  $FABRIC_CA_CLIENT_HOME/msp/admincerts

export FABRIC_CA_CLIENT_HOME=$PWD/client/astu/admin
ATTRIBUTES='"hf.AffiliationMgr=false:ecert","hf.Revoker=false:ecert","department=logistics:ecert","app.logistics.role=specialis:ecert"'
fabric-ca-client register --id.type user --id.name anil --id.secret pw --id.affiliation astu --id.attrs $ATTRIBUTES
export FABRIC_CA_CLIENT_HOME=$PWD/client/astu/anil
fabric-ca-client enroll -u http://anil:pw@localhost:7054
cp -a $PWD/client/astu/admin/msp/signcerts  $FABRIC_CA_CLIENT_HOME/msp/admincerts

# Shutdown CA
docker-compose -f docker-compose-ca.yaml down

# Setup network config
export FABRIC_CFG_PATH=$PWD/config
configtxgen -outputBlock  ./config/orderer/docs-genesis.block -channelID ordererchannel  -profile DocsOrdererGenesis
configtxgen -outputCreateChannelTx  ./config/docschannel.tx -channelID docschannel  -profile DocsChannel

ANCHOR_UPDATE_TX=./config/docs-anchor-update-astu.tx
configtxgen -profile DocsChannel -outputAnchorPeersUpdate $ANCHOR_UPDATE_TX -channelID docschannel -asOrg AstuMSP

ANCHOR_UPDATE_TX=./config/docs-anchor-update-astu-service.tx
configtxgen -profile DocsChannel -outputAnchorPeersUpdate $ANCHOR_UPDATE_TX -channelID docschannel -asOrg Astu-ServiceMSP
