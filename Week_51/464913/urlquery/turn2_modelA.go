package urlquery

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type QueryParams map[string][]string

// EncodeParameter safely encodes a single query parameter.
func EncodeParameter(key, value string) string {
	return url.QueryEscape(key) + "=" + url.QueryEscape(value)
}

// EncodeParameters safely encodes multiple query parameters into a query string.
func EncodeParameters(params QueryParams) string {
	values := url.Values{}
	for key, _values := range params {
		for _, value := range _values {
			values.Add(key, value)
		}
	}
	return values.Encode()
}

// DecodeParameter safely decodes a single query parameter.
func DecodeParameter(key, value string) (string, error) {
	if !isValidKey(key) || !isValidValue(value) {
		return "", fmt.Errorf("invalid key or value: %s=%s", key, value)
	}
	decodedKey, err := url.QueryUnescape(key)
	if err != nil {
		return "", fmt.Errorf("error decoding key: %s", err)
	}
	_, err = url.QueryUnescape(value)
	if err != nil {
		return "", fmt.Errorf("error decoding value: %s", err)
	}
	return decodedKey, nil
}

// DecodeParameters safely decodes multiple query parameters from a query string.
func DecodeParameters(queryString string) (QueryParams, error) {
	if !isValidQueryString(queryString) {
		return nil, fmt.Errorf("invalid query string: %s", queryString)
	}
	parsedQuery, err := url.ParseQuery(queryString)
	if err != nil {
		return nil, fmt.Errorf("error parsing query: %s", err)
	}
	return QueryParams(parsedQuery), nil
}

func isValidKey(key string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9_.-]+=$`).MatchString(key)
}

func isValidValue(value string) bool {
	return strings.TrimSpace(value) != ""
}

func isValidQueryString(queryString string) bool {
	return queryString != ""
}

func main() {
	encodedstring := EncodeParameter("key", "value1")
	fmt.Println(encodedstring)
	res, _ := DecodeParameter("key", "value1")
	fmt.Println(res)
	encodedstring = EncodeParameters(QueryParams{"key": []string{"value1", "value2"}})
	fmt.Println(encodedstring)
	r, _ := DecodeParameters(encodedstring)
	fmt.Println(r)
}
