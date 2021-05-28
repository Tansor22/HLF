package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type ChangeDocumentFunction struct {
	ChaincodeFunction
	Id      string
	Member  string
	Type    string
	Details string
}
type SignDocumentResponse struct {
	IsSigned bool `json:"isSigned"`
}

func (f *ChangeDocumentFunction) BindParams(args []string) {
	f.Id = args[0]
	f.Member = args[1]
	f.Type = args[2]
	f.Details = args[3]
}

func (f *ChangeDocumentFunction) Execute() peer.Response {
	documentKey := "doc" + f.Id
	documentJson, _ := f.stub.GetState(documentKey)
	document, _ := DocumentFromJson(documentJson)
	// register change
	change := NewChange(f.Member, f.Type, f.Details)
	if err := document.RegisterChange(change); err != nil {
		return errorResponse(err.Error(), 2)
	}
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
