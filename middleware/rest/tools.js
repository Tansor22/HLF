const fs = require('fs')
const yaml = require('js-yaml')
const {Gateway, FileSystemWallet} = require('fabric-network')

module.exports = {
    setupGateway: async function (walletPath, profilePath, userId) {
        let connectionProfile = yaml.load(fs.readFileSync(profilePath, 'utf8'));
        const wallet = new FileSystemWallet(walletPath)

        let connectionOptions = {
            identity: userId,
            wallet: wallet,
            discovery: {enabled: false, asLocalhost: true}
            /*** Uncomment lines below to disable commit listener on submit ****/
            // , eventHandlerOptions: {
            //     strategy: null
            // }
        }
        let gateway = new Gateway();
        await gateway.connect(connectionProfile, connectionOptions)
        return gateway
    }
}