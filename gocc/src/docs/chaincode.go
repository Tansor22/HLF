package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strings"
)

type DocsChaincode struct {
}

func (token *DocsChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init executed")
	// no init yet
	return successResponse("Initializing successful.")
}

func (token *DocsChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, ", args=", strings.Join(args, ","))

	return NewFunction(function, args, stub).Execute()
}
