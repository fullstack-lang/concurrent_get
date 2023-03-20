package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
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

	// Create router1 for serving the static files
	router1 := gin.Default()
	router1.Use(cors.Default())

	router1.GET("/time", currentTimeHandler)

	router1.Use(static.Serve("/", EmbedFolder(embeddedFiles, "ng/dist/ng")))
	router1.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	router1.Run(":8080")
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
