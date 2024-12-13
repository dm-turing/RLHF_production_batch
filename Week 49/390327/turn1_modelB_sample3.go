package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	b, _ := fetchData("https://www.github.com")
	os.Stdout.Write(b)
}
