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
		"Приказ о темах выпускных квалификационных работ студентов ФИТ (гр. ПИ-71,-72)",
		TypeGraduationThesisTopics,
		"Л.И. Сучкова",
		"Administration",
		"{}",
		[]string{"А.С. Авдеев", "С.А. Кантор"})
	if e != nil {
		return e
	}

	t := true
	f := false
	document.Attributes = &GraduationThesisTopicsAttributes{
		Group:      "ПИ-71",
		StudyType:  FullTime,
		Speciality: "09.03.04 Программная инженерия (профиль Разработка программно-информационных систем)",
		Students: []GraduationThesisTopicsStudent{
			{
				CommonInfo:              Student{FullName: "Агафонов Анатолий Алексеевич", Nationality: "Nationality1", Group: "ПИ-71", OnGovernmentPay: &t},
				Topic:                   "Применение технологии  верификации акторных взаимодействующих систем (на примере Erlang-программ)",
				AcademicAdvisorFullName: "Старолетов С.М., к.ф.-м.н., доцент, доцент кафедры ПМ",
			},
			{
				CommonInfo:              Student{FullName: "Борзенко Максим Александрович", Nationality: "Nationality2", Group: "Group2", OnGovernmentPay: &f},
				Topic:                   "Борзенко Максим Александрович",
				AcademicAdvisorFullName: "Корней А.И. программист ООО «Энтерра-Софт»; \nКрючкова Е.Н., к.ф.-м.н., доцент, профессор каф. ПМ\n",
			},
			// put more students here
		}}
	marshalledDocument, _ := json.Marshal(document)
	_ = stub.PutState("doc"+document.Id, marshalledDocument)
	return nil
}
func addGraduatedExpellingDocInBlockchain(stub shim.ChaincodeStubInterface) error {
	document, e := NewDocument(
		"Представление-8ПИ-81-отчисление",
		TypeGraduatedExpelling,
		"А.С. Авдеев",
		"Administration",
		"{}",
		[]string{"С.А. Кантор"})
	if e != nil {
		return e
	}

	t := true
	f := false
	course := 2
	document.Attributes = &GraduatedExpellingAttributes{
		Qualification: "МАГИСТР",
		Course:        &course,
		Faculty:       "ФИТ",
		Speciality:    "09.04.04 Программная инженерия (профиль Разработка программно-информационных систем) ",
		Students: []GraduatedExpellingStudent{
			{
				CommonInfo:       Student{FullName: "Инюшин Константин Олегович", Nationality: "РФ", Group: "8ПИ-81", OnGovernmentPay: &t},
				HasHonoursDegree: &f,
				ExamDate:         "от 08.07.2020 № 06",
			},
			{
				CommonInfo:       Student{FullName: "от 08.07.2020 № 06", Nationality: "РФ", Group: "8ПИ-81", OnGovernmentPay: &t},
				HasHonoursDegree: &f,
				ExamDate:         "от 09.07 2020 № 14",
			},
			// put more students
		}}
	marshalledDocument, _ := json.Marshal(document)
	_ = stub.PutState("doc"+document.Id, marshalledDocument)
	return nil
}
func addPracticePermissionDocInBlockchain(stub shim.ChaincodeStubInterface, title string, practiceType string, course int, speciality string, from string, to string, students []PracticePermissionStudent) error {
	document, e := NewDocument(
		title,
		TypePracticePermission,
		"Л.И. Сучкова",
		"Administration",
		"{}",
		[]string{"А.С. Авдеев"})
	if e != nil {
		return e
	}

	document.Attributes = &PracticePermissionAttributes{
		PracticeType: practiceType,
		Course:       &course,
		Speciality:   speciality,
		StudyType:    "FULL_TIME",
		DateFrom:     from,
		DateTo:       to,
		Students:     students,
	}
	marshalledDocument, _ := json.Marshal(document)
	_ = stub.PutState("doc"+document.Id, marshalledDocument)
	return nil
}

func addGeneralDocInBlockchain(stub shim.ChaincodeStubInterface) error {
	document, e := NewDocument(
		"Служебная записка о самоизоляции",
		TypeGeneral,
		"Л.И. Сучкова",
		"Administration",
		"{}",
		[]string{"С.А. Кантор"})
	if e != nil {
		return e
	}

	document.Attributes = &DocAttributes{
		Content: "Довожу до вашего сведения, что в связи с возвращением из поездки по США (неблагополучного региона) я должна соблюдать режим самоизоляции в течении 14 календарных дней.",
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
	t := true
	f := false
	if e := addPracticePermissionDocInBlockchain(stub,
		"Приказ о допуске/направлении на практику студентов ФИТ (гр. ПИ-71,-72)",
		"Преддипломная", 4, "09.03.04 Программная инженерия (профиль Разработка программно-информационных систем)",
		"18.06.21", "31.06.21", []PracticePermissionStudent{
			{
				CommonInfo:       Student{FullName: "Агафонов Анатолий Алексеевич", Nationality: "Nationality1", Group: "ПИ-71", OnGovernmentPay: &t},
				PracticeLocation: "АлтГТУ, каф. ПМ, г. Барнаул (стационарная)",
				HeadFullName:     "Лукоянычев Виктор Геннадьевич, доцент каф. ПМ",
			},
			{
				CommonInfo:       Student{FullName: "Деулина Анастасия Дмитриевна", Nationality: "Nationality2", Group: "ПИ-71", OnGovernmentPay: &f},
				PracticeLocation: "АлтГТУ, каф. ПМ, г. Барнаул (стационарная)",
				HeadFullName:     "Лукоянычев Виктор Геннадьевич, доцент каф. ПМ",
			},
			// put more students
		}); e != nil {
		return errorResponse(e.Error(), 3)
	}
	if e := addPracticePermissionDocInBlockchain(stub,
		"Приказ о допуске/направлении на практику студентов ФИТ (гр. ПИ-61,-62)",
		"Производственная", 3, "09.03.04 Программная инженерия (профиль Разработка программно-информационных систем)",
		"17.06.21", "30.06.21", []PracticePermissionStudent{
			{
				CommonInfo:       Student{FullName: "Аванесян Камо Камоевич", Group: "ПИ-61", OnGovernmentPay: &t},
				PracticeLocation: "ООО «31», г. Барнаул (стационарная)",
				HeadFullName:     "Лукоянычев Виктор Геннадьевич, доцент каф. ПМ",
			},
			{
				CommonInfo:       Student{FullName: "Дашин Константин Владимирович", Group: "ПИ-71", OnGovernmentPay: &f},
				PracticeLocation: "АО «Сбербанк-Технологии», г. Барнаул (стационарная)",
				HeadFullName:     "Лукоянычев Виктор Геннадьевич, доцент каф. ПМ",
			},
			// put more students
		}); e != nil {
		return errorResponse(e.Error(), 3)
	}
	if e := addPracticePermissionDocInBlockchain(stub,
		"Приказ о допуске/направлении на практику студентов ФИТ (гр. ПИ-81,-82)",
		"Учебная", 2, "09.03.04 Программная инженерия (профиль Разработка программно-информационных систем)",
		"29.06.21", "12.07.21", []PracticePermissionStudent{
			{
				CommonInfo:       Student{FullName: "Баев Александр Евгеньевич", Group: "ПИ-81", OnGovernmentPay: &t},
				PracticeLocation: "ФГБОУ ВО АлтГТУ, кафедра ПМ, г. Барнаул (стационарная)",
				HeadFullName:     "Андреева Ангелика Юрьевна, доцент каф. ПМ",
			},
			{
				CommonInfo:       Student{FullName: "Бачище Ольга Игоревна", Group: "ПИ-71", OnGovernmentPay: &t},
				PracticeLocation: "ФГБОУ ВО АлтГТУ, кафедра ПМ, г. Барнаул (стационарная)",
				HeadFullName:     "Андреева Ангелика Юрьевна, доцент каф. ПМ",
			},
			// put more students
		}); e != nil {
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
