const db = require("../models");
const User = db.user;

module.exports = {
    checkDuplicateEmail: (req, res, next) => {
        // Email
        User.findOne({
            email: req.body.email
        }).exec((err, user) => {
            if (err) {
                return res.logAndSendError(err, 'Mongo exception.')
            }

            if (user) {
                return res.logAndSendError(400, 'EmailAlreadyInUse', 'There is an existing user with email specified.')
            }
            next();
        });
    }
}