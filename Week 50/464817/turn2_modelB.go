package querybuilder

import (
    "fmt"
    "net/url"
    "reflect"
    "strings"
)

type Operator string

const (
    Equals        Operator = "="
    NotEquals      Operator = "!="
    GreaterThan    Operator = "gt"
    LessThan       Operator = "lt"
    In             Operator = "in"
    Contains       Operator = "contains"
)

type QueryBuilder struct {
    values url.Values
}

func New() *QueryBuilder {
    return &QueryBuilder{values: make(url.Values)}
}

// Set adds a query parameter with the default operator "="
func (qb *QueryBuilder) Set(fieldName string, value interface{}) *QueryBuilder {
    return qb.AddOperator(fieldName, value, Equals)
}

// Add adds a query parameter with the given operator
func (qb *QueryBuilder) AddOperator(fieldName string, value interface{}, operator Operator) *QueryBuilder {
    encodeQueryParameter(qb.values, fieldName, value, operator)
    return qb
}

// Del removes all query parameters with the given key.
func (qb *QueryBuilder) Del(param string) *QueryBuilder {
    qb.values.Del(param)
    return qb
}

func encodeQueryParameter(values url.Values, fieldName string, value interface{}, operator Operator) {
    var vals string

    // Handle slice types
    switch v := value.(type) {
    case []string:
        vals = strings.Join(v, ",")
    default:
        vals = fmt.Sprint(v)
    }

    // Apply the specified operator
    switch operator {
    case In:
        // In operator is handled differently as it requires multiple values separated by commas
        values.Set(fieldName+string(operator), vals)
    case Contains:
        values.Set(fieldName+"__"+string(operator), vals)
    default:
        values.Set(fieldName, vals+string(operator))
    }
}
 
func (qb *QueryBuilder) Build() string {
    return qb.values.Encode()
}

// BuildQueryString creates a query string from a struct instance using query tags.
func BuildQueryString(v interface{}) string {
    qb := New()
    val := reflect.ValueOf(v)