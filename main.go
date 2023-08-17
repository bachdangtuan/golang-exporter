package main

import (
	"net"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ipAddress = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "machine_ip_address",
			Help: "IP address of the machine.",
		},
		[]string{"ip"},
	)
)

func main() {
	prometheus.MustRegister(ipAddress)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Exporter is running!"))
	})

	http.Handle("/metrics", promhttp.Handler())

	go startExporter()

	http.ListenAndServe(":8080", nil)
}

func startExporter() {
	for {
		ip, err := getMachineIP()
		if err == nil {
			ipAddress.WithLabelValues(ip).Set(1)
		} else {
			ipAddress.WithLabelValues("unknown").Set(0)
		}
	}
}

func getMachineIP() (string, error) {
	hostname, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		return "", err
	}
	return hostname[0], nil
}
