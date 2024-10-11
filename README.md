# AirBox Agent

The AirBox Agent is a lightweight server monitoring tool written in Go. It is designed to collect server metrics such as CPU usage, RAM usage, cache usage, and storage size every minute. These metrics are saved as JSON files in the `/var/log/airbox` directory, with each file named based on the timestamp of collection.

## Features

- **Cross-Platform Compatibility**: The agent works on Linux, Windows, and macOS systems.
- **Metrics Collected**:
  - CPU usage (% utilization)
  - RAM usage (% utilization)
  - Cache usage (bytes)
  - Storage size (total available storage in bytes)
- **Data Logging**: The metrics are logged to JSON files, which makes it easy to integrate with other tools or automate analysis.

## Requirements

- **Go**: This agent is written in Go and requires Go to be installed for compilation.
- **Permissions**: The agent needs permission to write to `/var/log/airbox`. If you're running it on macOS or Windows, consider changing the log directory to a path that has write access, such as your home directory.
- **Libraries**: The agent uses `gopsutil` to gather system metrics. You can install the dependencies with:

  ```sh
  go get github.com/shirou/gopsutil/cpu
  go get github.com/shirou/gopsutil/mem
  go get github.com/shirou/gopsutil/disk
  ```

## Installation

1. **Clone the Repository**:

   ```sh
   git clone https://github.com/airbox-build/agent.git
   cd agent
   ```

2. **Build the Agent**:

   ```sh
   go build -o airbox-agent
   ```

3. **Run the Agent**:

   ```sh
   sudo ./airbox-agent
   ```

   > Note: Running the agent with `sudo` is required to write to `/var/log/airbox`.

## Usage

The agent runs continuously, collecting and logging metrics every minute. Each log file is saved in `/var/log/airbox` and follows the naming convention `airbox-<timestamp>.json`, where `<timestamp>` is the Unix timestamp of when the metrics were collected.

### Example Log File

```json
{
  "timestamp": "2024-10-11T09:53:07+08:00",
  "cpu": {
    "usage": [
      22.10517039079614
    ],
    "cores": 8
  },
  "memory": {
    "total": 17179869184,
    "used": 13534248960,
    "used_percent": 78.77969741821289,
    "swap_total": 11811160064,
    "swap_used": 10768678912
  },
  "storage": {
    "total": 994662584320,
    "used": 787438243840,
    "free": 207224340480,
    "cache": 787438243840
  },
  "system": {
    "hostname": "Nasruls-MacBook-Pro-2.local",
    "os": "darwin",
    "platform": "darwin",
    "platform_version": "14.6.1",
    "kernel_version": "23.6.0",
    "uptime": 2949679
  }
}
```

## Customization

- **Log Directory**: You can change the log directory by modifying the `dir` variable in the `saveMetricsToFile` function.
- **Collection Frequency**: The default collection frequency is set to 1 minute. You can modify the `time.Sleep` value in the `main` function to adjust this interval.

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the functionality or add new features.

## Security Vulnerabilities

If you discover a security vulnerability within AirBox, please send an e-mail to Nasrul Hazim via [nasrulhazim.m@gmail.com](mailto:nasrulhazim.m@gmail.com). All security vulnerabilities will be promptly addressed.

## Contributors

<a href="https://github.com/airbox-build/agent/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=airbox-build/agent"  alt="AirBox Contributors"/>
</a>
