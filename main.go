package main

import (
	"bash_gen/jobs"

	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {
	var addr = flag.String("addr", ":8080", "http service address")
	flag.Parse()
	ctx := context.Background()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "jobs",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var srv jobs.Service
	{
		repo := jobs.NewRepo(logger)
		srv = jobs.NewService(repo, logger)
	}
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	endpoints := jobs.MakeEndpoints(srv)
	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *addr)
		handler := jobs.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*addr, handler)
	}()
	level.Error(logger).Log("exit", <-errs)

}
