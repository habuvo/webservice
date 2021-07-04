package servers

import (
	"context"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

const shutDownWait = 5 * time.Second

func Serve(ctx context.Context, srv *http.Server, logger *zap.Logger, wg *sync.WaitGroup) {
	defer wg.Done()

	// don't check error returned because shutdown will be called instead
	go srv.ListenAndServe()

	logger.Info("server started")

	<-ctx.Done()

	logger.Info("server stopped", zap.Error(srv.Shutdown(context.Background())))
}
