package state

import (
	"sync"
	"time"
)

type PageData struct {
	Timestamp int64
	Content   string
}

type CrawlJob struct {
	Url    string
	Paying bool
	Result chan string
}

const (
	NonPayingDelay   = 2 * time.Second
	DefaultWorkers   = 5
	DefaultMaxCrawls = 100
)

type MyServerState struct {
	NumWorkers           int
	MaxCrawlsPerHour     int
	PagesCrawledThisHour int
	LastCrawlReset       time.Time
	StateMutex           sync.Mutex // Define the mutex here
}

var MyState = MyServerState{
	NumWorkers:       DefaultWorkers,
	MaxCrawlsPerHour: DefaultMaxCrawls,
	LastCrawlReset:   time.Now(),
}

var MyCrawledPages = make(map[string]PageData)
