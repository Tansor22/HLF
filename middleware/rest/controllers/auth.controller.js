const config = require("../config/auth.config");
const db = require("../models");
const User = db.user;

const jwt = require("jsonwebtoken");
const bcrypt = require("bcryptjs");

module.exports = {
    signUp: (req, res) => {
        const user = new User({
            username: req.body.username,
            email: req.body.email,
            password: bcrypt.hashSync(req.body.password, 8)
        });
        user.save((err, user) => {
            if (err) {
                res.status(500).send({message: err});
                return;
            }
            res.send({result: "Ok"});
        });
    },
    signIn: (request, response, next) => {
        User.findOne({
            email: request.body.email
        }).exec((err, user) => {
            if (err) {
                next(err)
            } else if (!user) {
                return response.status(403).send({error: "UserNotFound"})
            } else {
                const passwordIsValid = bcrypt.compareSync(
                    request.body.password,
                    user.password
                )
                if (!passwordIsValid) {
                    return response.status(403).send({error: "ClientUnauthorized"})
                }
                const token = jwt.sign({id: user.id}, config.secret, {
                    expiresIn: 86400 // 24 hours
                });
                response.status(200).send({
                    username: user.username,
                    email: user.email,
                    accessToken: token
                });
            }
        })

    }
}