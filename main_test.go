package main

import (
	"testing"
	"ericivan/crawler/fetcher"
	"fmt"
	"ericivan/crawler/zhenai/parser"
)

func TestProfile(t *testing.T) {

	content, err := fetcher.Fetch("http://album.zhenai.com/u/1303633404")

	//content, _ := fetcher.Fetch("http://zhenai.com/zhenghun")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println(string(content))
	parser.ParseProfile(content)
}
