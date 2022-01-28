package domain

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	MethodHistogram *prometheus.HistogramVec
}

var METRICS Metrics

func Setup() {
	methodHistogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "method_histogram",
			Help:    "Histogram or latency request",
			Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"method"},
	)
	prometheus.MustRegister(methodHistogram)

	METRICS = Metrics{
		MethodHistogram: methodHistogram,
	}
}
