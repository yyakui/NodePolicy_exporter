package types

type Task struct {
    TaskID     string `json:"task_id"`
    SourceIP   string `json:"source_ip"`
    TargetIP   string `json:"target_ip"`
    TargetPort int    `json:"target_port"`
    Protocol   string `json:"protocol"`
    TimeoutMs  int    `json:"timeout_ms"`
}

