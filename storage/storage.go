package storage

import "metric_service/models"

var (
    Metrics []models.Metric
)

func AddMetric(metric models.Metric) {
    Metrics = append(Metrics, metric)
}

func GetMetrics() []models.Metric {
    return Metrics
}
