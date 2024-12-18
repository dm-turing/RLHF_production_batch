package main

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/stretchr/objx"
)

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	// Example complex XML data
	complexXMLData := `
	<items>
		<item>
			<id>1</id>
			<name>Apple</name>
			<details>
				<color>Red</color>
				<weight>150g</weight>
			</details>
		</item>
		<item>
			<id>2</id>
			<name>Banana</name>
			<details>
				<color>Yellow</color>
			</details>
		</item>
		<!-- More items -->
	</items>
	`

	// Parse the complex XML data using objx
	var data objx.Map
	if err := xml.Unmarshal([]byte(complexXMLData), &data); err != nil {
		log.Fatal(err)
	}

	// Pagination logic remains the same
	pageParam := r.URL.Query().Get("page")
	sizeParam := r.URL.Query().Get("size")
	// ... (rest of the pagination code)

	// Slice the items for the requested page
	startIndex := (page - 1) * size
	endIndex := min(len(data.Get("item").([]objx.Map)), startIndex+size)
	pagedItems := data.Get("item").([]objx.Map)[startIndex:endIndex]

	// Create the page response using objx
	pageResponse := objx.New(map[string]interface{}{
		"current_page": page,
		"total_pages":  totalPages,
		"total_items":  len(data.Get("item").([]objx.Map)),
		"items":        pagedItems,
	})

	// Set content type header and encode the response
	w.Header().Set("Content-Type", "application/xml")
	if err := xml.NewEncoder(w).Encode(pageResponse); err != nil {
		log.Fatal(err)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	http.HandleFunc("/items", itemsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
