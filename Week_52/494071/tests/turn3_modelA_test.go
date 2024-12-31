package main

import (
	"encoding/json"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMarshalPerson(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	want := `{"name":"Alice","age":30}`
	got, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("json.Marshal error: %v", err)
	}
	if string(got) != want {
		t.Errorf("json.Marshal(%v) = %q, want %q", p, got, want)
	}
}
