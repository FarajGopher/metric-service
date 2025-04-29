package main

import (
	"math/rand"
	"metric_service/auth"
	"metric_service/config"
	"metric_service/handlers"
	"metric_service/models"
	"metric_service/storage"
	"metric_service/utils" 
	"net/http"
	"time"
)

func main() {
	utils.InitLogger()               // Initialize logger
	config.LoadConfig()              // Load config using Viper
	utils.Log.Info("Config loaded")  // Log config loading

	go mockMetricsCollector()

	http.HandleFunc("/metrics", auth.BasicAuthMiddleware(handlers.GetMetrics))
	http.HandleFunc("/metrics/add", auth.BasicAuthMiddleware(handlers.PostMetric))

	port := config.AppConfig.Server.Port
	utils.Log.Infof("Server running on port :%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		utils.Log.Fatalf("Failed to start server: %v", err)
	}
}

func mockMetricsCollector() {
	nodeNames := []string{"node1", "node2", "node3"}
	for {
		for _, node := range nodeNames {
			metric := models.Metric{
				NodeName:  node,
				CPUUsage:  rand.Float64() * 100,
				MemUsage:  rand.Float64() * 100,
				Timestamp: time.Now().Unix(),
			}
			storage.AddMetric(metric)

			utils.Log.WithFields(map[string]interface{}{
				"node": node,
				"cpu":  metric.CPUUsage,
				"mem":  metric.MemUsage,
			}).Info("Mock metric added")
		}
		time.Sleep(10 * time.Second)
	}
}
