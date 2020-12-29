package main

import (
	"flag"
	"net/http"
	"os"
	"time"
	"handler/server"
	"github.com/go-kit/kit/log"
)


var (
	fs = flag.NewFlagSet("todo_app", flag.ExitOnError)
	httpAddr = fs.String("http-addr", ":8080", "HTTP server address")
)

func main() {
	logger := log.NewJSONLogger(os.Stdout)
	logger = log.WithPrefix(logger, "ts", log.DefaultTimestamp)

	if err := fs.Parse(os.Args[1:]); err != nil {
		logger.Log("binding", "flag", "err", err)
		os.Exit(1)
	}

	server := &http.Server {
		Handler: handler.NewHandler(),
		Addr: *httpAddr,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Log("http", "server", "addr", *httpAddr)

	if err := server.ListenAndServe(); err != nil {
		logger.Log("http", "server", "err", err)
		os.Exit(1)
	}
}