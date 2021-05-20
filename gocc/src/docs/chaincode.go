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
	var docsMarshalled []string
	for i := 0; i < 1; i++ {
		document := NewDocument(
			"Приказ о темах выпускных квалификационных работ студентов",
			PracticePermission,
			"Л.И. Сучкова",
			"Administration",
			"В соответствии с учебным планом и СК ОПД 01-139-2019 «Положение о выпускной квалификационной работе по образовательным программам высшего образования – программам бакалавриата, программам специалитета, программам магистратуры» утвердить темы и руководителей выпускных квалификационных работ студентам факультета информационных технологий групп ПИ-71,-72 очной формы обучения по направлению бакалавриата 09.03.04 Программная инженерия (профиль Разработка программно-информационных систем) согласно списку:\n\n"+
				"Группа ПИ-71\n\n"+
				"Агафонов Анатолий Алексеевич, Тема: Применение технологии верификации акторных взаимодействующих систем (на примере Erlang-программ), науч.рук.:Старолетов С.М., к.ф.-м.н., доцент, доцент кафедры ПМ\n\n"+
				"Куторкин Артем Сергеевич; Борзенко Максим Александрович, Тема: Проектирование и реализация системы распознавания сложных изображений на основе метода семантической коррекции, науч. рук.: Корней А.И. программист ООО «Энтерра-Софт»;Крючкова Е.Н., к.ф.-м.н., доцент, профессор каф. ПМ\n\n",
			[]string{"С.М. Старолетов", "С.А. Кантор"})
		marshalledDocument, _ := json.Marshal(document)
		_ = stub.PutState("doc"+document.Id, marshalledDocument)
		docsMarshalled = append(docsMarshalled, string(marshalledDocument))
	}
	j, _ := json.Marshal(docsMarshalled)
	return successResponse("Initializing successful. Docs added: " + string(j))
}

func (token *DocsChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, ", args=", strings.Join(args, ","))

	return NewFunction(function, args, stub).Execute()
}
