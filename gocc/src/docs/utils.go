package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func errorResponse(err string, code uint) peer.Response {
	codeStr := strconv.FormatUint(uint64(code), 10)
	errorString := "{\"error\":" + err + ", \"code\":" + codeStr + " \" }"
	return shim.Error(errorString)
}

func successResponse(dat string) peer.Response {
	success := "{\"response\": " + dat + ", \"code\": 0 }"
	return shim.Success([]byte(success))
}
