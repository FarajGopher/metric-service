package models

type Metric struct {
    NodeName  string  `json:"node_name"`
    CPUUsage  float64 `json:"cpu_usage"`
    MemUsage  float64 `json:"mem_usage"`
    Timestamp int64   `json:"timestamp"`
}
