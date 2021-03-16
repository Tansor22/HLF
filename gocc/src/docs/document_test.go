package main

import (
	"fmt"
	"testing"
	"time"
)

func TestIsSigned(t *testing.T) {
	dummyDoc := Document{
		Id:      "1",
		Date:    time.Now(),
		Content: "Some content",
		Status:  WaitingForApproval,
	}
	if !dummyDoc.IsSigned() {
		fmt.Println("Docs with no signs should be considered approved")
		t.Fail()
	}
}
