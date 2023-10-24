package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	state.MyState.StateMutex.Lock()
	state.MyState.NumWorkers = newNumWorkers
	state.MyState.StateMutex.Unlock()

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

	state.MyState.StateMutex.Lock()
	state.MyState.MaxCrawlsPerHour = newMaxCrawls
	state.MyState.StateMutex.Unlock()

	w.Write([]byte("Max crawls per hour updated successfully"))
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	state.MyState.StateMutex.Lock()
	defer state.MyState.StateMutex.Unlock()

	config := fmt.Sprintf("Number of Workers: %d, Max Crawls per Hour: %d", state.MyState.NumWorkers, state.MyState.MaxCrawlsPerHour)
	w.Write([]byte(config))
}

func GetConfigJSON(w http.ResponseWriter, r *http.Request) {
	state.MyState.StateMutex.Lock()
	defer state.MyState.StateMutex.Unlock()

	// Create a map to hold the configuration
	configData := map[string]int{
		"numWorkers":       state.MyState.NumWorkers,
		"maxCrawlsPerHour": state.MyState.MaxCrawlsPerHour,
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
