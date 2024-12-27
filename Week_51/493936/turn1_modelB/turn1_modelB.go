package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	// "github.com/pkg/errors"
)

type DatabaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database error: %s (code: %d)", e.Message, e.Code)
}

func CheckAge(age int) error {
	if age < 18 {
		return fmt.Errorf("You must be at least 18 years old to register")
	}
	return nil
}

func OpenFile(filename string) (file *os.File, err error) {
	file, err = os.Open(filename)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}
	return file, nil
}

func HandleError(err error) {
	if errors.Is(err, io.EOF) {
		fmt.Println("Reached the end of the file")
	} else {
		fmt.Println("An error occurred:", err)
	}
}

func HandleError1(err error) {
	var databaseError *DatabaseError
	if errors.As(err, &databaseError) {
		fmt.Println("Database error code:", databaseError.Code)
		fmt.Println("Database error message:", databaseError.Message)
	} else {
		fmt.Println("An error occurred:", err)
	}
}

func ReadFile(filename string) (content []byte, err error) {
	content, err = ioutil.ReadFile(filename)
	if err != nil {
		err = errors.Join(err, fmt.Errorf("Failed to read file"))
	}
	return content, err
}

func HandleError2(err error) {
	originalError := errors.Unwrap(err)
	if originalError != nil {
		fmt.Println("Original error:", originalError)
	} else {
		fmt.Println("An error occurred:", err)
	}
}

func main() {
	d := DatabaseError{
		Code:    101,
		Message: "some database error",
	}
	fmt.Println(d.Error())
	fmt.Println(CheckAge(10))
	_, e := OpenFile("example.txt")
	fmt.Println(e.Error())
	HandleError(io.EOF)
	HandleError(fmt.Errorf("%s", "some file error"))
	HandleError1(fmt.Errorf("%s", d.Error()))
	ReadFile("example.txt")
	HandleError2(fmt.Errorf("%s", d.Error()))
}
