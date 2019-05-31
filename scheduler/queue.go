package scheduler

import "ericivan/crawler/engine"

type QueueSchedule struct {
	WorkerChan []chan engine.Request

}
