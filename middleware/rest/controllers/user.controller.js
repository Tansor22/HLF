module.exports = {
    newDocument: async function (request, response) {
        // работает ка джойн
        let gateway = await Promise.resolve(request.gateway)
        let network = await gateway.getNetwork(request.app.get('NETWORK_NAME'))
        let contract = await network.getContract(request.app.get('CONTRACT_ID'));
        let {body} = request
        let responseHLF = await contract.submitTransaction('new-doc', body.org, body.content, JSON.stringify(body.signsRequired))
        response.send(responseHLF)
    },
    getDocuments: async function (request, response) {
        // работает ка джойн
        let gateway = await Promise.resolve(request.gateway)
        let network = await gateway.getNetwork(request.app.get('NETWORK_NAME'))
        let contract = await network.getContract(request.app.get('CONTRACT_ID'));
        let {body} = request
        let responseHLF = await contract.evaluateTransaction('get-docs', body.orgName)
        response.send(responseHLF)
    },
    signDocument: async function (request, response) {
        // работает ка джойн
        let gateway = await Promise.resolve(request.gateway)
        let network = await gateway.getNetwork(request.app.get('NETWORK_NAME'))
        let contract = await network.getContract(request.app.get('CONTRACT_ID'));
        let {body} = request
        let responseHLF = await contract.submitTransaction('sign-doc', body.documentId, JSON.stringify(body.signs))
        response.send(responseHLF)
    },
}