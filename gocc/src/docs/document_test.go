package main

import (
	"fmt"
	"testing"
	"time"
)

func TestIsSigned(t *testing.T) {
	dummyDoc := Document{
		DocumentId:     "1",
		Date:           time.Now(),
		Content:        "Some content",
		DocumentStatus: WaitingForApproval,
	}
	if !dummyDoc.IsSigned() {
		fmt.Println("Docs with no signs should be considered approved")
		t.Fail()
	}
}
