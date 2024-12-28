package main  
import (  
    "fmt"
    "net/url"
)  
// Strategy interface
type Strategy interface {  
    Parse(values url.Values) error
}  
// SearchStrategy struct
type SearchStrategy struct {  
    Query string
}  
// Parse method for SearchStrategy
func (s *SearchStrategy) Parse(values url.Values) error {  
    s.Query = values.Get("q")
    return nil
}  
// FilterStrategy struct
type FilterStrategy struct {  
    Category string
    Price    float64