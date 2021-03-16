package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"testing"
)

func TestSignDocumentFunction(t *testing.T) {
	stub := InitChaincode(t)
	// adding doc in blockchain
	signs := []string{"1", "2", "3"}
	documentId := AddDocumentInBlockchain(stub, signs)
	// calling sign-doc func
	marshalledSigns, _ := json.Marshal(signs)
	ccArgs := SetupArgs("sign-doc", documentId, string(marshalledSigns))
	response := stub.MockInvoke("TxUUID2", ccArgs)
	DumpResponse(ccArgs, response, true)

	tree := ParseJson(response.GetPayload())
	if tree.GetBool("response", "isSigned") == false {
		t.Fail()
		t.Log("Expected response contains isSigned = true")
	}
	// checking document state in blockchain
	documentBytes, _ := stub.GetState("doc" + documentId)
	var document Document
	if err := json.Unmarshal(documentBytes, &document); err != nil {
		panic(err)
	}
	if document.Status != Approved {
		t.Fail()
		t.Logf("Expected document has status '%s', but it's actual status is %s", Approved, document.Status)
	}
}

func AddDocumentInBlockchain(stub *shimtest.MockStub, signs []string) string {
	orgName := "orgName"
	content := "Some Content"
	marshalledSigns, _ := json.Marshal(signs)
	ccArgs := SetupArgs("new-doc", orgName, content, string(marshalledSigns))

	response := stub.MockInvoke("TxUUID", ccArgs)

	tree := ParseJson(response.GetPayload())
	return string(tree.GetStringBytes("response", "documentId"))
}
