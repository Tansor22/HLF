const jwt = require("jsonwebtoken");
const config = require("../config/auth.config.js");

module.exports = {
    verifyToken: (req, res, next) => {
        let token = req.headers["x-access-token"];

        if (!token) {
            return res.logAndSendError(403, 'NoTokenProvided', 'Header \'x-access-token\' is not specified.')
        }

        jwt.verify(token, config.secret, (err, decoded) => {
            if (err) {
                return res.logAndSendError(401, 'UserUnauthorized', 'Invalid token received.')
            }
            req.userId = decoded.id
            req.group = decoded.group
            next();
        });
    }
}