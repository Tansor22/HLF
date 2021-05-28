package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-protos-go/peer"
)

const query = `{
    "selector": {
       "group": "%s"
    }
 }`

type GetDocumentsFunction struct {
	ChaincodeFunction
	Group string
}

type GetDocumentsResponse struct {
	Documents []Document `json:"documents"`
}

func (f *GetDocumentsFunction) BindParams(args []string) {
	f.Group = args[0]
}

func (f *GetDocumentsFunction) Execute() peer.Response {
	query := fmt.Sprintf(query, f.Group)
	documents := f.ExecuteRichQuery(query)
	response := GetDocumentsResponse{make([]Document, len(documents))}
	for i, document := range documents {
		response.Documents[i], _ = DocumentFromJson(document)
	}
	marshalledResponse, _ := json.Marshal(response)
	return successResponse(string(marshalledResponse))
}
