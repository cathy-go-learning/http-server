package metrics

import (
	"fmt"
	"github.com/cathy-go-learning/http-server/pkg/envflag"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func Register() {
	err := prometheus.Register(functionLatency)
	if err != nil {
		fmt.Println(err)
	}
}

// ExecutionTimer measures execution time of a computation, split into major steps
// usual usage pattern is: timer := NewExecutionTimer(...) ; compute ; timer.ObserveStep() ; ... ; timer.ObserveTotal()
type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	start time.Time
	last  time.Time
}

var (
	functionLatency = CreateExecutionTimeMetric(envflag.PrometheusNamespace, "Time spent.")
)

const (
	// MaxVpaSizeLog - The metrics will distinguish VPA sizes up to 2^MaxVpaSizeLog (~1M)
	// Anything above that size will be reported in the top bucket.
	MaxVpaSizeLog = 20
)

// NewExecutionTimer provides a timer for Updater's RunOnce execution
func NewTimer() *ExecutionTimer {
	return NewExecutionTimer(functionLatency)
}

// NewExecutionTimer provides a timer for admission latency; call ObserveXXX() on it to measure
func NewExecutionTimer(histo *prometheus.HistogramVec) *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo: histo,
		start: now,
		last:  now,
	}
}

// ObserveStep measures the execution time from the last call to the ExecutionTimer
func (t *ExecutionTimer) ObserveStep(step string) {
	now := time.Now()
	(*t.histo).WithLabelValues(step).Observe(now.Sub(t.last).Seconds())
	t.last = now
}

// ObserveTotal measures the execution time from the creation of the ExecutionTimer
func (t *ExecutionTimer) ObserveTotal() {
	(*t.histo).WithLabelValues("total").Observe(time.Since(t.start).Seconds())
}

// CreateExecutionTimeMetric prepares a new histogram labeled with execution step
func CreateExecutionTimeMetric(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: envflag.PrometheusNamespace,
			Name:      "execution_latency_seconds",
			Help:      help,
			Buckets: []float64{0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1.0, 2.0, 5.0, 10.0,
				20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0, 90.0, 100.0, 120.0, 150.0, 240.0, 300.0},
		}, []string{"step"},
	)
}
