const {setupGateway} = require("./tools");
module.exports = {
    authenticateClient: async function (request, response, next) {
        // брать юзер айди из заголовка, сетить реквестуу юзер ид
        request.userId = 'Admin@acme.com'
        next()
        //throw new Error('Client unauthorized')
    },
    connectToHLF: async function (request, response, next) {
        let profilePath = request.app.get('CONNECTION_PROFILE_PATH')
        let walletPath = request.app.get('FILESYSTEM_WALLET_PATH')
        request.gateway = await setupGateway(walletPath, profilePath, request.userId)
        next()
    },
    newDocument: async function (request, response) {
        let network = await request.gateway.getNetwork(request.app.get('NETWORK_NAME'))
        let contract = await network.getContract(request.app.get('CONTRACT_ID'));
        let {body} = request
        let responseHLF = await contract.submitTransaction('new-doc', body.org, body.content, JSON.stringify(body.signsRequired))
        response.json(responseHLF)
    },
    getDocuments: async function (request, response) {
        let network = await request.gateway.getNetwork(request.app.get('NETWORK_NAME'))
        let contract = await network.getContract(request.app.get('CONTRACT_ID'));
        let {body} = request
        let responseHLF = await contract.evaluateTransaction('get-docs', body.orgName)
        response.json(responseHLF)
    },
    signDocument: async function (request, response) {
        let network = await request.gateway.getNetwork(request.app.get('NETWORK_NAME'))
        let contract = await network.getContract(request.app.get('CONTRACT_ID'));
        let {body} = request
        let responseHLF = await contract.submitTransaction('sign-doc', body.documentId, JSON.stringify(body.signs))
        response.json(responseHLF)
    },
    // triggered by errors in async only in express 5, if use below should get wrapper with try { await willThrow() } catch (e) { next(e) }
    errorHandler: async function (e, request, response) {
        console.log("Error occurred: " + e)
        response.json({error: e})
    }
}