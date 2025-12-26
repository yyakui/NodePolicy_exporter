package collector

import (
    "strconv"
    "nodepolicy_exporter/internal/types"
    "nodepolicy_exporter/internal/probe"

    "github.com/prometheus/client_golang/prometheus"
)

var NetStatus = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
        Name: "nodepolicy_connectivity_status",
        Help: "Connectivity status",
    },
    []string{"task_id", "source_ip", "target_ip", "target_port", "protocol"},
)

func RunTask(task types.Task) {
    result := probe.DoProbe(task)
    NetStatus.WithLabelValues(
        task.TaskID,
        task.SourceIP,
        task.TargetIP,
        strconv.Itoa(task.TargetPort),
        task.Protocol,
    ).Set(result)
}

