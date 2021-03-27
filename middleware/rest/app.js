const express = require('express');
const bodyParser = require('body-parser');
const middlewares = require("./middlewares");
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

// Error handlers
app.use(middlewares.errorHandler)

// Launch server
app.listen(8081, function () {
    const host = this.address().address;
    const port = this.address().port;
    console.log("REST server listening at http://%s:%s", host, port)
})