package main

import (
	"fmt"
	"strings"
)

func generateHTMLList(items []string) string {
	var buf strings.Builder
	buf.WriteString("<ul>\n")

	for _, item := range items {
		buf.WriteString(fmt.Sprintf("  <li>%s</li>\n", item))
	}

	buf.WriteString("</ul>\n")
	return buf.String()
}

func main() {
	items := []string{"Apple", "Banana", "Cherry", "Date"}
	htmlList := generateHTMLList(items)
	fmt.Println(htmlList)
}
