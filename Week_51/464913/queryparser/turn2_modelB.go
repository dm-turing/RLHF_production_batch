package queryparser

import (
	"errors"
	"net/url"
	"strconv"
)

// QueryParser provides a safe way to parse URL query parameters.
type QueryParser struct {
	values url.Values
}

// NewQueryParser creates a new QueryParser instance from the given query string.
func NewQueryParser(query string) (*QueryParser, error) {
	if query == "" {
		return nil, errors.New("query string is empty")
	}

	values, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}

	return &QueryParser{values: values}, nil
}

// GetString retrieves the string value for the specified key. It returns an empty string if the key is not found.
func (p *QueryParser) GetString(key string) string {
	values := p.values[key]
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

// GetInt retrieves the integer value for the specified key. It returns 0 if the key is not found or if the value cannot be converted to an integer.
func (p *QueryParser) GetInt(key string) int {
	strValue := p.GetString(key)
	if strValue == "" {
		return 0
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		return 0
	}

	return value
}

// GetFloat retrieves the float64 value for the specified key. It returns 0.0 if the key is not found or if the value cannot be converted to a float64.
func (p *QueryParser) GetFloat(key string) float64 {
	strValue := p.GetString(key)
	if strValue == "" {
		return 0.0
	}

	value, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return 0.0
	}

	return value
}

// GetBool retrieves the boolean value for the specified key. It returns false if the key is not found or if the value cannot be converted to a boolean (case-insensitive: "true", "1" are true, all others are false).
func (p *QueryParser) GetBool(key string) bool {
	strValue := p.GetString(key)
	if strValue == "" {
		return false
	}

	value, err := strconv.ParseBool(strValue)
	if err != nil {
		return false
	}

	return value
}

// GetArray retrieves all the string values for the specified key as a slice.
func (p *QueryParser) GetArray(key string) []string {
	return p.values[key]
}

// GetAll returns a map containing all the query parameters and their corresponding values.
func (p *QueryParser) GetAll() map[string][]string {
	return p.values
}
