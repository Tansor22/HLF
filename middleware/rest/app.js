const express = require('express');
const https = require('https');
const bodyParser = require('body-parser');
const middlewares = require("./middlewares");
const fs = require("fs");
const app = express();

// Configure app
app.set('CONNECTION_PROFILE_PATH', '../profiles/dev-connection.yaml')
app.set('NETWORK_NAME', 'airlinechannel')
app.set('CONTRACT_ID', 'docs')
app.set('FILESYSTEM_WALLET_PATH', '../gateway/user-wallet')
app.use(bodyParser.json())

// Interceptors
app.use(middlewares.authenticateClient)
app.use(middlewares.connectToHLF)

// Routes
app.post('/newDoc', middlewares.newDocument)
app.post('/getDocs', middlewares.getDocuments)
app.post('/signDoc', middlewares.signDocument)
/*todo delete*/
app.post('/test',  async function (request, response) {
    response.end(JSON.stringify({msg: 'Hello'}))
})

// Error handlers
app.use(middlewares.errorHandler)

// Configure server
const key = fs.readFileSync(__dirname + '/certs/server-key.pem');
const cert = fs.readFileSync(__dirname + '/certs/server-cert.pem');
const host = 'hlf-gtw.local'
const port = 443
const options = {
    key,
    cert,
    host,
    port,
    requestCert: true,
    rejectUnauthorized: false,
    ca: [ fs.readFileSync(__dirname + '/certs/client-cert.pem') ]
}
const server = https.createServer(options, app)

// Launch server
server.listen(port, host,() => {
    console.log("REST server listening at https://%s:%s", host, port)
})