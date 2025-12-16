package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Total numbers of HTTP Requests",
	},
	[]string{"path"},
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestTotal.WithLabelValues(r.URL.Path).Inc()
	w.Write([]byte("Hello!"))
}
func main() {
	prometheus.MustRegister(httpRequestTotal)
	http.HandleFunc("/", indexHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
