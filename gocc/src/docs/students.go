package main

type Student struct {
	// all string should be substitute to pointers?
	FullName        string `json:"fullName"`
	Nationality     string `json:"nationality"`
	Group           string `json:"group"`
	OnGovernmentPay *bool  `json:"onGovernmentPay"` // основа обучения бюджет=true\внебюджет=false
}

type GraduatedExpellingStudent struct {
	CommonInfo       Student `json:"commonInfo"`
	HasHonoursDegree *bool   `json:"hasHonoursDegree"`
	ExamDate         string  `json:"examDate"`
}

type GraduationThesisTopicsStudent struct {
	CommonInfo              Student `json:"commonInfo"`
	Topic                   string  `json:"topic"`
	AcademicAdvisorFullName string  `json:"academicAdvisorFullName"`
}

type PracticePermissionStudent struct {
	CommonInfo       Student `json:"commonInfo"`
	PracticeLocation string  `json:"practiceLocation"`
	HeadFullName     string  `json:"headFullName"`
}
