package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type Metrics struct {
	Type      string        `json:"type"`
	Data      MetricsData   `json:"data"`
}

type MetricsData struct {
	Timestamp string        `json:"timestamp"`
	CPU       CPUMetrics     `json:"cpu"`
	Memory    MemoryMetrics  `json:"memory"`
	Storage   StorageMetrics `json:"storage"`
	System    SystemMetrics  `json:"system"`
	Meta      MetaMetrics    `json:"meta"`
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
	Hostname        string `json:"hostname"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platform_version"`
	KernelVersion   string `json:"kernel_version"`
	Uptime          uint64 `json:"uptime"`
	IPAddress       string `json:"ip_address"`
}

type MetaMetrics struct {
	FilePath     string `json:"file_path"`
	Interval     int    `json:"interval"`
	FileCreation string `json:"file_creation"`
	User         string `json:"user"`
}

func main() {
	// Define command line flags
	logPath := flag.String("logpath", "./logs", "Directory to store the log files")
	interval := flag.Int("interval", 60, "Interval to collect metrics in seconds")
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		collectAndStoreMetrics(*logPath, *interval)
	}
}

func collectAndStoreMetrics(logPath string, interval int) {
	var wg sync.WaitGroup
	metricsData := &MetricsData{Timestamp: time.Now().Format(time.RFC3339)}

	// Collect CPU Metrics
	wg.Add(1)
	go func() {
		defer wg.Done()
		metricsData.CPU = collectCPUMetrics()
	}()

	// Collect Memory Metrics
	wg.Add(1)
	go func() {
		defer wg.Done()
		metricsData.Memory = collectMemoryMetrics()
	}()

	// Collect Storage Metrics
	wg.Add(1)
	go func() {
		defer wg.Done()
		metricsData.Storage = collectStorageMetrics()
	}()

	// Collect System Metrics
	wg.Add(1)
	go func() {
		defer wg.Done()
		metricsData.System = collectSystemMetrics()
	}()

	// Collect Meta Metrics
	metricsData.Meta = collectMetaMetrics(logPath, interval)

	wg.Wait()

	metrics := Metrics{
		Type: "metrics",
		Data: *metricsData,
	}

	saveMetricsToFile(metrics, logPath)
}

func collectCPUMetrics() CPUMetrics {
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("Error getting CPU usage: %v", err)
		return CPUMetrics{}
	}

	return CPUMetrics{
		Usage: cpuUsage,
		Cores: runtime.NumCPU(),
	}
}

func collectMemoryMetrics() MemoryMetrics {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error getting memory usage: %v", err)
		return MemoryMetrics{}
	}

	swapInfo, err := mem.SwapMemory()
	if err != nil {
		log.Printf("Error getting swap memory usage: %v", err)
		return MemoryMetrics{}
	}

	return MemoryMetrics{
		Total:       memInfo.Total,
		Used:        memInfo.Used,
		UsedPercent: memInfo.UsedPercent,
		SwapTotal:   swapInfo.Total,
		SwapUsed:    swapInfo.Used,
	}
}

func collectStorageMetrics() StorageMetrics {
	cacheInfo, err := disk.Usage("/")
	if err != nil {
		log.Printf("Error getting storage usage: %v", err)
		return StorageMetrics{}
	}

	return StorageMetrics{
		Total: cacheInfo.Total,
		Used:  cacheInfo.Used,
		Free:  cacheInfo.Free,
		Cache: cacheInfo.Used, // Assuming Cache refers to Used here
	}
}

func collectSystemMetrics() SystemMetrics {
	hostInfo, err := host.Info()
	if err != nil {
		log.Printf("Error getting host information: %v", err)
		return SystemMetrics{}
	}

	ipAddress := getMachineIPAddress()

	return SystemMetrics{
		Hostname:        hostInfo.Hostname,
		OS:              hostInfo.OS,
		Platform:        hostInfo.Platform,
		PlatformVersion: hostInfo.PlatformVersion,
		KernelVersion:   hostInfo.KernelVersion,
		Uptime:          hostInfo.Uptime,
		IPAddress:       ipAddress,
	}
}

func collectMetaMetrics(logPath string, interval int) MetaMetrics {
	timestamp := time.Now()
	return MetaMetrics{
		FilePath:     filepath.Join(logPath, fmt.Sprintf("%d.json", timestamp.Unix())),
		Interval:     interval,
		FileCreation: timestamp.Format(time.RFC3339),
		User:         os.Getenv("USER"),
	}
}

func getMachineIPAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Error getting network interfaces: %v", err)
		return ""
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					return ipNet.IP.String()
				}
			}
		}
	}
	return ""
}

func saveMetricsToFile(metrics Metrics, logPath string) {
	if err := os.MkdirAll(logPath, 0755); err != nil {
		log.Printf("Error creating directory: %v", err)
		return
	}

	filePath := metrics.Data.Meta.FilePath
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(metrics); err != nil {
		log.Printf("Error writing metrics to file: %v", err)
	}
}
