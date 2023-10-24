package middlewares

import (
	"io/ioutil"
	"net/http"
	"project-sendx/state"
	"time"
)

var myPageCache = make(map[string]myCachedPage)
var content string

type myCachedPage struct {
	Content    string
	Expiration time.Time
}

func myWorker(jobs <-chan state.CrawlJob) {
	for job := range jobs {
		if !job.Paying {
			time.Sleep(state.NonPayingDelay)
		}
		job.Result <- fetchPageContent(job.Url)
	}
}

func CrawlHandler(w http.ResponseWriter, r *http.Request) {
	state.MyState.StateMutex.Lock()
	if time.Since(state.MyState.LastCrawlReset) > time.Hour {
		state.MyState.PagesCrawledThisHour = 0
		state.MyState.LastCrawlReset = time.Now()
	}
	if state.MyState.PagesCrawledThisHour >= state.MyState.MaxCrawlsPerHour {
		state.MyState.StateMutex.Unlock()
		http.Error(w, "Hourly crawl limit exceeded", http.StatusTooManyRequests)
		return
	}
	state.MyState.PagesCrawledThisHour++
	state.MyState.StateMutex.Unlock()

	urls, ok := r.URL.Query()["url"]
	if !ok || len(urls[0]) < 1 {
		http.Error(w, "Url Param 'url' is missing", http.StatusBadRequest)
		return
	}
	url := urls[0]
	isPaying := r.URL.Query().Get("isPaying") == "true"

	workersCount := state.MyState.NumWorkers
	if isPaying {
		workersCount = state.DefaultWorkers
	}

	jobs := make(chan state.CrawlJob, workersCount)
	for i := 0; i < workersCount; i++ {
		go myWorker(jobs)
	}

	result := make(chan string)
	jobs <- state.CrawlJob{Url: url, Paying: isPaying, Result: result}
	if data, found := getMyCachedPage(url); found {
		content = data // Update the variable name
	} else {
		// If not found in cache, fetch the page content
		content = fetchPageContent(url) // Update the variable name

		// Store the fetched data in the cache
		storeMyCrawledPage(url, content) // Update the variable name
	}

	currentTime := time.Now().Unix()
	state.MyCrawledPages[url] = state.PageData{
		Timestamp: currentTime,
		Content:   content, // Update the field name
	}
	w.Write([]byte(content))
	close(jobs)
}

func fetchPageContent(url string) string {
	maxRetries := 3
	retryCount := 0
	var err error
	var content []byte

	for retryCount < maxRetries {
		resp, err := http.Get(url)
		if err != nil {
			retryCount++
			continue
		}
		defer resp.Body.Close()
		content, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			break // Successfully fetched the page
		}
		retryCount++
	}

	if err != nil {
		return err.Error()
	}

	return string(content)
}

func getMyCachedPage(url string) (string, bool) {
	cached, found := myPageCache[url]
	if !found || time.Since(cached.Expiration) > 60*time.Minute {
		// Data not found in cache or expired
		return "", false
	}
	return cached.Content, true
}

func storeMyCrawledPage(url, content string) {
	myPageCache[url] = myCachedPage{
		Content:    content,
		Expiration: time.Now(),
	}
}
