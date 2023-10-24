package state

import (
	"project-sendx/page"
	"sync"
	"time"
)

const (
	NON_PAYING_DELAY   = 2 * time.Second
	DEFAULT_WORKERS    = 5
	DEFAULT_MAX_CRAWLS = 100
)

type ServerState struct {
	NumWorkers           int
	MaxCrawlsPerHour     int
	PagesCrawledThisHour int
	LastCrawlReset       time.Time
	Mu                   sync.Mutex
}

var State = ServerState{
	NumWorkers:       DEFAULT_WORKERS,
	MaxCrawlsPerHour: DEFAULT_MAX_CRAWLS,
	LastCrawlReset:   time.Now(),
}

var CrawledPages = make(map[string]page.PageData)
