package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := "https://www.washingtonpost.com/news-sitemap-index.xml"

	// 设置 Timeout
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// Get Resp
	resp, err := client.Get(url)
	checkErr(err)

	// Read to bytes
	bytes, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	// fmt.Println(bytes) 这个输出的是 byte 结果
	fmt.Println(string(bytes))
}
