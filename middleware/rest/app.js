const fs = require('fs');
const yaml = require('js-yaml');
var express = require('express');
var bodyParser = require('body-parser');
var app = express();
const {Gateway, FileSystemWallet} = require('fabric-network');

// Constants for profile
const CONNECTION_PROFILE_PATH = '../profiles/dev-connection.yaml'
// Path to the wallet
const FILESYSTEM_WALLET_PATH = '../gateway/user-wallet'
// Identity context used
const USER_ID = 'Admin@acme.com'
// Channel name
const NETWORK_NAME = 'airlinechannel'
// Chaincode
const CONTRACT_ID = "docs"


const gateway = new Gateway();

main()

async function main() {
    await setupGateway()
    var jsonParser = bodyParser.json()
    app.post('/newDoc', jsonParser, async function (req, res) {
       // todo add user selection in each func
        let body = req.body
        console.log("Request = " + body)
        let network = await gateway.getNetwork(NETWORK_NAME)
        const contract = await network.getContract(CONTRACT_ID);
        await invoke(req, res, contract, 'new-doc', body.org, body.content, JSON.stringify(body.signsRequired))
    })
    app.post('/getDocs', jsonParser, async function (req, res) {
        let body = req.body
        console.log("Request = " + body)
        let network = await gateway.getNetwork(NETWORK_NAME)
        const contract = await network.getContract(CONTRACT_ID);
        await query(req, res, contract, 'get-docs', body.orgName)
    })
    app.post('/signDoc', jsonParser, async function (req, res) {
        let body = req.body
        console.log("Request = " + body)
        let network = await gateway.getNetwork(NETWORK_NAME)
        const contract = await network.getContract(CONTRACT_ID);
        await invoke(req, res, contract, 'sign-doc', body.documentId, JSON.stringify(body.signs))
    })
    var server = app.listen(8081, function () {
        var host = server.address().address
        var port = server.address().port
        console.log("REST server listening at http://%s:%s", host, port)
    })
}
async function query(req, res, contract, func, ... args){
    try{
        let response = await contract.evaluateTransaction(func, ... args)
        console.log(`Query Response=${response.toString()}`)
        res.setHeader('Content-Type', 'application/json');
        res.end(response)
    } catch(e){
        console.log(e)
        res.end("FAILED, see logs of middleware")
    }
}

async function invoke(req, res, contract, func, ... args){
    try{
        let response = await contract.submitTransaction(func, ... args)
        console.log(`Query Response=${response.toString()}`)
        res.setHeader('Content-Type', 'application/json');
        res.end(response)
    } catch(e){
        console.log(e)
        res.end("FAILED, see logs of middleware")
    }
}

async function setupGateway() {
    let connectionProfile = yaml.safeLoad(fs.readFileSync(CONNECTION_PROFILE_PATH, 'utf8'));
    const wallet = new FileSystemWallet(FILESYSTEM_WALLET_PATH)

    let connectionOptions = {
        identity: USER_ID,
        wallet: wallet,
        discovery: {enabled: false, asLocalhost: true}
        /*** Uncomment lines below to disable commit listener on submit ****/
        // , eventHandlerOptions: {
        //     strategy: null
        // }
    }

    await gateway.connect(connectionProfile, connectionOptions)
}
