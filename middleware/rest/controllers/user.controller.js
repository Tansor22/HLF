const fs = require("fs")
const path = require('path')

const db = require("../models")
const User = db.user

function parseHLFError(e) {
    const msg = e.message
    let indexOfError = msg.indexOf('Error: ');
    // from 'Error: ' to '.' (dot)
    return msg.substring(indexOfError, msg.indexOf('.'))
}

module.exports = {
    getFormConfig: async function (request, response) {
        const fileContent = fs.readFileSync(path.resolve(__dirname, '../config/forms/general.json')).toString()
        const formConfig = JSON.parse(fileContent)
        User.find({group: request.group, member: {$ne: request.member}}).exec((err, users) => {
            if (err) {
                return response.logAndSendError('UnexpectedError', 'Unexpected error, see middleware logs for details.')
            } else if (!(Array.isArray(users) && users.length)) {
                return response.logAndSendError('NoUsersInGroup', 'There are no any users in group ' + request.group)
            } else {
                // todo handle doc type
                let docSignsConfigIndex = formConfig.findIndex(it => it._id === 'doc_signs_multi_spinner')
                if (docSignsConfigIndex !== -1) {
                    formConfig[docSignsConfigIndex].list = []
                    for (let i = 0; i < users.length; i++) {
                        formConfig[docSignsConfigIndex].list.push({
                            index: i, index_text: users[i].member
                        })
                    }
                }
                return response.logAndSendOk({config: formConfig})
            }
        })
    },
    newDocument: async function (request, response) {
        try {
            let gateway = await Promise.resolve(request.gateway)
            let network = await gateway.getNetwork(request.app.get('NETWORK_NAME'))
            let contract = await network.getContract(request.app.get('CONTRACT_ID'));
            let {body} = request
            let responseHLF = await contract.submitTransaction('new-doc',
                body.title, body.type || "Unknown", body.owner, body.group, body.content,
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
                body.documentId, body.member, body.type, body.details || "")
            response.logAndSendOk(JSON.parse(responseHLF))
        } catch (e) {
            return response.logAndSendError("HLFError", e.message)
        }
    },
}