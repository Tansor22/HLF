const richConsole = require('rich-console')
const stringify = require('json-stringify-safe');
module.exports = {
    configure: async function (request, response, next) {
        response.logAndSend = (status, data) => {
            // log
            const resp = data.payload ? data : {payload: data, code: status === 200 ? 'Ok' : 'Nok'}
            let msg = `\tEndpoint: ${request.path}\n\tRequest: ${stringify(request.body)}\n\tResponse (Code:${status}): ${stringify(resp)}\n`
            let richMsg = status === 200 ? '<green>' + msg + '</green>>' : '<red>' + msg + '</red>>'
            richConsole.log(richMsg)
            // send
            response.status(status).send(resp)
        };
        response.logAndSendOk = (data) => {
            response.logAndSend(200, data ? data : {result: 'Ok'})
        }
        const logAndSendErrorCustom = (status, error, details) => {
            response.logAndSend(status, {error, details})
        }
        const logAndSendError500 = (error, details) => {
            response.logAndSend(500, {error, details})
        }
        response.logAndSendError = (... args) => {
            if (args.length === 2) {
                logAndSendError500(args[0], args[1])
            } else if (arguments.length === 3) {
                logAndSendErrorCustom(args[0], args[1], args[2])
            }
        }
        next()
    },
    verifyCert: async function (request, response, next) {
        if (request.socket.authorized) {
            next()
        } else {
            // todo fix
            next()
            // certificate incorrect
           // return response.logAndSendError(403, 'ClientUnauthorized', 'Client app certificate is not valid.')
        }
    },
    // triggered by errors in async only in express 5, if use below should get wrapper with try { await willThrow() } catch (e) { next(e) }
    errorHandler: async function (e, request, response) {
        console.log("Error occurred: " + e.stack)
        response.logAndSendError('UnexpectedError', 'See server logs for details.')
    }
}