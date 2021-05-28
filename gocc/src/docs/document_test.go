package main

import (
	"fmt"
	"testing"
	"time"
)

func TestIsSigned(t *testing.T) {
	attrs, _ := AttributesFromJson("", "{\"content\": \"Some content\"}")
	dummyDoc := Document{
		Id:         "1",
		Date:       time.Now(),
		Attributes: attrs,
		Status:     Processing,
	}
	if !dummyDoc.IsSigned() {
		fmt.Println("Docs with no signs should be considered approved")
		t.Fail()
	}
}
