package handlers

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	handlersDurationMetric = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "server",
			Subsystem: "handlers",
			Name:      "duration",
			Help:      "Duration in seconds the handler took",
		},
		[]string{"name"},
	)

	handlersCountMetric = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "server",
			Subsystem: "handlers",
			Name:      "count",
			Help:      "Count of the handler was called with a result got",
		},
		[]string{"name", "status"},
	)
)
