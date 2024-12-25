package tests

import (
	"testing"

	"464913/urlquery"
)

func TestEncodeParameter(t *testing.T) {
	key := "Hello World"
	value := "This is a test"
	expected := "Hello+World=This+is+a+test"
	result := urlquery.EncodeParameter(key, value)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestEncodeParameters(t *testing.T) {
	params := urlquery.QueryParams{
		"name":  []string{"John Doe"},
		"age":   []string{"30"},
		"email": []string{"john@example.com"},
	}
	expected := "name=John+Doe&age=30&email=john%40example.com"
	result := urlquery.EncodeParameters(params)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestDecodeParameter(t *testing.T) {
	key := "Hello+World"
	value := "This+is+a+test"
	expectedKey := "Hello World"
	expectedValue := "This is a test"
	decodedKey, err := urlquery.DecodeParameter(key, value)
	if err != nil {
		t.Errorf("Error decoding: %s", err)
	}
	if decodedKey != expectedKey {
		t.Errorf("Expected key %s, got %s", expectedKey, decodedKey)
	}
	decodedValue, err := urlquery.DecodeParameter("", value)
	if err == nil {
		t.Errorf("Expected error, got %s", decodedValue)
	}
	if decodedValue != expectedValue {
		t.Errorf("Expected key %s, got %s", expectedValue, decodedValue)
	}
}

func TestDecodeParameters(t *testing.T) {
	queryString := "name=John+Doe&age=30&email=john%40example.com"
	/*expected := urlquery.QueryParams{
		"name":  []string{"John Doe"},
		"age":   []string{"30"},
		"email": []string{"john@example.com"},
	}*/
	parsed, err := urlquery.DecodeParameters(queryString)
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}
	// if parsed != expected {
	// t.Errorf("Expected %v, got %v", expected, parsed)
	// }
	parsed, err = urlquery.DecodeParameters("")
	if err == nil {
		t.Errorf("Expected error, got %v", parsed)
	}
}
