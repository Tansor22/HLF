const {chaincode, authJwt} = require("../middlewares")
const userController = require("../controllers/user.controller")

module.exports = function (app) {
    app.post('/api/service/getFormConfig',
        [
            authJwt.verifyToken
        ],
        userController.getFormConfig)

    app.post('/api/chaincode/newDoc',
        [
            authJwt.verifyToken,
            chaincode.connectToHLF
        ],
        userController.newDocument)

    app.post('/api/chaincode/getDocs',
        [
            authJwt.verifyToken,
            chaincode.connectToHLF
        ],
        userController.getDocuments)

    app.post('/api/chaincode/changeDoc',
        [
            authJwt.verifyToken,
            chaincode.connectToHLF
        ],
        userController.changeDocument)
}
