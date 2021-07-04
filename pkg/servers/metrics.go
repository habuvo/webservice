package servers

import (
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMetricsServer(lst string, healthHandler http.Handler) (srv *http.Server) {
	if healthHandler == nil {
		healthHandler = http.HandlerFunc(defaultHealthHandler)
	}

	mux := http.NewServeMux()
	mux.Handle("/health", healthHandler)

	mux.Handle("/metrics", promhttp.Handler())

	mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	srv.Handler = mux
	srv.Addr = lst

	return
}

func defaultHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
