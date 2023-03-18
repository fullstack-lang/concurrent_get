package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"
)

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	currentTime := time.Now().UTC().Format("05.999Z07:00")
	fmt.Fprint(w, currentTime)
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
