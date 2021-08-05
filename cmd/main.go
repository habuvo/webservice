package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/habuvo/webservice/pkg/loggers"
	"github.com/habuvo/webservice/pkg/servers"
	"github.com/habuvo/webservice/pkg/storage"
	"go.uber.org/zap"
)

func main() {
	if err := loggers.NewLogger(); err != nil {
		log.Fatal(err)
	}

	lstmain := os.Getenv("LISTEN_MAIN")
	if len(lstmain) == 0 {
		lstmain = ":8080"
	}

	lstmetrics := os.Getenv("LISTEN_METRICS")
	if len(lstmetrics) == 0 {
		lstmetrics = ":8081"
	}

	storage, err := storage.NewPostgreStorage(fmt.Sprintf("host=%q user=%q password=%q dbname=%q", os.Getenv("DATABASE_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")), zap.L().Named("postgre"))
	if err != nil {
		zap.L().Fatal("Postgre init", zap.Error(err))
	}

	wg := &sync.WaitGroup{}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go servers.Serve(ctx, &http.Server{Addr: lstmain, Handler: servers.NewBaseRouter(storage)}, zap.L().With(zap.String("server", "main")), wg)

	go servers.Serve(ctx, servers.NewMetricsServer(lstmetrics, nil), zap.L().With(zap.String("server", "metrics")), wg)

	sig := <-c
	zap.L().Info("system interrupt", zap.Stringer("signal", sig))
	cancel()
}
