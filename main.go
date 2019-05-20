package main

import (
	"ericivan/crawler/engine"
	"ericivan/crawler/zhenai/parser"
	"ericivan/crawler/scheduler"
)

func main() {

	webSite := "http://www.zhenai.com/zhenghun"

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:        webSite,
		ParserFunc: parser.ParseCityList,
	})

}
