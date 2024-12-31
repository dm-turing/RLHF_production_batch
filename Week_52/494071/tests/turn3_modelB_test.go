package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPerson(t *testing.T) {
	t.Run("Get existing person", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/person?id=1", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(getPerson)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		expected := `{"id":1,"name":"Alice","age":30,"city":"New York"}`

