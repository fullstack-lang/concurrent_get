package main

import (
	"embed"
	"encoding/json"
	"fmt"
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
	http.HandleFunc("/time", currentTimeHandler)

	// Access the embedded files in the "../ng/dist/ng" directory
	subFS, err := fs.Sub(embeddedFiles, "ng/dist/ng")
	if err != nil {
		panic(err)
	}

	// Serve embedded static files
	fs := http.FileServer(http.FS(subFS))
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
