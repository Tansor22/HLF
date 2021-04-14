const db = require("../models");
const User = db.user;

module.exports = {
    checkDuplicateUsernameOrEmail: (req, res, next) => {
        // Username
        User.findOne({
            username: req.body.username
        }).exec((err, user) => {
            if (err) {
                res.status(500).send({message: err});
                return;
            }

            if (user) {
                res.status(400).send({error: "UsernameAlreadyInUse"});
                return;
            }

            // Email
            User.findOne({
                email: req.body.email
            }).exec((err, user) => {
                if (err) {
                    res.status(500).send({message: err});
                    return;
                }

                if (user) {
                    res.status(400).send({message: "EmailAlreadyInUse"});
                    return;
                }

                next();
            });
        });
    }
}