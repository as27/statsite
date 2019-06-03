package page

import (
	"strings"
	"testing"
)

const testPage = `
Title:   Das ist der Titel

---
# Überschrift 1

Hier steht noch mehr Text.
`

func TestReadHeader(t *testing.T) {
	r := strings.NewReader(testPage)
	header, _ := ReadHeader(r)
	wantTitle := "Das ist der Titel"
	if header.Title != wantTitle {
		t.Errorf("ReadHeader() = '%v', want %v", header.Title, wantTitle)
	}
}
