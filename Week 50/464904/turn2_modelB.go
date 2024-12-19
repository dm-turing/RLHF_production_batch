package main  
import (  
    "fmt"
    "net/url"
    "strconv"
    "strings"
    
    "github.com/PuerkitoBio/goquery"
)
type FilterOperator string
const (  
    Eq  FilterOperator = "="
    Neq FilterOperator = "!="
    Gt  FilterOperator = ">"
    Lt  FilterOperator = "<"
    Gte FilterOperator = ">="
    Lte FilterOperator = "<="
)
type FilterCondition struct {  
    Key     string
    Operator FilterOperator
    Value   string
}  
type AggregateQuery struct {  
    Type        string         `json:"type"` // Aggregate query type like 'sum', 'count', 'average'
    Aggregator  string         `json:"aggregator"` // Field name for aggregation
    Filter      []FilterCondition `json:"filter"`
    Page        int               `json:"page"`
    Limit       int               `json:"limit"`
}  
func NewQuery(values url.Values) *AggregateQuery {  
    query := &AggregateQuery{  
        Page:   1,
        Limit:  10,
        Filter: []FilterCondition{},
    }  

    // Parse query parameters
    for key, value := range values {  
        if len(value) == 1 {  
            switch key {  
            case "type":  
                query.Type = value[0]
            case "aggregator":
                query.Aggregator = value[0]
            case "page":  
                page, err := strconv.Atoi(value[0])  
                if err == nil {  
                    query.Page = page  
                }  
            case "limit":  
                limit, err := strconv.Atoi(value[0])  
                if err == nil {  
                    query.Limit = limit  
                }
            default:
                filter := parseFilterCondition(key, value[0])
                if filter != nil {
                    query.Filter = append(query.Filter, *filter)
                }
            }  
        }  
    }  
    return query  
}  
func parseFilterCondition(key string, value string) *FilterCondition {
    parts := strings.SplitN(key, "__", 2)
    if len(parts) != 2 {
        return nil // Invalid filter condition
    }

    operator, err := parseFilterOperator(parts[1])
    if err != nil {
        return nil // Invalid operator
    }

    return &FilterCondition{
        Key:     parts[0],
        Operator: operator,
        Value:   value,
    }
}

func parseFilterOperator(operator string) (FilterOperator, error) {
    switch operator {
    case "eq":
        return Eq, nil
    case "neq":
        return Neq, nil
    case "gt":
        return Gt, nil
    case "lt":
        return Lt, nil
    case "gte":
        return Gte, nil
    case "lte":
        return Lte, nil
    default:
        return "", fmt.Errorf("invalid filter operator: %s", operator)
    }
}

func main() {  
    // Example usage:
    urlValues := url.Values{}
	