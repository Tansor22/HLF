If you get a message that certs have expired execute the script:

./init.sh

ca-dev-init.sh


export FABRIC_CFG_PATH=$PWD/config
configtxgen -outputBlock  ./config/docsgenesis.block -channelID ordererchannel  -profile DocsOrdererGenesis
configtxgen -outputCreateChannelTx  ./config/docschannel.tx -channelID docschannel  -profile DocsChannel