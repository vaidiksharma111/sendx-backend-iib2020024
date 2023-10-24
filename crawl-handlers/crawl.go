package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"project-sendx/page"
	"project-sendx/state"
	"time"
)

var pageCache = make(map[string]cachedPage)

type cachedPage struct {
	Data       string
	ExpiryTime time.Time
}

func worker(jobs <-chan page.CrawlJob) {
	for job := range jobs {
		if !job.IsPaying {
			time.Sleep(state.NON_PAYING_DELAY)
		}
		job.Result <- crawlPage(job.Url)
	}
}

func CrawlHandler(w http.ResponseWriter, r *http.Request) {
	state.State.Mu.Lock()
	if time.Since(state.State.LastCrawlReset) > time.Hour {
		state.State.PagesCrawledThisHour = 0
		state.State.LastCrawlReset = time.Now()
	}
	if state.State.PagesCrawledThisHour >= state.State.MaxCrawlsPerHour {
		state.State.Mu.Unlock()
		http.Error(w, "Hourly crawl limit exceeded", http.StatusTooManyRequests)
		return
	}
	state.State.PagesCrawledThisHour++
	state.State.Mu.Unlock()

	urls, ok := r.URL.Query()["url"]
	if !ok || len(urls[0]) < 1 {
		http.Error(w, "Url Param 'url' is missing", http.StatusBadRequest)
		return
	}
	url := urls[0]
	isPaying := r.URL.Query().Get("isPaying") == "true"

	workersCount := state.State.NumWorkers
	if isPaying {
		workersCount = state.DEFAULT_WORKERS
	}

	jobs := make(chan page.CrawlJob, workersCount)
	for i := 0; i < workersCount; i++ {
		go worker(jobs)
	}

	result := make(chan string)
	jobs <- page.CrawlJob{Url: url, IsPaying: isPaying, Result: result}
	pageData := <-result

	// Check if the URL has been crawled recently and exists in the cache
	if data, found := getCachedPage(url); found {
		fmt.Println("content found in cache")
		pageData = data
	} else {
		// If not found in cache, crawl the page
		pageData = crawlPage(url)

		// Store the crawled data in the cache
		storeCrawledPage(url, pageData)
	}

	currentTime := time.Now().Unix()
	state.CrawledPages[url] = page.PageData{
		Timestamp: currentTime,
		Data:      pageData,
	}
	w.Write([]byte(pageData))
	close(jobs)
}

func crawlPage(url string) string {
	maxRetries := 3
	retryCount := 0
	var err error
	var body []byte

	for retryCount < maxRetries {
		fmt.Println("try: ", retryCount)
		resp, err := http.Get(url)
		if err != nil {
			retryCount++
			continue
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			break // Successfully fetched the page
		}
		retryCount++
	}

	if err != nil {
		return err.Error()
	}

	return string(body)
}

func getCachedPage(url string) (string, bool) {
	cached, found := pageCache[url]
	if !found || time.Since(cached.ExpiryTime) > 60*time.Minute {
		// Data not found in cache or expired
		return "", false
	}
	return cached.Data, true
}

func storeCrawledPage(url, data string) {
	pageCache[url] = cachedPage{
		Data:       data,
		ExpiryTime: time.Now(),
	}
}
