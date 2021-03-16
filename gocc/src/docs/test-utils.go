package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/valyala/fastjson"
	"strings"
	"testing"
)

func InitChaincode(t *testing.T) *shimtest.MockStub {
	stub := shimtest.NewMockStub("TestStub", new(DocsChaincode))
	response := stub.MockInit("mockTxId", nil)
	status := response.GetStatus()
	t.Logf("Received status = %d", status)
	if response.GetStatus() != shim.OK {
		t.FailNow()
	}
	return stub
}

func SetupArgs(funcName string, args ...string) [][]byte {
	// Create an args array with 1 additional element for the funcName
	ccArgs := make([][]byte, 1+len(args))

	// Setup the function name
	ccArgs[0] = []byte(funcName)

	// Set up the args array
	for i, arg := range args {
		ccArgs[i+1] = []byte(arg)
	}

	return ccArgs
}
func ParseJson(jsonBytes []byte) *fastjson.Value {
	var p fastjson.Parser
	tree, err := p.ParseBytes(jsonBytes)
	if err != nil {
		panic(err)
	}
	return tree
}
func DumpResponse(args [][]byte, response peer.Response, printFlag bool) {
	if !printFlag {
		return
	}
	argsArray := make([]string, len(args))
	for i, arg := range args {
		argsArray[i] = string(arg)
	}
	fmt.Println("Call:    ", strings.Join(argsArray, ","))
	fmt.Println("RetCode: ", response.Status)
	fmt.Println("RetMsg:  ", response.Message)
	fmt.Println("Payload: ", string(response.Payload))
}
