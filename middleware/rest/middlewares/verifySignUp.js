const db = require("../models");
const User = db.user;

module.exports = {
    checkDuplicateUsernameOrEmail: (req, res, next) => {
        // Username
        User.findOne({
            username: req.body.username
        }).exec((err, user) => {
            if (err) {
                res.logAndSendError(err, 'Mongo exception.')
                return
            }

            if (user) {
                res.logAndSendError(400, 'UsernameAlreadyInUse', 'There is an existing user with login specified.')
                return
            }

            // Email
            User.findOne({
                email: req.body.email
            }).exec((err, user) => {
                if (err) {
                    res.logAndSendError(err, 'Mongo exception.')
                    return;
                }

                if (user) {
                    res.logAndSendError(400, 'EmailAlreadyInUse', 'There is an existing user with email specified.')
                    return
                }

                next();
            });
        });
    }
}