package main

import (
	"crawler_doubanTop250/crawler"
	"fmt"
)

func main() {
	list, err := crawler.GetTop250MovieData()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		list.SaveAsJson("../data/", "test0")
	}

}
