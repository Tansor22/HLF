const jwt = require("jsonwebtoken");
const config = require("../config/auth.config.js");

module.exports = {
    verifyToken: (req, res, next) => {
        let token = req.headers["x-access-token"];

        if (!token) {
            return res.status(403).send({error: "NoTokenProvided"});
        }

        jwt.verify(token, config.secret, (err, decoded) => {
            if (err) {
                return res.status(401).send({error: "UserUnauthorized"});
            }
            req.userId = decoded.id
            next();
        });
    }
}