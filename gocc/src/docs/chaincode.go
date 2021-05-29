package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strings"
)

type DocsChaincode struct {
}

func (token *DocsChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init executed")
	// add sample docs
	document, e := NewDocument(
		"TypeGraduationThesisTopics",
		TypeGraduationThesisTopics,
		"Л.И. Сучкова",
		"Administration",
		"{}",
		[]string{"С.М. Старолетов", "С.А. Кантор"})
	if e != nil {
		return errorResponse(e.Error(), 3)
	}
	t := true
	f := false
	document.Attributes = &GraduationThesisTopicsAttributes{
		Group:      "Group",
		StudyType:  FullTime,
		Speciality: "Speciality",
		Students: []GraduationThesisTopicsStudent{
			{
				CommonInfo:              Student{FullName: "Student name1", Nationality: "Nationality1", Group: "Group1", OnGovernmentPay: &t},
				Topic:                   "Topic1",
				AcademicAdvisorFullName: "AcademicAdvisorFullName1",
			},
			{
				CommonInfo:              Student{FullName: "Student name2", Nationality: "Nationality2", Group: "Group2", OnGovernmentPay: &f},
				Topic:                   "Topic2",
				AcademicAdvisorFullName: "AcademicAdvisorFullName2",
			},
		}}
	marshalledDocument, _ := json.Marshal(document)
	_ = stub.PutState("doc"+document.Id, marshalledDocument)
	return successResponse("Initializing successful.")
}

func (token *DocsChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, ", args=", strings.Join(args, ","))

	return NewFunction(function, args, stub).Execute()
}
