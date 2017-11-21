package main

import (
	"encoding/xml"
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

type Sitemap struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return l.Loc
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

	var s Sitemap
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}
