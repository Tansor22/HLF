const fs = require('fs');
const yaml = require('js-yaml');
var express = require('express');
var app = express();
const { Gateway, FileSystemWallet} = require('fabric-network');

// Constants for profile
const CONNECTION_PROFILE_PATH = '../profiles/dev-connection.yaml'
// Path to the wallet
const FILESYSTEM_WALLET_PATH = '../gateway/user-wallet'
// Identity context used
const USER_ID = 'Admin@acme.com'
// Channel name
const NETWORK_NAME = 'airlinechannel'
// Chaincode
const CONTRACT_ID = "erc20"



const gateway = new Gateway();

main()

async function main() {
    await setupGateway()
    app.get('/accessHLF', async function (req, res) {

        let network = await gateway.getNetwork(NETWORK_NAME)
        const contract = await network.getContract(CONTRACT_ID);
        try{
            // Submit the transaction
            let response = await contract.submitTransaction('transfer', 'john','sam','2')
            console.log("Submit Response=",response.toString())
            res.setHeader('Content-Type', 'application/json');
            res.end(response)
        } catch(e){
            console.log(e)
        }
    })
    var server = app.listen(8081, function () {
        var host = server.address().address
        var port = server.address().port
        console.log("REST server listening at http://%s:%s", host, port)
    })
}

async function setupGateway() {
    let connectionProfile = yaml.safeLoad(fs.readFileSync(CONNECTION_PROFILE_PATH, 'utf8'));
    const wallet = new FileSystemWallet(FILESYSTEM_WALLET_PATH)

    let connectionOptions = {
        identity: USER_ID,
        wallet: wallet,
        discovery: { enabled: false, asLocalhost: true }
        /*** Uncomment lines below to disable commit listener on submit ****/
        // , eventHandlerOptions: {
        //     strategy: null
        // }
    }

    await gateway.connect(connectionProfile, connectionOptions)
}
