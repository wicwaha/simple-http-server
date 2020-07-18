package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)

	envDir := os.Getenv("WICWAHA_XCONTENT")
	envPort := os.Getenv("WICWAHA_XPORT")

	dir := flag.String("dir", "/content", "Directory containing content")
	port := flag.Int("port", 5000, "Port to serve on")
	flag.Parse()

	if envDir != "" {
		dir = &envDir
	}
	stat, err := os.Stat(*dir)
	if err != nil || !stat.IsDir() {
		log.Fatalf("%s is not a directory: %s", err)
	}

	if envPort != "" {
		if i, err := strconv.Atoi(envPort); err != nil {
			log.Fatalf("%s is not a valid port: %s", err)
		} else {
			port = &i
		}
	}
	if *port <= 1024 || *port > 65535 {
		log.Fatalf("%d is an invalid port number", *port)
	}

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	server := http.Server{
		Addr:              fmt.Sprintf(":%d", *port),
		TLSConfig:         nil,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	log.Printf("Listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
