function parseHLFError(e) {
    const msg = e.message
    let indexOfError = msg.indexOf('Error: ');
    // from 'Error: ' to '.' (dot)
    return msg.substring(indexOfError, msg.indexOf('.'))
}

module.exports = {
    newDocument: async function (request, response) {
        try {
            let gateway = await Promise.resolve(request.gateway)
            let network = await gateway.getNetwork(request.app.get('NETWORK_NAME'))
            let contract = await network.getContract(request.app.get('CONTRACT_ID'));
            let {body} = request
            let responseHLF = await contract.submitTransaction('new-doc',
                body.title, body.type, body.owner, body.group, body.content,
                JSON.stringify(body.signsRequired))
            response.logAndSendOk(JSON.parse(responseHLF))
        } catch (e) {
            return response.logAndSendError("HLFError", e.message)
        }
    },
    getDocuments: async function (request, response) {
      try {
            let gateway = await Promise.resolve(request.gateway)
            let network = await gateway.getNetwork(request.app.get('NETWORK_NAME'))
            let contract = await network.getContract(request.app.get('CONTRACT_ID'));
            let {body} = request
            let responseHLF = await contract.evaluateTransaction('get-docs', body.group)
            response.logAndSendOk(JSON.parse(responseHLF))
        } catch (e) {
            return response.logAndSendError("HLFError", parseHLFError(e.message))
        }

    },
    changeDocument: async function (request, response) {
        try {
            let gateway = await Promise.resolve(request.gateway)
            let network = await gateway.getNetwork(request.app.get('NETWORK_NAME'))
            let contract = await network.getContract(request.app.get('CONTRACT_ID'));
            let {body} = request
            let responseHLF = await contract.submitTransaction('change-doc',
                body.documentId, body.member, body.type, body.details)
            response.send(responseHLF)
        } catch (e) {
            return response.logAndSendError("HLFError", e.message)
        }
    },
}