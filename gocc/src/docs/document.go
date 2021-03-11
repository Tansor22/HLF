package main

import "time"

const (
	WaitingForApproval = "WAITING_FOR_APPROVAL"
	Approved           = "APPROVED"
	Closed             = "CLOSED"
	Rejected           = "REJECTED"
)

type Document struct {
	DocumentId     string    `json:"documentId"`
	Organization   string    `json:"org"`
	Date           time.Time `json:"date"`
	Content        string    `json:"content"`
	DocumentStatus string    `json:"status"`
	SignsRequired  []string  `json:"signsRequired"`
	SignedBy       []string  `json:"signedBy"`
}

func (d Document) IsSigned() bool {
	for _, signRequired := range d.SignsRequired {
		if !contains(d.SignedBy, signRequired) {
			// no sign required
			return false
		}
	}
	// document is signed
	d.DocumentStatus = Approved
	return true
}
