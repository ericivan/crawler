package main

import (
	"ericivan/crawler/engine"
	"ericivan/crawler/scheduler"
	"ericivan/crawler/zhenai/parser"
)

func main() {

	webSite := "https://www.xl720.com/category/dongzuopian"

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:        webSite,
		ParserFunc: parser.ParseMoveList,
	})

}
