package parser

import (
	"strings"
	"testing"
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

func TestParseXML(t *testing.T) {
	reader := strings.NewReader(xmlData)
	people, err := ParseXML(reader)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(people.Persons) != 2 {
		t.Fatalf("Expected 2 persons, got %d", len(people.Persons))
	}
}
