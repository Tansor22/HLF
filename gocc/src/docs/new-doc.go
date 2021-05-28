package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type CreateNewDocumentFunction struct {
	ChaincodeFunction
	Title     string
	Type      string
	Owner     string
	Group     string
	attrsJson string
	Signs     []string
}

type CreateNewDocumentResponse struct {
	DocumentId string `json:"documentId"`
}

func (f *CreateNewDocumentFunction) BindParams(args []string) {
	f.Title = args[0]
	f.Type = args[1]
	f.Owner = args[2]
	f.Group = args[3]
	f.attrsJson = args[4]
	_ = json.Unmarshal([]byte(args[5]), &f.Signs)
}

func (f *CreateNewDocumentFunction) Execute() peer.Response {
	document, e := NewDocument(f.Title, f.Type, f.Owner, f.Group, f.attrsJson, f.Signs)
	if e != nil {
		return errorResponse(e.Error(), 3)
	}
	marshalledDocument, _ := json.Marshal(document)
	_ = f.stub.PutState("doc"+document.Id, marshalledDocument)
	response := CreateNewDocumentResponse{
		DocumentId: document.Id,
	}
	marshalledResponse, _ := json.Marshal(response)
	return successResponse(string(marshalledResponse))
}
