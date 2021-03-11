const fs = require('fs');
const path = require('path');
const {FileSystemWallet, X509WalletMixin} = require('fabric-network');

// Location of the crypto for the dev environment
const CRYPTO_CONFIG = path.resolve(__dirname, '../../network/crypto/crypto-config');
const CRYPTO_CONFIG_PEER_ORGS = path.join(CRYPTO_CONFIG, 'peerOrganizations')

// Folder for creating the wallet - All identities written under this
const WALLET_FOLDER = './user-wallet'
const wallet = new FileSystemWallet(WALLET_FOLDER);
const supportedActions = ['list', 'add', 'export']
if (process.argv.length > 2) {
    let action = process.argv[2]
    if (!supportedActions.includes(action)) {
        console.log("Wrong action specified: Possible values are"
            + supportedActions.map(it => " " + "'" + it + "'"))
        return
    }
    if (action === 'list') {
        console.log("List of identities in wallet:")
        listIdentities()
    } else if (action === 'add' || action === 'export') {
        if (process.argv.length < 5) {
            console.log("For 'add' & 'export' - Org & User are needed!!!")
            process.exit(1)
        }
        if (action === 'add') {
            addToWallet(process.argv[3], process.argv[4])
            console.log('Done adding/updating.')
        } else {
            exportIdentity(process.argv[3], process.argv[4])
        }
    }
} else {
    console.log("No action specified!!!")
}



async function addToWallet(org, user) {
    // Read the cert & key file content
    try {
        // Read the certificate file content
        var cert = readCertCryptogen(org, user)

        // Read the keyfile content
        var key = readPrivateKeyCryptogen(org, user)

    } catch (e) {
        console.log("Error reading certificate or key!!! " + org + "/" + user)
        process.exit(1)
    }

    // Create the MSP ID
    let mspId = createMSPId(org)

    // Create the label
    const identityLabel = createIdentityLabel(org, user);

    // Create the X509 identity 
    const identity = X509WalletMixin.createIdentity(mspId, cert, key);

    // Add to the wallet
    await wallet.import(identityLabel, identity);
}

async function listIdentities() {
    console.log("Identities in Wallet:")

    // Retrieve the identities in folder
    let list = await wallet.list()

    for (var i = 0; i < list.length; i++) {
        console.log((i + 1) + '. ' + list[i].label)
    }
}

async function exportIdentity(org, user) {
    // Label is used for identifying the identity in wallet
    let label = createIdentityLabel(org, user)

    // To retrieve execute export
    let identity = await wallet.export(label)

    if (identity == null) {
        console.log(`Identity ${user} for ${org} Org Not found!!!`)
    } else {
        // Prints all attributes : label, Key, Cert
        console.log(identity)
    }
}

function readCertCryptogen(org, user) {
    let certPath = CRYPTO_CONFIG_PEER_ORGS + "/" + org + ".com/users/" + user + "@" + org + ".com/msp/signcerts/" + user + "@" + org + ".com-cert.pem";
    return fs.readFileSync(certPath).toString()
}

function readPrivateKeyCryptogen(org, user) {
    const pkFolder = CRYPTO_CONFIG_PEER_ORGS + "/" + org + ".com/users/" + user + "@" + org + ".com/msp/keystore";
    fs.readdirSync(pkFolder).forEach(file => {
        // return the first file
        pkfile = file
        return
    })

    return fs.readFileSync(pkFolder + "/" + pkfile).toString()
}

function createMSPId(org) {
    return org.charAt(0).toUpperCase() + org.slice(1) + 'MSP'
}

function createIdentityLabel(org, user) {
    return user + '@' + org + '.com';
}