package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	"github.com/hyperledger/fabric-protos-go/peer"
)

const (
	// function names for dispatching
	NewDoc  = "new-doc"
	SignDoc = "sign-doc"
	GetDocs = "get-docs"
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

func NewFunction(name string, args []string, stub shim.ChaincodeStubInterface) IChaincodeFunction {
	defaultFunction := ChaincodeFunction{name: name, stub: stub}
	var output IChaincodeFunction
	switch name {
	case NewDoc:
		output = &CreateNewDocumentFunction{ChaincodeFunction: defaultFunction}
	case SignDoc:
		output = &SignDocumentFunction{ChaincodeFunction: defaultFunction}
	case GetDocs:
		output = &GetDocumentsFunction{ChaincodeFunction: defaultFunction}
	default:
		output = &defaultFunction
	}
	output.BindParams(args)
	return output

}
func (f *ChaincodeFunction) ExecuteRichQuery(query string) [][]byte {
	fmt.Printf("Query JSON=%s \n\n", query)
	iterator, err := f.stub.GetQueryResult(query)
	// return empty map in case of any errors
	if err != nil {
		fmt.Println("Error occurred during executing query =" + err.Error())
		return make([][]byte, 0)
	}
	var output [][]byte
	for iterator.HasNext() {
		var resultKV *queryresult.KV
		var err error
		resultKV, err = iterator.Next()
		if err != nil {
			fmt.Println("Error occurred during parsing value =" + err.Error())
			return make([][]byte, 0)
		}
		output = append(output, resultKV.GetValue())
	}
	_ = iterator.Close()
	return output
}
