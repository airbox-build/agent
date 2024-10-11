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
sudo bash < <(curl -sL https://raw.githubusercontent.com/airbox-build/agent/main/install)
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
  "timestamp": "2024-10-11T11:30:18+08:00",
  "cpu": {
    "usage": [
      22.23625623468918
    ],
    "cores": 8
  },
  "memory": {
    "total": 17179869184,
    "used": 13222641664,
    "used_percent": 76.96590423583984,
    "swap_total": 11811160064,
    "swap_used": 10595598336
  },
  "storage": {
    "total": 994662584320,
    "used": 787459915776,
    "free": 207202668544,
    "cache": 787459915776
  },
  "system": {
    "hostname": "Nasruls-MacBook-Pro-2.local",
    "os": "darwin",
    "platform": "darwin",
    "platform_version": "14.6.1",
    "kernel_version": "23.6.0",
    "uptime": 2955510,
    "ip_address": "192.168.0.105"
  },
  "network": {
    "interfaces": [
      {
        "name": "lo0",
        "bytes_sent": 23583009561,
        "bytes_recv": 23583009561,
        "packets_sent": 44482621,
        "packets_recv": 44482621
      },
      {
        "name": "gif0",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "stf0",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "anpi0",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "anpi1",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "en3",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "en4",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "en1",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "en2",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "bridge0",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "ap1",
        "bytes_sent": 31367021,
        "bytes_recv": 0,
        "packets_sent": 116011,
        "packets_recv": 0
      },
      {
        "name": "en0",
        "bytes_sent": 22295660432,
        "bytes_recv": 81311602125,
        "packets_sent": 55187252,
        "packets_recv": 87588903
      },
      {
        "name": "awdl0",
        "bytes_sent": 5565058,
        "bytes_recv": 31673,
        "packets_sent": 50420,
        "packets_recv": 97
      },
      {
        "name": "llw0",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "utun0",
        "bytes_sent": 80,
        "bytes_recv": 0,
        "packets_sent": 1,
        "packets_recv": 0
      },
      {
        "name": "utun1",
        "bytes_sent": 144244,
        "bytes_recv": 0,
        "packets_sent": 742,
        "packets_recv": 0
      },
      {
        "name": "utun2",
        "bytes_sent": 144304,
        "bytes_recv": 0,
        "packets_sent": 743,
        "packets_recv": 0
      },
      {
        "name": "utun3",
        "bytes_sent": 144304,
        "bytes_recv": 0,
        "packets_sent": 743,
        "packets_recv": 0
      },
      {
        "name": "en10",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      },
      {
        "name": "en9",
        "bytes_sent": 0,
        "bytes_recv": 0,
        "packets_sent": 0,
        "packets_recv": 0
      }
    ]
  },
  "meta": {
    "file_path": "/tmp/airbox/airbox-1728617418.json",
    "interval": 60,
    "file_creation": "2024-10-11T11:30:18+08:00",
    "user": "nasrulhazim"
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
