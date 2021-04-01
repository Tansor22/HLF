package main

import (
	"github.com/google/uuid"
	"time"
)

const (
	WaitingForApproval = "WAITING_FOR_APPROVAL"
	Approved           = "APPROVED"
	Closed             = "CLOSED"
	Rejected           = "REJECTED"
	InitialStatus      = WaitingForApproval
)

type Document struct {
	//TODO: add title, description
	Id            string    `json:"documentId"`
	Organization  string    `json:"org"`
	Date          time.Time `json:"date"`
	Content       string    `json:"content"`
	Status        string    `json:"status"`
	SignsRequired []string  `json:"signsRequired"`
	SignedBy      []string  `json:"signedBy"`
}

func NewDocument(org string, content string, signs []string) Document {
	return Document{
		Id:            uuid.NewString(),
		Organization:  org,
		Date:          time.Now(),
		Content:       content,
		Status:        InitialStatus,
		SignsRequired: signs,
		SignedBy:      make([]string, 0),
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
