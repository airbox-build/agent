[![Build](https://github.com/airbox-build/agent/actions/workflows/release.yml/badge.svg)](https://github.com/airbox-build/agent/actions/workflows/release.yml) [![Unit Test](https://github.com/airbox-build/agent/actions/workflows/unit-test.yml/badge.svg)](https://github.com/airbox-build/agent/actions/workflows/unit-test.yml)

# AirBox Agent

The AirBox Agent is a lightweight server monitoring tool written in Go. It is designed to collect server metrics such as CPU usage, RAM usage, cache usage, and storage size at a configurable interval. These metrics are saved as JSON files in a specified directory, with each file named based on the timestamp of collection.

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
- **Permissions**: The agent needs permission to write to the log directory. By default, it writes to `/tmp/airbox`, which should be accessible on most systems.
- **Libraries**: The agent uses `gopsutil` to gather system metrics. You can install the dependencies with:

  ```sh
  go get github.com/shirou/gopsutil/cpu
  go get github.com/shirou/gopsutil/mem
  go get github.com/shirou/gopsutil/disk
  go get github.com/shirou/gopsutil/host
  ```

## Installation

### Install via bash script (Linux & Mac)

Linux & Mac users can install it directly to `/usr/local/bin/airbox` with:

```bash
sudo bash -c "$(curl -sL https://raw.githubusercontent.com/airbox-build/agent/refs/heads/main/install)"
```

### Download static binary (Windows, Linux and Mac)

Run the following command which will download latest version and configure default configuration for Windows.

```batch
powershell -command "(New-Object Net.WebClient).DownloadFile('https://raw.githubusercontent.com/airbox-build/agent/main/install.ps1', '%TEMP%\install.ps1') && %TEMP%\install.ps1 && del %TEMP%\install.ps1"
```

## Usage

The agent runs continuously, collecting and logging metrics at a configurable interval. The default interval is 60 seconds, and the default log directory is `/tmp/airbox` for Windows is `C:\ProgramData\AirBox\Logs`. You can customize these values using command-line flags.

### Command-Line Flags

- `--logpath`: Specifies the directory to store the log files (default is `/tmp/airbox`).
- `--interval`: Specifies the interval to collect metrics in seconds (default is `60`).

### Example Log File

```json
{
  "type": "metrics",
  "data": {
    "timestamp": "2024-10-13T07:02:56+08:00",
    "cpu": {
      "usage": [
        38.642078793807144
      ],
      "cores": 8
    },
    "memory": {
      "total": 17179869184,
      "used": 13063290880,
      "used_percent": 76.03836059570312,
      "swap_total": 4294967296,
      "swap_used": 2719744000
    },
    "storage": {
      "total": 994662584320,
      "used": 776870703104,
      "free": 217791881216,
      "cache": 776870703104
    },
    "system": {
      "hostname": "Nasruls-MacBook-Pro-2.local",
      "os": "darwin",
      "platform": "darwin",
      "platform_version": "14.6.1",
      "kernel_version": "23.6.0",
      "uptime": 3112268,
      "ip_address": "192.168.0.39"
    },
    "meta": {
      "file_path": "logs/1728774176.json",
      "interval": 3,
      "file_creation": "2024-10-13T07:02:56+08:00",
      "user": "nasrulhazim"
    }
  }
}
```

## Customization

- **Log Directory**: You can change the log directory by using the `--logpath` flag when running the agent.
- **Collection Frequency**: The default collection frequency is 60 seconds. You can change this by using the `--interval` flag.

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
