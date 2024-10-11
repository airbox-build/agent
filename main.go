package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

type Metrics struct {
	Timestamp string        `json:"timestamp"`
	CPU       CPUMetrics     `json:"cpu"`
	Memory    MemoryMetrics  `json:"memory"`
	Storage   StorageMetrics `json:"storage"`
	System    SystemMetrics  `json:"system"`
}

type CPUMetrics struct {
	Usage []float64 `json:"usage"`
	Cores int       `json:"cores"`
}

type MemoryMetrics struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	SwapTotal   uint64  `json:"swap_total"`
	SwapUsed    uint64  `json:"swap_used"`
}

type StorageMetrics struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
	Cache uint64 `json:"cache"`
}

type SystemMetrics struct {
	Hostname   string `json:"hostname"`
	OS         string `json:"os"`
	Platform   string `json:"platform"`
	PlatformVersion string `json:"platform_version"`
	KernelVersion string `json:"kernel_version"`
	Uptime     uint64 `json:"uptime"`
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

	cpuCores := runtime.NumCPU()

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error getting memory usage: %v\n", err)
		return
	}

	swapInfo, err := mem.SwapMemory()
	if err != nil {
		fmt.Printf("Error getting swap memory usage: %v\n", err)
		return
	}

	cacheInfo, err := disk.Usage("/")
	if err != nil {
		fmt.Printf("Error getting cache usage: %v\n", err)
		return
	}

	hostInfo, err := host.Info()
	if err != nil {
		fmt.Printf("Error getting host information: %v\n", err)
		return
	}

	metrics := Metrics{
		Timestamp: time.Now().Format(time.RFC3339),
		CPU: CPUMetrics{
			Usage: cpuUsage,
			Cores: cpuCores,
		},
		Memory: MemoryMetrics{
			Total:       memInfo.Total,
			Used:        memInfo.Used,
			UsedPercent: memInfo.UsedPercent,
			SwapTotal:   swapInfo.Total,
			SwapUsed:    swapInfo.Used,
		},
		Storage: StorageMetrics{
			Total: cacheInfo.Total,
			Used:  cacheInfo.Used,
			Free:  cacheInfo.Free,
			Cache: cacheInfo.Used,
		},
		System: SystemMetrics{
			Hostname:       hostInfo.Hostname,
			OS:             hostInfo.OS,
			Platform:       hostInfo.Platform,
			PlatformVersion: hostInfo.PlatformVersion,
			KernelVersion:  hostInfo.KernelVersion,
			Uptime:         hostInfo.Uptime,
		},
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
