package handlers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type Storage interface {
}

type HandlersBundle struct {
	storage Storage
}

func NewBundle(s Storage) *HandlersBundle {
	return &HandlersBundle{storage: s}
}
func (hb *HandlersBundle) Create(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(handlersDurationMetric.WithLabelValues("create"))
	defer timer.ObserveDuration()

	//storage operations

	handlersCountMetric.WithLabelValues("create", "ok")
	//WriteJSONObject(w, map[string]interface{}{"id": id})
}
