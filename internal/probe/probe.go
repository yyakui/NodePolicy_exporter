package probe

import (
    "fmt"
    "net"
    "time"
    "nodepolicy_exporter/internal/types"
)

func DoProbe(task types.Task) float64 {
    addr := fmt.Sprintf("%s:%d", task.TargetIP, task.TargetPort)
    timeout := time.Duration(task.TimeoutMs) * time.Millisecond

    conn, err := net.DialTimeout("tcp", addr, timeout)
    if err != nil {
        return 0
    }
    conn.Close()
    return 1
}

