package main

import (
	"ericivan/crawler/engine"
	"ericivan/crawler/zhenai/parser"
)

func main() {

	webSite := "http://www.zhenai.com/zhenghun"

	engine.Run(engine.Request{
		Url:        webSite,
		ParserFunc: parser.ParseCityList,
	})


}
