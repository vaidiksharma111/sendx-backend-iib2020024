package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project-sendx/models"
	"strconv"
)

func SetNumWorkers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	newNumWorkers, err := strconv.Atoi(string(body))
	// fmt.Println("number of workers assigned = ", newNumWorkers)
	if err != nil {
		http.Error(w, "Invalid number provided", http.StatusBadRequest)
		return
	}

	models.State.Mu.Lock()
	models.State.NumWorkers = newNumWorkers
	models.State.Mu.Unlock()

	w.Write([]byte("Number of workers updated successfully"))
}

func SetMaxCrawlsPerHour(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	newMaxCrawls, err := strconv.Atoi(string(body))
	// fmt.Println("new max crawls assigned = ", newMaxCrawls)
	if err != nil {
		http.Error(w, "Invalid number provided", http.StatusBadRequest)
		return
	}

	models.State.Mu.Lock()
	models.State.MaxCrawlsPerHour = newMaxCrawls
	models.State.Mu.Unlock()

	w.Write([]byte("Max crawls per hour updated successfully"))
	// fmt.Println("Max crawls per hour updated successfully")
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	models.State.Mu.Lock()
	defer models.State.Mu.Unlock()

	config := fmt.Sprintf("Number of Workers: %d, Max Crawls per Hour: %d", models.State.NumWorkers, models.State.MaxCrawlsPerHour)
	w.Write([]byte(config))
}

func GetConfigJSON(w http.ResponseWriter, r *http.Request) {
	models.State.Mu.Lock()
	defer models.State.Mu.Unlock()

	// Create a map to hold the configuration
	configData := map[string]int{
		"numWorkers":       models.State.NumWorkers,
		"maxCrawlsPerHour": models.State.MaxCrawlsPerHour,
	}

	// Convert the map to JSON
	configJSON, err := json.Marshal(configData)
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}

	// Set the response header for JSON output
	w.Header().Set("Content-Type", "application/json")
	w.Write(configJSON)
}
