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

func addGraduationThesisTopicsDocInBlockchain(stub shim.ChaincodeStubInterface) error {
	document, e := NewDocument(
		"TypeGraduationThesisTopics",
		TypeGraduationThesisTopics,
		"Л.И. Сучкова",
		"Administration",
		"{}",
		[]string{"С.М. Старолетов", "С.А. Кантор"})
	if e != nil {
		return e
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
	return nil
}
func addGraduatedExpellingDocInBlockchain(stub shim.ChaincodeStubInterface) error {
	document, e := NewDocument(
		"TypeGraduatedExpelling",
		TypeGraduatedExpelling,
		"Л.И. Сучкова",
		"Administration",
		"{}",
		[]string{"С.М. Старолетов", "С.А. Кантор"})
	if e != nil {
		return e
	}

	t := true
	f := false
	course := 1
	document.Attributes = &GraduatedExpellingAttributes{
		Qualification: "Qualification",
		Course:        &course,
		Faculty:       "Faculty",
		Speciality:    "Speciality",
		Students: []GraduatedExpellingStudent{
			{
				CommonInfo:       Student{FullName: "Student name1", Nationality: "Nationality1", Group: "Group1", OnGovernmentPay: &t},
				HasHonoursDegree: &t,
				ExamDate:         "20.06.21",
			},
			{
				CommonInfo:       Student{FullName: "Student name2", Nationality: "Nationality2", Group: "Group2", OnGovernmentPay: &f},
				HasHonoursDegree: &f,
				ExamDate:         "23.06.21",
			},
		}}
	marshalledDocument, _ := json.Marshal(document)
	_ = stub.PutState("doc"+document.Id, marshalledDocument)
	return nil
}
func addPracticePermissionDocInBlockchain(stub shim.ChaincodeStubInterface) error {
	document, e := NewDocument(
		"TypePracticePermission",
		TypePracticePermission,
		"Л.И. Сучкова",
		"Administration",
		"{}",
		[]string{"С.М. Старолетов", "С.А. Кантор"})
	if e != nil {
		return e
	}

	t := true
	f := false
	course := 1
	document.Attributes = &PracticePermissionAttributes{
		PracticeType: "Учебная",
		Course:       &course,
		Speciality:   "Speciality",
		StudyType:    "FULL_TIME",
		DateFrom:     "20.05.21",
		DateTo:       "20.06.21",
		Students: []PracticePermissionStudent{
			{
				CommonInfo:       Student{FullName: "Student name1", Nationality: "Nationality1", Group: "Group1", OnGovernmentPay: &t},
				PracticeLocation: "PracticeLocation",
				HeadFullName:     "HeadFullName",
			},
			{
				CommonInfo:       Student{FullName: "Student name2", Nationality: "Nationality2", Group: "Group2", OnGovernmentPay: &f},
				PracticeLocation: "PracticeLocation",
				HeadFullName:     "HeadFullName",
			},
		}}
	marshalledDocument, _ := json.Marshal(document)
	_ = stub.PutState("doc"+document.Id, marshalledDocument)
	return nil
}
func addGeneralDocInBlockchain(stub shim.ChaincodeStubInterface) error {
	document, e := NewDocument(
		"TypeGeneral",
		TypeGeneral,
		"Л.И. Сучкова",
		"Administration",
		"{}",
		[]string{"С.М. Старолетов", "С.А. Кантор"})
	if e != nil {
		return e
	}

	document.Attributes = &DocAttributes{
		Content: "Some custom content",
	}
	marshalledDocument, _ := json.Marshal(document)
	_ = stub.PutState("doc"+document.Id, marshalledDocument)
	return nil
}
func (token *DocsChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init executed")
	// add sample docs
	if e := addGraduationThesisTopicsDocInBlockchain(stub); e != nil {
		return errorResponse(e.Error(), 3)
	}
	if e := addGraduatedExpellingDocInBlockchain(stub); e != nil {
		return errorResponse(e.Error(), 3)
	}
	if e := addPracticePermissionDocInBlockchain(stub); e != nil {
		return errorResponse(e.Error(), 3)
	}
	if e := addGeneralDocInBlockchain(stub); e != nil {
		return errorResponse(e.Error(), 3)
	}
	return successResponse("Initializing successful.")
}

func (token *DocsChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, ", args=", strings.Join(args, ","))

	return NewFunction(function, args, stub).Execute()
}
