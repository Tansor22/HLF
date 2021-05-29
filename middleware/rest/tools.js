const fs = require('fs')
const yaml = require('js-yaml')
const {Gateway, FileSystemWallet} = require('fabric-network')
const path = require('path')
const ejs = require('ejs')

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
    },
    configureTemplates: function () {
        const templatesDir = path.resolve(__dirname, './config/templates')
        const templates = {}
        fs.readdir(templatesDir, (err, files) => {
            if (err) {
                return console.log('Unable to scan directory: ' + err);
            }
            files.forEach(file => {
                const postfixIndex = file.indexOf('.txt')
                if (file.startsWith('_doc_type_') && postfixIndex !== -1) {
                    const fileContent = fs.readFileSync(path.resolve(templatesDir, file), 'utf-8').toString()
                    templates[file.substring('_doc_type_'.length, postfixIndex)] = ejs.compile(fileContent, {rmWhitespace: true})
                }
            });
            console.log("Templates parsed: " + Object.keys(templates))
        })
        return templates
    },
    configureForms: function () {
        const formsDir = path.resolve(__dirname, './config/forms')
        const docTypes = {}
        fs.readdir(formsDir, (err, files) => {
            if (err) {
                return console.log('Unable to scan directory: ' + err);
            }
            files.forEach(file => {
                const postfixIndex = file.indexOf('.json')
                if (file.startsWith('_doc_type_') && postfixIndex !== -1) {
                    const fileContent = fs.readFileSync(path.resolve(formsDir, file), 'utf-8').toString()
                    docTypes[file.substring('_doc_type_'.length, postfixIndex)] = JSON.parse(fileContent)
                }
            });
            // substitute doc types
            let docTypesArr = Object.keys(docTypes)
            docTypesArr.forEach(type => {
                let docTypesConfigIndex = docTypes[type].findIndex(it => it._id === 'type')
                if (docTypesConfigIndex !== -1) {
                    docTypes[type][docTypesConfigIndex].list = []
                    for (let i = 0; i < docTypesArr.length; i++) {
                        docTypes[type][docTypesConfigIndex].list.push({
                            index: i, index_text: docTypesArr[i]
                        })
                    }
                }
            })
            console.log("Doc types parsed: " + Object.keys(docTypes))
        })
        return docTypes
    }
}