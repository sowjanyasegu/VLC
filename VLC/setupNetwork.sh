echo "Setting up the network.."

echo "Creating channel genesis block.."

# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlbc.com/users/Admin@manufacturer.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlbc.com:7051" cli peer channel create -o orderer.vlbc.com:7050 -c crchannel -f /etc/hyperledger/configtx/crchannel.tx


sleep 5

echo "Channel genesis block created."

echo "peer0.manufacturer.vlbc.com joining the channel..."
# Join peer0.manufacturer.vlbc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlbc.com/users/Admin@manufacturer.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlbc.com:7051" cli peer channel join -b crchannel.block

echo "peer0.manufacturer.vlbc.com joined the channel"

echo "peer0.dealer.vlbc.com joining the channel..."

# Join peer0.dealer.vlbc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=DealerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dealer.vlbc.com/users/Admin@dealer.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.dealer.vlbc.com:7051" cli peer channel join -b crchannel.block

echo "peer0.dealer.vlbc.com joined the channel"

echo "peer0.insurer.vlbc.com joining the channel..."
# Join peer0.insurer.vlbc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=InsurerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insurer.vlbc.com/users/Admin@insurer.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.insurer.vlbc.com:7051" cli peer channel join -b crchannel.block
sleep 5

echo "peer0.rto.vlbc.com joined the channel"

echo "peer0.rto.vlbc.com joining the channel..."
# Join peer0.rto.vlbc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=RTOMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rto.vlbc.com/users/Admin@rto.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.rto.vlbc.com:7051" cli peer channel join -b crchannel.block
sleep 5

echo "peer0.rto.vlbc.com joined the channel"

echo "Following is the docker network....."

docker ps

