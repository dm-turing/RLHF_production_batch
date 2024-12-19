package main

import (
	"fmt"
	"main/querybuilder"
)

func main() {
	qb := querybuilder.NewQueryBuilder()
	qb.AddParam("sort", "desc")
	qb.AddParam("limit", 10)
	qb.AddParam("offset", 0)

	queryString := qb.Build()
	fmt.Println("Query String:", queryString)
}
