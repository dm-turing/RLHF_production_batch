package parser

import (
	"strings"
	"testing"

	"github.com/example/models"
)

const xmlData = `
<people>
    <person>
        <name>John Doe</name>
        <age>30</age>
        <address>123 Main St</address>
    </person>
    <person>
        <name>Jane Smith</name>
        <age>25</age>
        <address>456 Oak St</address>
    </person>
</people>`

func TestStreamParseXML(t *testing.T) {
	reader := strings.NewReader(xmlData)
	count := 0

	err := StreamParseXML(reader, func(person models.Person) {
		count++
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if count != 2 {
		t.Fatalf("Expected 2 persons, got %d", count)
	}
}
