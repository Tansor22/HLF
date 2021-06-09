package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
	"testing"
)

func TestSignDocumentFunction(t *testing.T) {
	stub := InitChaincode(t)
	// adding doc in blockchain
	signs := []string{"1", "2", "3"}
	documentId := AddDocumentInBlockchain(stub, signs)
	// calling change-doc func
	var response peer.Response
	for i := 1; i <= 3; i++ {
		ccArgs := SetupArgs("change-doc", documentId, strconv.Itoa(i), "APPROVE", "")
		response = stub.MockInvoke("TxUUID2", ccArgs)
		DumpResponse(ccArgs, response, true)
	}

	tree := ParseJson(response.GetPayload())
	if tree.GetBool("payload", "isSigned") == false {
		t.Fail()
		t.Log("Expected response contains isSigned = true")
	}
	// checking document state in blockchain
	documentBytes, _ := stub.GetState("doc" + documentId)
	var document Document
	var err error
	if document, err = DocumentFromJson(documentBytes); err != nil {
		panic(err)
	}
	if document.Status != Approved {
		t.Fail()
		t.Logf("Expected document has status '%s', but it's actual status is %s", Approved, document.Status)
	}
}

func TestChangeDocumentFunctionEdit(t *testing.T) {
	stub := InitChaincode(t)
	// adding doc in blockchain
	signs := []string{"1", "2", "3"}
	documentId := AddDocumentInBlockchain(stub, signs)
	oldContent := "Some attrsJson"
	newContent := "New Content"

	oldDocumentBytes, _ := stub.GetState("doc" + documentId)
	var oldDocument Document
	var e error
	if oldDocument, e = DocumentFromJson(oldDocumentBytes); e != nil {
		panic(e)
	}
	if oldDocument.Attributes.(*DocAttributes).Content != oldContent {
		t.Fail()
		t.Logf("Expected document has  '%s', but it's actual status is %s", "Some attrsJson", oldDocument.Attributes.(*DocAttributes).Content)
	}
	// calling change-doc func
	var response peer.Response
	ccArgs := SetupArgs("change-doc", documentId, "1", "EDIT", "Changed by some member.", "{\"content\": \""+newContent+"\"}")
	response = stub.MockInvoke("TxUUID2", ccArgs)
	DumpResponse(ccArgs, response, true)

	documentBytes, _ := stub.GetState("doc" + documentId)
	var document Document
	var err error
	if document, err = DocumentFromJson(documentBytes); err != nil {
		panic(err)
	}
	if document.Attributes.(*DocAttributes).Content != newContent {
		t.Fail()
		t.Logf("Expected document has content '%s', but it's actual content is %s", newContent, document.Attributes.(*DocAttributes).Content)
	}
	// todo attrs of changes not parsed
	/*	change:= document.Changes[len(document.Changes) - 1]
		if (change.Attributes).(*DocAttributes).Content != oldContent {
			t.Fail()
			t.Logf("Expected document's last cahge has content '%s', but it's actual content is %s", oldContent, (change.Attributes).(*DocAttributes).Content)
		}*/
}

func TestSignDocumentFunctionErrorsCheck(t *testing.T) {
	stub := InitChaincode(t)
	// adding doc in blockchain
	signs := []string{"1", "2", "3"}
	documentId := AddDocumentInBlockchain(stub, signs)
	// calling change-doc func
	var response peer.Response
	var ccArgs [][]byte
	ccArgs = SetupArgs("change-doc", documentId, "1", "APPROVE", "")
	response = stub.MockInvoke("TxUUID2", ccArgs)
	DumpResponse(ccArgs, response, true)

	// should be 2
	ccArgs = SetupArgs("change-doc", documentId, "3", "APPROVE", "")
	response = stub.MockInvoke("TxUUID2", ccArgs)
	DumpResponse(ccArgs, response, true)

	if response.GetMessage() == "" {
		t.Fail()
		t.Log("Expected error")
	}
	// checking document state in blockchain
	documentBytes, _ := stub.GetState("doc" + documentId)
	var document Document
	var err error
	if document, err = DocumentFromJson(documentBytes); err != nil {
		panic(err)
	}
	if sign, _ := document.getCurrentSign(); sign != "2" {
		t.Fail()
		t.Logf("Expected current sign to be '2', but got %s", sign)
	}
	if document.Changes[0].Member != "1" || document.Changes[0].Type != "APPROVE" {
		t.Fail()
		t.Logf("Expected document has changed")
	}
	// bad sign
	ccArgs = SetupArgs("change-doc", documentId, "32", "REJECT", "")
	response = stub.MockInvoke("TxUUID2", ccArgs)
	DumpResponse(ccArgs, response, true)
	if response.GetMessage() == "" {
		t.Fail()
		t.Log("Expected error")
	}
}

func AddDocumentInBlockchain(stub *shimtest.MockStub, signs []string) string {
	title := "title"
	_type := "General"
	owner := "owner"
	group := "group"
	content := "Some attrsJson"
	marshalledSigns, _ := json.Marshal(signs)
	ccArgs := SetupArgs("new-doc", title, _type, owner, group, "{\"content\": \""+content+"\"}", string(marshalledSigns))

	response := stub.MockInvoke("TxUUID", ccArgs)

	tree := ParseJson(response.GetPayload())
	return string(tree.GetStringBytes("payload", "documentId"))
}
