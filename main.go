package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)

type Metrics struct {
	Timestamp   string  `json:"timestamp"`
	CPUUsage    float64 `json:"cpu_usage"`
	RAMUsage    float64 `json:"ram_usage"`
	CacheUsage  uint64  `json:"cache_usage"`
	StorageSize uint64  `json:"storage_size"`
}

func main() {
	for {
		collectAndStoreMetrics()
		time.Sleep(10 * time.Second)
	}
}

func collectAndStoreMetrics() {
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Printf("Error getting CPU usage: %v\n", err)
		return
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error getting memory usage: %v\n", err)
		return
	}

	cacheInfo, err := disk.Usage("/")
	if err != nil {
		fmt.Printf("Error getting cache usage: %v\n", err)
		return
	}

	storageSize := cacheInfo.Total

	metrics := Metrics{
		Timestamp:   time.Now().Format(time.RFC3339),
		CPUUsage:    cpuUsage[0],
		RAMUsage:    memInfo.UsedPercent,
		CacheUsage:  cacheInfo.Used,
		StorageSize: storageSize,
	}

	saveMetricsToFile(metrics)
}

func saveMetricsToFile(metrics Metrics) {
	dir := "/var/log/airbox"
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	filename := fmt.Sprintf("airbox-%d.json", time.Now().Unix())
	filePath := filepath.Join(dir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(metrics); err != nil {
		fmt.Printf("Error writing metrics to file: %v\n", err)
	}
}
