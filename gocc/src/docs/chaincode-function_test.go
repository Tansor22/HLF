package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"testing"
)

func PopulateBlockchain(stub *shimtest.MockStub) {
	for i := 0; i < 10; i++ {
		title := "title" + string(rune(i))
		_type := "type" + string(rune(i))
		owner := "owner" + string(rune(i))
		group := "group" + string(rune(i))
		content := "Some attrsJson" + string(rune(i))
		signs := [3]string{"1", "2", "3"}
		marshalledSigns, _ := json.Marshal(signs)
		ccArgs := SetupArgs("new-doc", title, _type, owner, group, content, string(marshalledSigns))
		_ = stub.MockInvoke("TxUUID", ccArgs)
	}
}
func TestExecuteRichQuery(t *testing.T) {
	stub := InitChaincode(t)
	PopulateBlockchain(stub)
	_ = ChaincodeFunction{name: "test-func", stub: stub}
	_ = `{
    "selector": {
       "org": "orgName"
    }
 }
 `
	// cannot be tested, not implemented in shimtest
	//result := function.ExecuteRichQuery(query)
	//t.Logf("Result %v", result)
}
