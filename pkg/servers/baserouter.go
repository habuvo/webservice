package servers

import (
	"net/http"

	"github.com/habuvo/webservice/pkg/handlers"
)

func NewBaseRouter(storage handlers.Storage) (router *http.ServeMux) {
	hb := handlers.NewBundle(storage)

	r := http.NewServeMux()

	// OPTIONS
	// r.Handle("/*", http.HandlerFunc(handlers.Preflight))
	r.Handle("/create", handlers.CORSHandler(http.HandlerFunc(hb.Create)))

	return
}
