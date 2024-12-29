package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

var (
	// cache for transformed data
	cache sync.Map
)

func main() {
	router := gin.Default()
	router.POST("/data", middleware, handler)
	log.Println("Server starting on :8080")
	log.Fatal(router.Run(":8080"))
}

func handler(c *gin.Context) {
	// Assuming the request body is in XML format
	data := c.MustGet("data").(RequestData)

	// Do something with the data
	c.JSON(http.StatusOK, gin.H{"name": data.Name, "age": data.Age})
}

func middleware(c *gin.Context) {
	contentType := c.GetHeader("Content-Type")

	if !strings.HasPrefix(contentType, "application/xml") {
		// Skip if the content type is not XML
		c.Next()
		return
	}

	// Get the raw request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	defer c.Request.Body.Close()

	// Check if the data is already in the cache
	if value, ok := cache.Load(string(body)); ok {
		c.Set("data", value)
		c.Next()
		return
	}

	var data RequestData
	body = []byte(`<data><name>Immortal</name><age>100</age></data>`)

	if err := xml.Unmarshal(body, &data); err != nil {
		log.Println("Error unmarshaling XML:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Perform conversion to JSON
	jsonBody, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Store the transformed data in the cache
	cache.Store(string(body), data)

	// Modify the request to have a JSON body
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Proceed to the next handler
	c.Next()
}
