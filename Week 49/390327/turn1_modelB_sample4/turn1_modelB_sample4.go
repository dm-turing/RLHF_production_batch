package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Get(string) (*http.Response, error)
}

func fetchData(url string, client Client) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	fmt.Println(fetchData("https://www.github.com", &http.Client{}))
}
