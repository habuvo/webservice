package storage

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	handlersDurationMetric = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "server",
			Subsystem: "storage",
			Name:      "duration",
			Help:      "Duration in seconds the method took",
		},
		[]string{"name"},
	)

	handlersCountMetric = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "server",
			Subsystem: "handlers",
			Name:      "count",
			Help:      "Count of the method was called with a result got",
		},
		[]string{"name", "status"},
	)
)
