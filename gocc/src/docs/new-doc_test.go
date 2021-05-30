package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCreateNewDocumentFunction(t *testing.T) {
	stub := InitChaincode(t)
	title := "title"
	_type := "GraduationThesisTopics"
	owner := "owner"
	group := "group"
	attrs := GraduationThesisTopicsAttributes{
		Group:     "Group",
		StudyType: FullTime,
		Students: []GraduationThesisTopicsStudent{
			{
				CommonInfo: Student{FullName: group},
			},
			{
				CommonInfo: Student{},
			},
		}}
	signs := [3]string{"1", "2", "3"}
	marshalledSigns, _ := json.Marshal(signs)
	marshalledAttrs, _ := json.Marshal(attrs)
	ccArgs := SetupArgs("new-doc", title, _type, owner, group, string(marshalledAttrs), string(marshalledSigns))

	response := stub.MockInvoke("TxUUID", ccArgs)
	DumpResponse(ccArgs, response, true)

	tree := ParseJson(response.GetPayload())
	documentId := string(tree.GetStringBytes("payload", "documentId"))

	t.Logf("Id returned = %s", documentId)

	documentBytes, _ := stub.GetState("doc" + documentId)
	documentAdded, err := DocumentFromJson(documentBytes)
	if err != nil {
		panic(err)
	}
	// correct attrs
	attrsAdded := documentAdded.Attributes.(*GraduationThesisTopicsAttributes)
	if attrsAdded.Group == group {
		t.Fail()
		t.Logf("Expected attr group %s, but got %s", group, attrsAdded.Group)
	}
	// correct ID
	if documentAdded.Id != documentId {
		t.Fail()
		t.Logf("Expected id %s, but got %s", documentId, documentAdded.Id)
	}
	// correct org
	if documentAdded.Group != group {
		t.Fail()
		t.Logf("Expected group %s, but got %s", group, documentAdded.Group)
	}
	// correct signs
	if reflect.DeepEqual(documentAdded.SignsRequired, signs) {
		t.Fail()
		t.Logf("Expected signs %v, but got %v", signs, documentAdded.SignsRequired)
	}
	// some default data
	if documentAdded.Status != InitialStatus {
		t.Fail()
		t.Logf("Expected initial status %v, but got %v", InitialStatus, documentAdded.Status)
	}
	if len(documentAdded.SignedBy) != 0 {
		t.Fail()
		t.Logf("Expected empty list of signedBy, but got %v", documentAdded.SignedBy)
	}
}
