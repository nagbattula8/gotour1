package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

//PostRepos

func GetRepos(username string) (map[string]interface{}, error) {

	values := map[string]string{"name": "John Doe", "occupation": "gardener"}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("https://httpbin.org/post")
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(json_data))

	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}
	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(response.Body)

	defer response.Body.Close()
	m := map[string]interface{}{}
	err = json.NewDecoder(response.Body).Decode(&m)

	if err != nil {
		return nil, err
	}
	return m, nil

}

func main() {

	fmt.Println(GetRepos("name"))

	// values := map[string]string{"name": "John Doe", "occupation": "gardener"}
	// json_data, err := json.Marshal(values)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// resp, err := http.Post("https://httpbin.org/post", "application/json",
	// 	bytes.NewBuffer(json_data))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var res map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&res)

	// fmt.Println(&resp.Body, resp.Body)

	// fmt.Println(res["json"])
}
