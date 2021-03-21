rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile VLBCOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile CRChannel -outputCreateChannelTx ./config/crchannel.tx -channelID crchannel
