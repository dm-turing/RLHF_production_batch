package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readNonClosableResource() (io.Reader, error) {
	// This example uses a non-closable reader (e.g., from a buffer)
	return ioutil.NopCloser(os.Stdin), nil
}

func main() {
	reader, err := readNonClosableResource()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reader.(io.Closer).Close() // Use io.Closer to close the resource

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
