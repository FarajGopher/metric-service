package handlers

import (
	"encoding/json"
	"metric_service/models"
	"metric_service/storage"
	constant "metric_service/utils"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := storage.GetMetrics()

	log.WithFields(log.Fields{
		"count": len(metrics),
		"route": "/metrics",
		"method": r.Method,
	}).Info("Fetched metrics")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		log.WithError(err).Error("Failed to encode metrics response")
		http.Error(w, constant.MsgFailedToSendResponse, http.StatusInternalServerError)
	}
}

func PostMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Warnf("Invalid method %s on /metrics/add", r.Method)
		http.Error(w, constant.MsgOnlyPOSTAllowed, http.StatusMethodNotAllowed)
		return
	}

	var metric models.Metric
	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		log.WithError(err).Warn("Failed to decode metric payload")
		http.Error(w, constant.MsgInvalidMetricData, http.StatusBadRequest)
		return
	}

	storage.AddMetric(metric)

	log.WithFields(log.Fields{
		"node":     metric.NodeName,
		"cpuUsage": metric.CPUUsage,
		"memUsage": metric.MemUsage,
		"ts":       metric.Timestamp,
	}).Info("Stored new metric")

	response := map[string]interface{}{
		"message": constant.MsgMetricStoredSuccess,
		"metric":  metric,
	}

	w.Header().Set("Content-Type", constant.ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.WithError(err).Error("Failed to send metric stored response")
		http.Error(w, constant.MsgFailedToSendResponse, http.StatusInternalServerError)
	}
}
