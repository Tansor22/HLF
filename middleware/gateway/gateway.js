const fs = require('fs');
const yaml = require('js-yaml');
const { Gateway, FileSystemWallet} = require('fabric-network');

// Constants for profile
const CONNECTION_PROFILE_PATH = '../profiles/dev-connection.yaml'
// Path to the wallet
const FILESYSTEM_WALLET_PATH = './user-wallet'
// Identity context used
const USER_ID = 'Admin@astu.com'
// Channel name
const NETWORK_NAME = 'docschannel'
// Chaincode
const CONTRACT_ID = "erc20"



const gateway = new Gateway();

main()

async function main() {
    await setupGateway()
    let network = await gateway.getNetwork(NETWORK_NAME)
    const contract = await network.getContract(CONTRACT_ID);
    await queryContract(contract)
    await submitTxnContract(contract)
}

async function queryContract(contract){
    try{
        let response = await contract.evaluateTransaction('balanceOf', 'john')
        console.log(`Query Response=${response.toString()}`)
    } catch(e){
        console.log(e)
    }
}

async function submitTxnContract(contract){
    try{
        let response = await contract.submitTransaction('transfer', 'john','sam','2')
        console.log("Submit Response=",response.toString())
    } catch(e){
        console.log(e)
    }
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