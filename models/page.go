package models

type PageData struct {
	Timestamp int64
	Data      string
}

type CrawlJob struct {
	Url      string
	IsPaying bool
	Result   chan string
}
