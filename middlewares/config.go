package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "project-sendx/models"
	"project-sendx/state"
	"strconv"
)

func UpdateNumWorkers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read the request body", http.StatusBadRequest)
		return
	}

	newNumWorkers, err := strconv.Atoi(string(body))
	if err != nil {
		http.Error(w, "Invalid number provided", http.StatusBadRequest)
		return
	}

	state.State.Mu.Lock()
	state.State.NumWorkers = newNumWorkers
	state.State.Mu.Unlock()

	w.Write([]byte("Number of workers updated successfully"))
}

func UpdateMaxCrawlsPerHour(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read the request body", http.StatusBadRequest)
		return
	}

	newMaxCrawls, err := strconv.Atoi(string(body))
	if err != nil {
		http.Error(w, "Invalid number provided", http.StatusBadRequest)
		return
	}

	state.State.Mu.Lock()
	state.State.MaxCrawlsPerHour = newMaxCrawls
	state.State.Mu.Unlock()

	w.Write([]byte("Max crawls per hour updated successfully"))
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	state.State.Mu.Lock()
	defer state.State.Mu.Unlock()

	config := fmt.Sprintf("Number of Workers: %d, Max Crawls per Hour: %d", state.State.NumWorkers, state.State.MaxCrawlsPerHour)
	w.Write([]byte(config))
}

func GetConfigJSON(w http.ResponseWriter, r *http.Request) {
	state.State.Mu.Lock()
	defer state.State.Mu.Unlock()

	// Create a map to hold the configuration
	configData := map[string]int{
		"numWorkers":       state.State.NumWorkers,
		"maxCrawlsPerHour": state.State.MaxCrawlsPerHour,
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