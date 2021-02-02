package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AlastorTh/httpGetRequest/models"
)

func main() {
	s := getContentAsString("https://httpbin.org/get")
	fmt.Println(s)

	str := getContentAsStruct("https://httpbin.org/get")
	fmt.Println(str)
}

func getContentAsString(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func getContentAsStruct(url string) models.Response {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	var r models.Response

	pErr := json.Unmarshal(b, &r)

	if pErr != nil {
		panic(pErr)
	}

	return r
}
