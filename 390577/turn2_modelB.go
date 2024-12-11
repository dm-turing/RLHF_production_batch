package main

import (
	"fmt"

	"github.com/hashicorp/go-semver/semver"
)

func main() {
	v1, _ := semver.NewVersion("1.2.3")
	v2, _ := semver.NewVersion("1.2.4")

	fmt.Println(v1.GreaterThan(v2)) // false
	fmt.Println(v2.GreaterThan(v1)) // true
}
