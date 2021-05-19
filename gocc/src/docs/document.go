package main

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

const (
	// doc.status
	Processing    = "PROCESSING"
	Approved      = "APPROVED"
	Closed        = "CLOSED"
	Rejected      = "REJECTED"
	InitialStatus = Processing
	// change.type
	Reject  = "REJECT"
	Edit    = "EDIT"
	Approve = "APPROVE"
	// doc.type
	GraduatedExpelling     = "GraduatedExpelling"
	PracticePermission     = "PracticePermission"
	GraduationThesisTopics = "GraduationThesisTopics"
)

type Change struct {
	// Dean
	Member string `json:"documentId"`
	// REJECT
	Type string `json:"type"`
	// 21.05.2021
	Date time.Time `json:"date"`
	// Отсутствует студент Иванов И.И.
	Details string `json:"details"`
}

type Document struct {
	Id    string `json:"documentId"`
	Title string `json:"title"`
	// User identity
	Owner string `json:"owner"`
	// группа: администрация, сервис
	Group string `json:"group"`
	// GraduatedExpelling - Представление-<группа(ы)>-отчисление
	// PracticePermission - Приказ о допуске на практику студентов
	// GraduationThesisTopics - Приказ о темах выпускных квалификационных работ
	Type string `json:"type"`
	// Дата создания
	Date time.Time `json:"date"`
	// Текст документа
	Content string `json:"content"`
	// PROCESSING - на рассмотрении (не подписан и не отклонен)
	// APPROVED - подписан всеми участниками, финальный статус
	// REJECTED - отклонен участником с комментарием
	// ABORTED - отмененный
	Status string `json:"status"`
	// История изменений по типу:
	/*
		[
		  {
		    "member": "Dean",
		    "type": "REJECT",
		    "date": "21.05.2021",
		    "details": "Отсутствует студент Иванов И.И."
		  },
		   {
		    "member": "Owner",
		    "type": "EDIT",
		    "date": "22.05.2021",
		    "details": "Добавлен студент Иванов И.И."
		  },
		   {
		    "member": "Dean",
		    "type": "APPROVE",
		    "date": "23.05.2021",
		    "details": null
		  }
		]
	*/
	Changes []Change `json:"changes"`
	// Список участников чьи подписи требуются в порядке указанном в списке
	SignsRequired []string `json:"signsRequired"`
	// Список участников чьи подписи уже поставлены (кем одобрен документ)
	SignedBy []string `json:"signedBy"`
}

func NewChange(member string, _type string, details string) Change {
	return Change{
		Member:  member,
		Type:    _type,
		Date:    time.Now(),
		Details: details,
	}
}
func NewDocument(title string, _type string, owner string, group string, content string, signs []string) Document {
	return Document{
		Id:            uuid.NewString(),
		Title:         title,
		Type:          _type,
		Owner:         owner,
		Group:         group,
		Date:          time.Now(),
		Content:       content,
		Status:        InitialStatus,
		SignsRequired: signs,
		SignedBy:      make([]string, 0),
		Changes:       make([]Change, 0),
	}
}

func (d *Document) IsSigned() bool {
	for _, signRequired := range d.SignsRequired {
		if !contains(d.SignedBy, signRequired) {
			// no sign required
			return false
		}
	}
	// document is signed
	d.Status = Approved
	return true
}

func (d *Document) RegisterChange(change Change) error {
	switch change.Type {
	case Approve:
		if !contains(d.SignsRequired, change.Member) {
			marshalledSigns, _ := json.Marshal(d.SignsRequired)
			return errors.New("Sign " + change.Member + " is not applicable to the doc, only " + string(marshalledSigns) + " are")
		}
		if sign, err := d.getCurrentSign(); err != nil {
			return err
		} else {
			if sign != change.Member {
				return errors.New("Out of queue, current sign should be " + sign)
			}
			d.SignedBy = append(d.SignedBy, change.Member)
			d.Changes = append(d.Changes, change)
		}
	case Reject:
		if !contains(d.SignsRequired, change.Member) {
			marshalledSigns, _ := json.Marshal(d.SignsRequired)
			return errors.New("Reject " + change.Member + " is not applicable to the doc, only " + string(marshalledSigns) + " are")
		}
		if sign, err := d.getCurrentSign(); err != nil {
			return err
		} else {
			if sign != change.Member {
				return errors.New("Out of queue, current sign should be " + sign)
			}
			d.Status = Rejected
			d.Changes = append(d.Changes, change)
		}
	case Edit:
		// todo how to process edit?
		return errors.New("Not supported change type yet:" + change.Type)
	}
	return nil
}

func (d *Document) getCurrentSign() (string, error) {
	if d.Status == Approved {
		return "", errors.New("already signed")
	}
	if len(d.SignedBy) == 0 {
		return d.SignsRequired[0], nil
	}
	lastSigned := d.SignedBy[len(d.SignedBy)-1]
	// next sign required
	if lastSigned != d.SignsRequired[len(d.SignsRequired)-1] {
		for i := 0; i < len(d.SignsRequired)-1; i++ {
			if d.SignsRequired[i] == lastSigned {
				return d.SignsRequired[i+1], nil
			}
		}
	}
	// last sign remained
	return d.SignsRequired[len(d.SignsRequired)-1], nil

}
