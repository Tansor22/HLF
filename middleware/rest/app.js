const express = require('express');
const https = require('https');
const bodyParser = require('body-parser');
const middlewares = require("./middlewares");
const path = require('path')
const fs = require("fs");
const app = express();

// Configure app
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
        let docTypesConfigIndex = docTypes[type].findIndex(it => it._id === 'doc_type_spinner')
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
app.set('DOC_TYPES', docTypes)
app.set('CONNECTION_PROFILE_PATH', '../profiles/dev-connection.yaml')
app.set('NETWORK_NAME', 'airlinechannel')
app.set('CONTRACT_ID', 'docs')
app.set('FILESYSTEM_WALLET_PATH', '../gateway/user-wallet')
app.use(bodyParser.json())

// Interceptors
// todo should be configuring, but actual middleware
app.use(middlewares.common.configure)
app.use(middlewares.common.verifyCert)

// Routes
require('./routes/auth.routes')(app);
require('./routes/user.routes')(app);

// Error handlers ???
app.use(middlewares.common.errorHandler)

// Configure server
const key = fs.readFileSync(__dirname + '/certs/server-key.pem');
const cert = fs.readFileSync(__dirname + '/certs/server-cert.pem');
//const host = 'hlf-gtw.local'
const host = 'localhost'
const port = 443
const options = {
    key,
    cert,
    host,
    port,
    requestCert: true,
    rejectUnauthorized: false,
    ca: [fs.readFileSync(__dirname + '/certs/client-cert.pem')]
}
const server = https.createServer(options, app)

// Launch server
server.listen(port, host, () => {
    console.log("REST server listening at https://%s:%s", host, port)
})

// Connect to db
const db = require("./models");

db.mongoose
    .connect(`mongodb://${db.config.HOST}:${db.config.PORT}/${db.config.DB}`, {
        useNewUrlParser: true,
        useUnifiedTopology: true
    })
    .then(() => {
        console.log("Successfully connect to MongoDB.");
        // todo delete add test user
        db.user.estimatedDocumentCount((e, count) => {
            if (!e && count === 0) {
                new db.user(
                    {username: 'admin', email: 'admin@com', password: 'admin'}
                ).save(e => {
                    if (e) {
                        console.log("error", e);
                        process.exit()
                    } else {
                        console.log("added 'user' to roles collection");
                    }
                })
            }
        })
    })
    .catch(err => {
        console.error("Connection error", err);
        process.exit();
    });