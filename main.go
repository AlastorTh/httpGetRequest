package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	s := getContentAsString("https://httpbin.org/get")
	fmt.Println(s)
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
