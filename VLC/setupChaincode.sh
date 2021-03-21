# install chaincode
# Install code on manufacturer peer
echo "Installing tfbc chaincode to peer0.manufacturer.vlbc.com:"
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlbc.com/users/Admin@manufacturer.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlbc.com:7051" cli peer chaincode install -n carRegistration -v 2.0 -p github.com/carRegistration -l golang
echo "Installed carRegistration chaincode to peer0.manufacturer.vlbc.com:"

# Install code on dealer peer
echo "Installing carRegistration chaincode to peer0.dealer.vlbc.com:"
docker exec -e "CORE_PEER_LOCALMSPID=DealerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dealer.vlbc.com/users/Admin@dealer.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.dealer.vlbc.com:7051" cli peer chaincode install -n carRegistration -v 2.0 -p github.com/carRegistration -l golang
echo "Installed carRegistration chaincode to peer0.dealer.vlbc.com:"

# Install code on insurer peer
echo "Installing carRegistration chaincode to peer0.insurer.vlbc.com:"
sleep 5
docker exec -e "CORE_PEER_LOCALMSPID=InsurerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insurer.vlbc.com/users/Admin@insurer.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.insurer.vlbc.com:7051" cli peer chaincode install -n carRegistration -v 2.0 -p github.com/carRegistration -l golang
echo "Installed carRegistration chaincode to peer0.insurer.vlbc.com:"

# Install code on rto peer
echo "Installing carRegistration chaincode to peer0.rto.vlbc.com:"
docker exec -e "CORE_PEER_LOCALMSPID=RTOMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rto.vlbc.com/users/Admin@rto.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.rto.vlbc.com:7051" cli peer chaincode install -n carRegistration -v 2.0 -p github.com/carRegistration -l golang
echo "Installed carRegistration chaincode to peer0.rto.vlbc.com:"
sleep 5

echo "Instantiating carRegistration chaincode.."
sudo docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlbc.com/users/Admin@manufacturer.vlbc.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlbc.com:7051" cli peer chaincode instantiate -o orderer.vlbc.com:7050 -C crchannel -n carRegistration -l golang -v 2.0 -c '{"Args":[""]}'
echo "Instantiated carRegistration chaincode."
echo "Following is the docker network....."

docker ps
