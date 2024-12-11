package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(uuid.Max)
	fmt.Println(logrus.GetLevel().String())
	fmt.Println("Hello, World!")
}
