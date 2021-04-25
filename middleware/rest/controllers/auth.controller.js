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
                res.logAndSendError(err, 'Error saving user data.')
                return
            }
            res.logAndSendOk()
        });
    },
    signIn: (request, response, next) => {
        User.findOne({
            username: request.body.username
        }).exec((err, user) => {
            if (err) {
                next(err)
            } else if (!user) {
                return response.logAndSendError(403, 'UserNotFound', 'A user with the username provided doesn\'t exist.')
            } else {
                const passwordIsValid = bcrypt.compareSync(
                    request.body.password,
                    user.password
                )
                if (!passwordIsValid) {
                    return response.logAndSendError(403, 'ClientUnauthorized', 'Password provided is not valid.')
                }
                const token = jwt.sign({id: user.id}, config.secret, {
                    expiresIn: 86400 // 24 hours
                });

                return response.logAndSendOk({
                    username: user.username,
                    email: user.email,
                    accessToken: token
                })
            }
        })

    }
}