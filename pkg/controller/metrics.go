package controller

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/renomarx/fizzbuzz/pkg/core/model"
)

type metricsController struct {
}

func NewMetricsController() *metricsController {
	return &metricsController{}
}

func (s *metricsController) GetMetrics() *model.Metrics {
	return model.AppMetrics
}

func (s *metricsController) HTTPHandler() http.Handler {
	return promhttp.Handler()
}
