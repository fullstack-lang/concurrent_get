package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type TimeResponse struct {
	UnixMilli int64
}

var startTime time.Time
var once sync.Once

func currentTimeHandler(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, string(responseJSON))
}

//go:embed ng/dist/ng
var embeddedFiles embed.FS

func main() {
	subFS, err := fs.Sub(embeddedFiles, "ng/dist/ng")
	if err != nil {
		panic(err)
	}

	// Create router1 for serving the static files
	router1 := gin.Default()
	router1.StaticFS("/", http.FS(subFS))
	go router1.Run(":8081")

	// Create router2 for serving the time endpoint
	router2 := gin.Default()
	router2.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})
	router2.GET("/time", currentTimeHandler)
	router2.Run(":8070")
}
