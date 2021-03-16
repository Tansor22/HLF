package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

const (
	// function names for dispatching
	NewDoc  = "new-doc"
	SignDoc = "sign-doc"
)

type IChaincodeFunction interface {
	BindParams([]string)
	Execute() peer.Response
}

type ChaincodeFunction struct {
	stub shim.ChaincodeStubInterface
	name string
}

func (f *ChaincodeFunction) BindParams([]string) {
	// default: no params
}

func (f *ChaincodeFunction) Execute() peer.Response {
	// default: error! unimplemented func
	return errorResponse("Invalid function: "+f.name, 1)
}

func NewChaincodeFunction(name string, stub shim.ChaincodeStubInterface) ChaincodeFunction {
	return ChaincodeFunction{name: name, stub: stub}
}

func NewFunction(name string, args []string, stub shim.ChaincodeStubInterface) IChaincodeFunction {
	defaultFunction := NewChaincodeFunction(name, stub)
	var output IChaincodeFunction
	switch name {
	case NewDoc:
		output = &CreateNewDocumentFunction{ChaincodeFunction: defaultFunction}
	default:
		output = &defaultFunction
	}
	output.BindParams(args)
	return output

}
