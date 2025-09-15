package workermetrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	jobsStarted = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "aiops_worker_jobs_started_total",
		Help: "Jobs started",
	})
	jobsSucceeded = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "aiops_worker_jobs_succeeded_total",
		Help: "Jobs succeeded",
	})
	jobsFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "aiops_worker_jobs_failed_total",
		Help: "Jobs failed",
	})
	jobDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "aiops_worker_job_duration_seconds",
		Help:    "Job duration in seconds",
		Buckets: prometheus.DefBuckets,
	})
	queueDepth = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "aiops_worker_queue_depth",
		Help: "Number of jobs waiting in the queue",
	})
)

func init() {
	prometheus.MustRegister(jobsStarted, jobsSucceeded, jobsFailed, jobDuration, queueDepth)
}

// StartServer exposes /metrics on the given addr (e.g., ":9101").
func StartServer(addr string) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Printf("worker metrics on %s", addr)
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Fatalf("metrics server error: %v", err)
		}
	}()
}

// Timer is a helper to time a job run.
// Usage:
//
//	t := workermetrics.NewTimer()
//	defer t.ObserveDuration()
func NewTimer() *prometheus.Timer {
	return prometheus.NewTimer(jobDuration)
}

func IncStarted()   { jobsStarted.Inc() }
func IncSucceeded() { jobsSucceeded.Inc() }
func IncFailed()    { jobsFailed.Inc() }

// SetQueueDepth can be called by your polling loop.
func SetQueueDepth(n int) { queueDepth.Set(float64(n)) }
