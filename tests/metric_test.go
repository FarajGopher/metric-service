package tests

import (
	"metric_service/models"
	"metric_service/storage"
	"testing"
)

func TestAddAndGetMetric(t *testing.T) {
    storage.AddMetric(models.Metric{NodeName: "test-node", CPUUsage: 50, MemUsage: 30, Timestamp: 123456})

    metrics := storage.GetMetrics()

    if len(metrics) == 0 {
        t.Fatal("Expected metrics to contain one item")
    }
    if metrics[0].NodeName != "test-node" {
        t.Fatalf("Expected node name to be 'test-node', got %s", metrics[0].NodeName)
    }
}
