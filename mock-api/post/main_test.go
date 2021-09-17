package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// Custom type that allows setting the func that our Mock Do func will run instead
type MockDoType func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
type MockClient struct {
	MockDo MockDoType
}

// Overriding what the Do function should "do" in our MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}
func TestGitHubCallSuccess(t *testing.T) {
	// build our response JSON
	jsonResponse := `{
   "full_name": "mock-repo"
  }`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
	result, err := GetRepos("atkinsonbg")
	if err != nil {

		fmt.Println(err, " This is the error")

		t.Error("TestGitHubCallSuccess failed.")
		return
	} else {
		fmt.Println("Passed testcase", result)
	}
}
func TestGitHubCallFail(t *testing.T) {
	// create a client that throws and returns an error
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       nil,
			}, errors.New("Mock Error")
		},
	}
	_, err := GetRepos("atkinsonbgthisusershouldnotexist")
	if err == nil {
		t.Error("TestGitHubCallFail failed.")
		return
	}
}
