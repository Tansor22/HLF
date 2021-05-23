const config = require("../config/auth.config");
const db = require("../models");
const User = db.user;

const jwt = require("jsonwebtoken");
const bcrypt = require("bcryptjs");

module.exports = {
    signUp: (req, res) => {
        // todo add member, group
        const user = new User({
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
            email: request.body.email
        }).exec((err, user) => {
            if (err) {
                next(err)
            } else if (!user) {
                return response.logAndSendError(403, 'UserNotFound', 'A user with the email provided doesn\'t exist.')
            } else {
                const passwordIsValid = bcrypt.compareSync(
                    request.body.password,
                    user.password
                )
                if (!passwordIsValid) {
                    return response.logAndSendError(403, 'ClientUnauthorized', 'Password provided is not valid.')
                }
                const token = jwt.sign({id: user.id, member: user.member, email: user.email, group: user.group}, config.secret, {
                    expiresIn: 86400 // 24 hours
                });

                return response.logAndSendOk({
                    member: user.member,
                    email: user.email,
                    accessToken: token,
                    avatar: user.avatar
                })
            }
        })

    }
}