package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCollectAndStoreMetrics(t *testing.T) {
	tempDir := filepath.Join(os.TempDir(), "airbox_test")
	defer os.RemoveAll(tempDir)

	interval := 30
	collectAndStoreMetrics(tempDir, interval)

	// Check if the file is created
	files, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Failed to read the directory: %v", err)
	}

	if len(files) == 0 {
		t.Fatalf("Expected metrics file, but none was created")
	}

	filePath := filepath.Join(tempDir, files[0].Name())
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Metrics file was not created at expected location: %v", filePath)
	}
}

func TestSaveMetricsToFile(t *testing.T) {
	tempDir := filepath.Join(os.TempDir(), "airbox_test_save")
	defer os.RemoveAll(tempDir)

	metrics := Metrics{
		Type: "metrics",
		Data: MetricsData{
			Timestamp: time.Now().Format(time.RFC3339),
			CPU: CPUMetrics{
				Usage: []float64{0.0},
				Cores: 1,
			},
			Memory: MemoryMetrics{
				Total:       1024,
				Used:        512,
				UsedPercent: 50.0,
				SwapTotal:   2048,
				SwapUsed:    1024,
			},
			Storage: StorageMetrics{
				Total: 10000,
				Used:  5000,
				Free:  5000,
				Cache: 2000,
			},
			System: SystemMetrics{
				Hostname:        "localhost",
				OS:              "linux",
				Platform:        "ubuntu",
				PlatformVersion: "20.04",
				KernelVersion:   "5.4.0",
				Uptime:          1000,
				IPAddress:       "192.168.1.1",
			},
			Meta: MetaMetrics{
				FilePath:     filepath.Join(tempDir, "test_metrics.json"),
				Interval:     30,
				FileCreation: time.Now().Format(time.RFC3339),
				User:         "test_user",
			},
		},
	}

	saveMetricsToFile(metrics, tempDir)

	// Check if the file is created
	if _, err := os.Stat(metrics.Data.Meta.FilePath); os.IsNotExist(err) {
		t.Errorf("Metrics file was not created at expected location: %v", metrics.Data.Meta.FilePath)
	}
}

func TestFlagsDefaultValues(t *testing.T) {
	logPath := "./logs"
	interval := 60

	if logPath != "./logs" {
		t.Errorf("Expected default logPath to be ./logs, but got %v", logPath)
	}

	if interval != 60 {
		t.Errorf("Expected default interval to be 60, but got %v", interval)
	}
}
