package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var requestCount = 0
var mu sync.Mutex

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/chaos", chaosHandler)

	log.Println("🔥 Chaos API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func chaosHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	requestCount++
	currentCount := requestCount
	mu.Unlock()

	// Simulate Rate Limiting
	if currentCount%5 == 0 {
		http.Error(w, "429 Too Many Requests - Simulated Rate Limit", http.StatusTooManyRequests)
		return
	}

	// Random Delay (0–3 seconds)
	delay := rand.Intn(3000)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// Random 500 Error (30% chance)
	if rand.Float32() < 0.3 {
		http.Error(w, "500 Internal Server Error - Chaos Injected", http.StatusInternalServerError)
		return
	}

	// Normal JSON Response
	response := map[string]interface{}{
		"status":  "success",
		"delay":   delay,
		"message": "Chaos response generated",
	}

	// Random JSON Corruption (20% chance)
	if rand.Float32() < 0.2 {
		w.Write([]byte("{invalid_json: true,,,}"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
