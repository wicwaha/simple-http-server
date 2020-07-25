package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

var Version = "dev"

func main() {
	log.SetFlags(log.LUTC | log.Ldate | log.Ltime)

	envDir := os.Getenv("CONTENT_DIR")
	envPort := os.Getenv("PORT")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	dir := flag.String("content-dir", ".", "Directory to serve over HTTP")
	port := flag.Int("port", 5000, "Port to serve on")
	flag.Parse()

	log.Printf("%s %s", path.Base(os.Args[0]), Version)

	if envDir != "" {
		dir = &envDir
	}

	stat, err := os.Stat(*dir)

	if err != nil || !stat.IsDir() {
		log.Fatalf("%s is not a directory: %s", *dir, err)
	}

	if envPort != "" {
		if i, err := strconv.Atoi(envPort); err != nil {
			log.Fatalf("%s is not a valid port: %s", envPort, err)
		} else {
			port = &i
		}
	}

	if *port <= 1024 || *port > 65535 {
		log.Fatalf("invalid port %d, must be between (1024, 65535]", *port)
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
