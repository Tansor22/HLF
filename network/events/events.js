/**
 * Launches the listener for events
 * 
 * node events   [event-type="block" | "txn" | "chaincode"]
 *               [chaincode-name="cc-name"]
 *               [chaincode-event="cc-event"]
 *               [crypto-type=defaut="cgen" | "ca"] 
 * 
 * https://fabric-sdk-node.github.io/release-1.3/tutorial-network-config.html
 * 
 * https://fabric-sdk-node.github.io/tutorial-network-config.html
 * 
 */
'use strict';

var Client = require('fabric-client');
const fs = require('fs');
const yaml = require('node-yaml')

//const cryptoFolder="../crypto/crypto-config/peerOrganizations"

/**
 * Arguments
 */
var listenerType=process.argv[2]
var ccName=process.argv[3]
var ccEvent=process.argv[4]
var channelId=process.argv[5]
var cryptoType=process.argv[6]
var cryptoFolder=process.argv[7]


console.log("Launching Listener: "," type=",listenerType," cc Name=",ccName," event Name=",ccEvent)
// console.log("---->",cryptoType, channelId)

/**
 * Set up the nw config file
 */
var obj = yaml.readSync('./nw-config.template.yaml')
createYAMLCryptogen(obj)


/** Launch the listener */
Client.setConfigSetting('nw-config','./nw-config.yaml');
Client.setConfigSetting('astu-con-profile','./astu-client.yaml');

let client = Client.loadFromConfig(Client.getConfigSetting('nw-config'));
client.loadFromConfig(Client.getConfigSetting('astu-con-profile'));


async function launch(){
    await client.initCredentialStores();

    if(listenerType=="block"){
        blockListener()
    } else if(listenerType=="chaincode") {
        chaincodeEventListener()
    }
}

launch()



function    blockListener() {

    var channel = client.getChannel(channelId)

    var channel_event_hub = channel.newChannelEventHub('astu-admin-peer1.astu.com');

    var block_reg = channel_event_hub.registerBlockEvent((block) => {
        console.log('Successfully received the block number=',block.header.number);
        // console.log(block.header.number)
        // console.log(block.data.data[0].payload.data.actions[0].payload.action.endorsements)
    }, (error)=> {
        console.log('Failed to receive the block event ::'+error);
        console.log("Error:", error)
    }
    );

    channel_event_hub.connect(true); 
}

function chaincodeEventListener() {
    var channel = client.getChannel(channelId)

    var channel_event_hub = channel.newChannelEventHub('astu-admin-peer1.astu.com');

    var cc_reg = channel_event_hub.registerChaincodeEvent(ccName,ccEvent,(event, block_num, txnid, status)=>{
        var payload = event.payload
        let INFO_SYMBOL='\u2705'
        console.log(INFO_SYMBOL, ' event=',event.event_name,'\n block#',block_num,"\n status=",status,"\n payload=",payload.toString('utf8') )
    }, (error)=>{
        console.log("Error:", error)
    })
    channel_event_hub.connect(true);
}





// Generates the YAML based on the 
function createYAMLCryptogen(obj){
    let astuCert = genCertPathCryptogen('astu')
    let astuPk = genPkCryptogen("astu")
    obj.organizations.Astu.signedCert.path=astuCert
    obj.organizations.Astu.adminPrivateKey.path=astuPk
    let astu-serviceCert = genCertPathCryptogen('astu-service')
    let astu-servicePk = genPkCryptogen("astu-service")
    obj.organizations.Astu-Service.signedCert.path=astu-serviceCert
    obj.organizations.Astu-Service.adminPrivateKey.path=astu-servicePk

    // console.log(astuCert)
    yaml.writeSync("nw-config.yaml", obj)

    console.log("+ Successfully generated Network config YAML")
}

function genCertPathCryptogen(org){ 
    //astu-service.com/users/Admin@astu-service.com/msp/signcerts/Admin@astu-service.com-cert.pem"
    var certPath=cryptoFolder+"/"+org+".com/users/Admin@"+org+".com/msp/signcerts/Admin@"+org+".com-cert.pem"
    return certPath
}

// looks for the PK files in the org folder
function    genPkCryptogen(org){
    // ../crypto/crypto-config/peerOrganizations/astu-service.com/users/Admin@astu-service.com/msp/keystore/05beac9849f610ad5cc8997e5f45343ca918de78398988def3f288b60d8ee27c_sk
    var pkFolder=cryptoFolder+"/"+org+".com/users/Admin@"+org+".com/msp/keystore"
    let pkfile=""
    fs.readdirSync(pkFolder).forEach(file => {
        // console.log(file);
        // return the first file
        pkfile = file
        return
    })

    return pkFolder+"/"+pkfile
}