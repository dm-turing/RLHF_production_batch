package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockClient struct {
	resp *http.Response
	err  error
}

func (m *mockClient) Get(url string) (*http.Response, error) {
	return m.resp, m.err
}

func TestFetchDataSuccess(t *testing.T) {
	mockResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte("test data"))),
	}
	client := &mockClient{resp: mockResp, err: nil}

	data, err := fetchData("http://example.com", client)
	assert.NoError(t, err)
	assert.Equal(t, []byte("test data"), data)
}

func TestFetchDataError(t *testing.T) {
	client := &mockClient{resp: nil, err: errors.New("test error")}

	data, err := fetchData("http://example.com", client)
	assert.Error(t, err)
	assert.Nil(t, data)
}
