package servers

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/habuvo/webservice/pkg/handlers"
)

func NewBaseRouter(storage handlers.Storage) (router *http.ServeMux) {
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		//some useness checks
		time.Sleep(10 * time.Second)
		isReady.Store(true)
	}()

	hb := handlers.NewBundle(storage)

	r := http.NewServeMux()

	// OPTIONS
	// r.Handle("/*", http.HandlerFunc(handlers.Preflight))
	r.Handle("/create", handlers.CORSHandler(http.HandlerFunc(hb.Create)))
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))

	return
}

// healthz is a liveness probe.
func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// readyz is a readiness probe.
func readyz(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
