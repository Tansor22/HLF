package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type ChangeDocumentFunction struct {
	ChaincodeFunction
	Id             string
	Member         string
	Type           string
	Details        string
	AttributesJson string
}
type SignDocumentResponse struct {
	IsSigned bool `json:"isSigned"`
}

func (f *ChangeDocumentFunction) BindParams(args []string) {
	f.Id = args[0]
	f.Member = args[1]
	f.Type = args[2]
	f.Details = args[3]
	if len(args) > 4 {
		f.AttributesJson = args[4]
	} else {
		f.AttributesJson = ""
	}
}

func (f *ChangeDocumentFunction) Execute() peer.Response {
	documentKey := "doc" + f.Id
	documentJson, _ := f.stub.GetState(documentKey)
	document, _ := DocumentFromJson(documentJson)
	var attrs IDocAttributes
	if f.AttributesJson != "" {
		var e error
		if attrs, e = AttributesFromJson(document.Type, f.AttributesJson); e != nil {
			return errorResponse(e.Error(), 2)
		}
	} else {
		attrs = nil
	}
	// register change
	change := NewChange(f.Member, f.Type, f.Details, attrs)
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
