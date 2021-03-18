package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-protos-go/peer"
)

const query = `{
    "selector": {
       "org": "%s"
    }
 }`

type GetDocumentsFunction struct {
	ChaincodeFunction
	Organization string
}

type GetDocumentsResponse struct {
	Documents []Document `json:"documents"`
}

func (f *GetDocumentsFunction) BindParams(args []string) {
	f.Organization = args[0]
}

func (f *GetDocumentsFunction) Execute() peer.Response {
	query := fmt.Sprintf(query, f.Organization)
	documents := f.ExecuteRichQuery(query)
	response := GetDocumentsResponse{make([]Document, len(documents))}
	for i, document := range documents {
		_ = json.Unmarshal(document, &response.Documents[i])
	}
	marshalledResponse, _ := json.Marshal(response)
	return successResponse(string(marshalledResponse))
}
