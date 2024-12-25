package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"464913/queryparser"
)

func TestQueryParser_NewQueryParser(t *testing.T) {
	_, err := queryparser.NewQueryParser("")
	assert.Error(t, err, "Empty query string should return an error")

	// _, err = queryparser.NewQueryParser("invalid query")
	// assert.Error(t, err, "Invalid query string should return an error")

	parser, err := queryparser.NewQueryParser("key=value")
	assert.NoError(t, err, "Valid query string should return no error")
	assert.NotNil(t, parser, "QueryParser instance should not be nil")
}

func TestQueryParser_GetString(t *testing.T) {
	parser, _ := queryparser.NewQueryParser("key=value")
	assert.Equal(t, "value", parser.GetString("key"), "Should return the correct string value")
	assert.Equal(t, "", parser.GetString("non-existent-key"), "Should return an empty string for a non-existent key")
}

func TestQueryParser_GetInt(t *testing.T) {
	parser, _ := queryparser.NewQueryParser("int_key=10")
	assert.Equal(t, 10, parser.GetInt("int_key"), "Should return the correct integer value")
	assert.Equal(t, 0, parser.GetInt("non-existent-key"), "Should return 0 for a non-existent key")
	assert.Equal(t, 0, parser.GetInt("string_key=value"), "Should return 0 for a string value")
}

func TestQueryParser_GetFloat(t *testing.T) {
	parser, _ := queryparser.NewQueryParser("float_key=3.14")
	assert.Equal(t, 3.14, parser.GetFloat("float_key"), "Should return the correct float value")
	assert.Equal(t, 0.0, parser.GetFloat("non-existent-key"), "Should return 0.0 for a non-existent key")
	assert.Equal(t, 0.0, parser.GetFloat("string_key=value"), "Should return 0.0 for a string value")
}

func TestQueryParser_GetBool(t *testing.T) {
	parser, _ := queryparser.NewQueryParser("bool_key=true")
	assert.True(t, parser.GetBool("bool_key"), "Should return the correct boolean value")
	assert.False(t, parser.GetBool("non-existent-key"), "Should return false for a non-existent key")
	assert.False(t, parser.GetBool("string_key=value"), "Should return false for a string value")
	assert.False(t, parser.GetBool("bool_key=false"), "Should handle lowercase false value")
}

func TestQueryParser_GetArray(t *testing.T) {
	parser, _ := queryparser.NewQueryParser("key=value1&key=value2")
	assert.Equal(t, []string{"value1", "value2"}, parser.GetArray("key"), "Should return all the values for a given key")
	assert.Equal(t, []string(nil), parser.GetArray("non-existent-key"), "Should return an empty array for a non-existent key")
}

func TestQueryParser_GetAll(t *testing.T) {
	parser, _ := queryparser.NewQueryParser("key1=value1&key2=value2&key3=value3")
	expected := map[string][]string{
		"key1": {"value1"},
		"key2": {"value2"},
		"key3": {"value3"},
	}
	assert.Equal(t, expected, parser.GetAll(), "Should return a map of all the query parameters")
}
