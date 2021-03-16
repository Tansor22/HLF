package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type CreateNewDocumentFunction struct {
	ChaincodeFunction
	Organization string
	Content      string
	Signs        []string
}

type CreateNewDocumentResponse struct {
	DocumentId string `json:"documentId"`
}

func (f *CreateNewDocumentFunction) BindParams(args []string) {
	f.Organization = args[0]
	f.Content = args[1]
	_ = json.Unmarshal([]byte(args[2]), &f.Signs)
}

func (f *CreateNewDocumentFunction) Execute() peer.Response {
	document := NewDocument(f.Organization, f.Content, f.Signs)
	marshalledDocument, _ := json.Marshal(document)
	_ = f.stub.PutState("doc"+document.Id, marshalledDocument)
	response := CreateNewDocumentResponse{
		DocumentId: document.Id,
	}
	marshalledResponse, _ := json.Marshal(response)
	return successResponse(string(marshalledResponse))
}
