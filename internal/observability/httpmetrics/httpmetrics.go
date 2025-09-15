package httpmetrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	apiRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "aiops_api_requests_total",
			Help: "Total API requests received",
		},
		[]string{"route", "method", "code"},
	)
	apiLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "aiops_api_request_duration_seconds",
			Help:    "HTTP request latency",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"route", "method"},
	)
)

func init() {
	prometheus.MustRegister(apiRequests, apiLatency)
}

// StatusRecorder captures response codes for metrics.
type StatusRecorder struct {
	http.ResponseWriter
	Code int
}

func (sr *StatusRecorder) WriteHeader(c int) {
	sr.Code = c
	sr.ResponseWriter.WriteHeader(c)
}

// Instrument wraps a handler to record latency and status codes.
func Instrument(route string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &StatusRecorder{ResponseWriter: w, Code: 200}
		next.ServeHTTP(rec, r)

		apiRequests.WithLabelValues(route, r.Method, http.StatusText(rec.Code)).Inc()
		apiLatency.WithLabelValues(route, r.Method).Observe(time.Since(start).Seconds())
	})
}

// Handler returns the Prometheus /metrics handler.
func Handler() http.Handler {
	return promhttp.Handler()
}
