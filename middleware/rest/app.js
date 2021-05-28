const express = require('express');
const https = require('https');
const bodyParser = require('body-parser');
const middlewares = require("./middlewares");
const fs = require("fs");
const app = express();
const tools = require('./tools')


// Configure app
app.set('TEMPLATES', tools.configureTemplates())
app.set('DOC_TYPES', tools.configureForms())
app.set('CONNECTION_PROFILE_PATH', '../profiles/dev-connection.yaml')
app.set('NETWORK_NAME', 'airlinechannel')
app.set('CONTRACT_ID', 'docs')
app.set('FILESYSTEM_WALLET_PATH', '../gateway/user-wallet')
app.use(bodyParser.json())

// Interceptors
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