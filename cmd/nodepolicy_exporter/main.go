package main

import (
    "nodepolicy_exporter/internal/collector"
    "nodepolicy_exporter/internal/server"

    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    prometheus.MustRegister(collector.NetStatus)
    server.StartServer()
}

