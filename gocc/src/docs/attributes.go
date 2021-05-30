package main

import "encoding/json"

const (
	FullTime  = "FULL_TIME"  // очный
	PartTime  = "PART_TIME"  // очно-заочный
	SelfStudy = "SELF_STUDY" // заочный
)

type IDocAttributes interface {
	// Текст документа
	GenerateContent()
}

// ===== General ====
type DocAttributes struct {
	// Текст документа
	Content string `json:"content"`
}

func (attrs *DocAttributes) GenerateContent() {}

// ===== GraduationThesisTopics ====

type GraduationThesisTopicsAttributes struct {
	Content    string                          `json:"content"`
	Group      string                          `json:"group"`
	Speciality string                          `json:"speciality"`
	StudyType  string                          `json:"studyType"`
	Students   []GraduationThesisTopicsStudent `json:"students"`
}

func (attrs *GraduationThesisTopicsAttributes) GenerateContent() {}

// ===== GraduatedExpelling ====

type GraduatedExpellingAttributes struct {
	Content       string                      `json:"content"`
	Course        *int                        `json:"course"`
	Speciality    string                      `json:"speciality"`
	Faculty       string                      `json:"faculty"`
	StudyType     string                      `json:"studyType"`
	Qualification string                      `json:"qualification"`
	Students      []GraduatedExpellingStudent `json:"students"`
}

func (attrs *GraduatedExpellingAttributes) GenerateContent() {}

// ===== PracticePermission ====

type PracticePermissionAttributes struct {
	Content      string                      `json:"content"`
	PracticeType string                      `json:"practiceType"`
	Course       *int                        `json:"course"`
	Speciality   string                      `json:"speciality"`
	StudyType    string                      `json:"studyType"`
	DateFrom     string                      `json:"dateFrom"`
	DateTo       string                      `json:"dateTo"`
	Students     []PracticePermissionStudent `json:"students"`
}

func (attrs *PracticePermissionAttributes) GenerateContent() {}

func AttributesFromJson(_type string, attrsJson string) (IDocAttributes, error) {
	switch _type {
	case TypeGraduationThesisTopics:
		var attrs GraduationThesisTopicsAttributes
		e := json.Unmarshal([]byte(attrsJson), &attrs)
		attrs.GenerateContent()
		return &attrs, e
	case TypeGraduatedExpelling:
		var attrs GraduatedExpellingAttributes
		e := json.Unmarshal([]byte(attrsJson), &attrs)
		attrs.GenerateContent()
		return &attrs, e
	case TypePracticePermission:
		var attrs PracticePermissionAttributes
		e := json.Unmarshal([]byte(attrsJson), &attrs)
		attrs.GenerateContent()
		return &attrs, e
	default:
		var attrs DocAttributes
		e := json.Unmarshal([]byte(attrsJson), &attrs)
		return &attrs, e
	}
}
