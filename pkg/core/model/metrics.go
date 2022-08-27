package model

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	prometheusMetrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
)

// AppMetrics Global variable for all app metrics
var AppMetrics *Metrics

func init() {
	AppMetrics = &Metrics{}
	InitMetrics()
}

// Metrics defining app metrics
type Metrics struct {
	Middleware middleware.Middleware

	// RouterHTTPNotFound counter of 404 not found sent by router
	RouterHTTPNotFound prometheus.Counter

	// DatabaseErrors by message
	DatabaseErrors *prometheus.CounterVec
}

func (m *Metrics) HTTPHandler() http.Handler {
	return promhttp.Handler()
}

// InitMetrics init all app metrics
func InitMetrics() {
	AppMetrics.Middleware = middleware.New(middleware.Config{
		Recorder: prometheusMetrics.NewRecorder(prometheusMetrics.Config{}),
	})
	AppMetrics.RouterHTTPNotFound = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "app_router_http_not_found",
			Help: "How many 404 HTTP not found responses sent by router.",
		},
	)
	AppMetrics.DatabaseErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_database_errors",
			Help: "How many database errors occured, by message.",
		},
		[]string{"message"},
	)
	prometheus.MustRegister(AppMetrics.RouterHTTPNotFound)
	prometheus.MustRegister(AppMetrics.DatabaseErrors)
}

// IncDatabaseErrors by msg, by 1
func (m *Metrics) IncDatabaseErrors(msg string) {
	m.DatabaseErrors.WithLabelValues(msg).Inc()
}
