package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type SignDocumentFunction struct {
	ChaincodeFunction
	Id    string
	Signs []string
}
type SignDocumentResponse struct {
	IsSigned bool `json:"isSigned"`
}

func (f *SignDocumentFunction) BindParams(args []string) {
	f.Id = args[0]
	_ = json.Unmarshal([]byte(args[1]), &f.Signs)
}

func (f *SignDocumentFunction) Execute() peer.Response {
	documentKey := "doc" + f.Id
	documentJson, _ := f.stub.GetState(documentKey)
	var document Document
	_ = json.Unmarshal(documentJson, &document)
	document.SignedBy = append(document.SignedBy, f.Signs...)
	// check for final status
	isSigned := document.IsSigned()
	// update document in blockchain
	marshalledDocument, _ := json.Marshal(document)
	_ = f.stub.PutState(documentKey, marshalledDocument)
	// return actual status
	response := SignDocumentResponse{isSigned}
	marshalledResponse, _ := json.Marshal(response)
	return successResponse(string(marshalledResponse))
}
