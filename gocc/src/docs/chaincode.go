package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"time"
)

const (
	// function names for dispatching
	Dummy = "dummy-func"
	// todo add newDoc func
	NewDoc = "new-doc"
)

type DocsChaincode struct {
}

func (token *DocsChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init executed")
	// creating dummy document
	dummyDoc := Document{
		DocumentId:     "1",
		Date:           time.Now(),
		Content:        "Some content",
		DocumentStatus: WaitingForApproval,
	}
	jsonDummyDoc, _ := json.Marshal(dummyDoc)
	_ = stub.PutState("dummyDoc", jsonDummyDoc)
	return successResponse("Dummy doc added: " + string(jsonDummyDoc))
}

func (token *DocsChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, " args=", args)

	switch function {
	case Dummy:
		return successResponse("Dummy function called successfully")
	case NewDoc:
		return successResponse("Dummy function called successfully")
	default:
		return errorResponse("Invalid function: "+function, 1)
	}
}
