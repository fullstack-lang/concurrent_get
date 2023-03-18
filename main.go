package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"
	"sync"
	"time"
)

type TimeResponse struct {
	UnixMilli int64
}

var startTime time.Time
var once sync.Once

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {

	// Set the Access-Control-Allow-Origin header to allow requests from any domain
	w.Header().Set("Access-Control-Allow-Origin", "*")

	once.Do(func() {
		startTime = time.Now()
	})

	elapsed := time.Since(startTime)

	now := int64(elapsed / time.Millisecond)

	timeResponse := TimeResponse{
		UnixMilli: now,
	}

	responseJSON, err := json.Marshal(timeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

//go:embed ng/dist/ng
var embeddedFiles embed.FS

func main() {

	// Access the embedded files in the "../ng/dist/ng" directory
	subFS, err := fs.Sub(embeddedFiles, "ng/dist/ng")
	if err != nil {
		panic(err)
	}

	// Create a multiplexer to handle requests on different ports
	mux := http.NewServeMux()

	// Serve embedded static files
	fs := http.FileServer(http.FS(subFS))

	mux.Handle("/", http.StripPrefix("/", fs))
	go http.ListenAndServe(":8081", mux)

	mux.HandleFunc("/time", currentTimeHandler)
	http.ListenAndServe(":8070", mux)

}
