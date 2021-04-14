module.exports = {
    verifyCert: async function (request, response, next) {
        if (request.socket.authorized) {
            next()
        } else {
            // certificate incorrect
            return response.status(403).send({error: "ClientUnauthorized"});
        }
    },
    // triggered by errors in async only in express 5, if use below should get wrapper with try { await willThrow() } catch (e) { next(e) }
    errorHandler: async function (e, request, response) {
        console.log("Error occurred: " + e.stack)
        response.status(500).send({error: e})
    }
}