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
        const docTypes = request.app.get('DOC_TYPES')
        const formConfig = docTypes[request.body.documentType || 'General']
        if (formConfig === undefined) {
            return response.logAndSendError('NoSuchForm', 'There is no form config for doc type ' + request.body.documentType + '.')
        }
        User.find({group: request.group, member: {$ne: request.member}}).exec((err, users) => {
            if (err) {
                return response.logAndSendError('UnexpectedError', 'Unexpected error, see middleware logs for details.')
            } else if (!(Array.isArray(users) && users.length)) {
                return response.logAndSendError('NoUsersInGroup', 'There are no any users in group ' + request.group + '.')
            } else {
                // substitute signs
                let docSignsConfigIndex = formConfig.findIndex(it => it._id === 'signs')
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
                body.title, body.type || "Unknown", body.owner, body.group, JSON.stringify(body.attributes),
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
            let docsResponse = JSON.parse(responseHLF)
            if (body.withContent === true) {
                const templates = request.app.get('TEMPLATES')
                for (let i = 0; i < docsResponse.payload.documents.length; i++) {
                    const parse = templates[docsResponse.payload.documents[i].type]
                    docsResponse.payload.documents[i].attributes.content = parse
                        ? parse(docsResponse.payload.documents[i].attributes) : null
                }
            }
            response.logAndSendOk(docsResponse)
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