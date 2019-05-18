package main

import (
	"testing"
	"ericivan/crawler/fetcher"
	"fmt"
	"ericivan/crawler/zhenai/parser"
)

func TestProfile(t *testing.T) {

	content, err := fetcher.Fetch("http://album.zhenai.com/u/1587689291")

	//content, _ := fetcher.Fetch("http://zhenai.com/zhenghun")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println(string(content))
	parser.ParseProfile(content)
}

func TestParseCity(t *testing.T) {

	content,err:=fetcher.Fetch("http://www.zhenai.com/zhenghun/anyang")


	if err!=nil{

		fmt.Errorf(err.Error())
	}

	parser.ParseCity(content)
}
