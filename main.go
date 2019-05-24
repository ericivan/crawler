package main

import (
	"ericivan/crawler/engine"
	"ericivan/crawler/zhenai/parser"
)

func main() {

	webSite := "https://www.xl720.com/category/dongzuopian"

	var tasks []engine.Request

	tasks = append(tasks, engine.Request{
		Url:        webSite,
		ParserFunc: parser.ParseMoveList,
	})

	engine.Run(tasks...)

}
