const {setupGateway} = require("../tools");
const db = require("../models");
const User = db.user;
module.exports = {
    connectToHLF: function (request, response, next) {
        User.findById(request.userId).exec((err, user) => {
            if (err) {
                response.logAndSendError(err, 'User not found.')
                return
            }
            let profilePath = request.app.get('CONNECTION_PROFILE_PATH')
            let walletPath = request.app.get('FILESYSTEM_WALLET_PATH')
            /* todo connection depends on user, already here*/
            console.log("User Identity : " + user.username + user.email)
            request.gateway = setupGateway(walletPath, profilePath, /*request.userId*/'Admin@acme.com')
            next()
        })
    }
}