package main

import (
	"ericivan/crawler/engine"
	"ericivan/crawler/scheduler"
	"ericivan/crawler/zhenai/parser"
	"ericivan/crawler/persist"
)

func main() {

	webSite := "https://www.xl720.com/category/dongzuopian"

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
		ItemChan: persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        webSite,
		ParserFunc: parser.ParseMoveList,
	})

}
