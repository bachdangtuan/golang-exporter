package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    requestTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests.",
        },
        []string{"method"},
    )
)




func main() {
    // Đăng ký metric với Prometheus
    prometheus.MustRegister(requestTotal)

    // Tạo một HTTP handler để tăng giá trị của metric mỗi khi có request
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        requestTotal.WithLabelValues(r.Method).Inc()
        w.Write([]byte("Hello, Prometheus!"))
    })

    // Đăng ký endpoint metrics cho Prometheus scrape
    http.Handle("/metrics", promhttp.Handler())

    // Bắt đầu server
    http.ListenAndServe(":8080", nil)
}
