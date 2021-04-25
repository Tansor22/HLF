const fs = require('fs')
const yaml = require('js-yaml')
const {Gateway, FileSystemWallet} = require('fabric-network')
//const richConsole = require('rich-console')

module.exports = {
   /* logAndSend: (request, response, status, data) => {
        // log
        let msg = `\tEndpoint: ${request.path}\n\tRequest: ${JSON.stringify(request.body)}\n\tResponse (Code:${status}): ${JSON.stringify(data)}\n`
        let richMsg = status === 200 ? '<green>' + msg + '</green>>' : '<red>' + msg + '</red>>'
        richConsole.log(richMsg)
        // send
        response.status(status).send(data)
    },
    logAndSendOk: (request, response, data) => {
        module.exports.logAndSend(request, response, 200, data ? data : {result: 'Ok'})
    },
    logAndSendError: () => {
        const logAndSendErrorCustom = (request, response,  status, error, details) => {
            module.exports.logAndSend(request, response, status, {error, details})
        }
        const logAndSendError500 = (request, response, error, details) => {
            module.exports.logAndSend(request, 500, {error, details})
        }
        if (arguments.length === 4) {
            logAndSendError500(arguments[0], arguments[1], arguments[2], arguments[3])
        } else if (arguments.length === 5) {
            logAndSendErrorCustom(arguments[0], arguments[1], arguments[2], arguments[3], arguments[4])
        }
    },*/
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