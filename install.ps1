# Check if running with administrative privileges
$isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)

if (-not $isAdmin) {
    Write-Host "Please run this script as an administrator."
    exit
}

$GH_REPO = "airbox-build/agent"
$TIMEOUT = 90
$LOG_PATH = "C:\ProgramData\AirBox\Logs"
$INTERVAL = 60

# Parse command line arguments
param(
    [string]$logpath = $LOG_PATH,
    [int]$interval = $INTERVAL
)

$LOG_PATH = $logpath
$INTERVAL = $interval

# Get the current logged-in user
$USERNAME = $env:USERNAME

$VERSION = (Invoke-WebRequest -Uri "https://api.github.com/repos/$GH_REPO/releases/latest" -UseBasicParsing).Content | ConvertFrom-Json | Select-Object -ExpandProperty tag_name
if (-not $VERSION) {
    Write-Host "`nThere was an error trying to check what is the latest version of airbox.`nPlease try again later.`n"
    exit 1
}

$OS_type = $env:PROCESSOR_ARCHITECTURE
switch ($OS_type) {
    "AMD64", "x86_64" {
        $OS_type = "amd64"
    }
    "x86", "i386" {
        $OS_type = "386"
    }
    "ARM64" {
        $OS_type = "arm64"
    }
    default {
        Write-Host "OS type not supported"
        exit 2
    }
}

$GH_REPO_BIN = "agent${VERSION}-windows-${OS_type}.tar.gz"

# Create tmp directory
$TMP_DIR = New-TemporaryFile -Directory | Select-Object -ExpandProperty FullName
Write-Host "Change to temporary directory $TMP_DIR"
Set-Location $TMP_DIR

Write-Host "Downloading AirBox Agent $VERSION"
$LINK = "https://github.com/$GH_REPO/releases/download/$VERSION/$GH_REPO_BIN"
Write-Host "Downloading $LINK"

Invoke-WebRequest -Uri $LINK -OutFile "$TMP_DIR\$GH_REPO_BIN"
if ($?) {
    Write-Host "Error downloading"
    exit 2
}

$BINARY_PATH = "C:\Program Files\AirBox"
$null = New-Item -Path $BINARY_PATH -ItemType Directory -Force

Copy-Item -Path "$TMP_DIR\airbox.exe" -Destination $BINARY_PATH -Force
if ($?) {
    exit 2
}

$BINARY_DIRECTORY = "C:\ProgramData\AirBox"
$null = New-Item -Path $BINARY_DIRECTORY -ItemType Directory -Force

Remove-Item -Path $TMP_DIR -Recurse -Force
Write-Host "Installed successfully to $BINARY_PATH\airbox.exe"

# Create the service
$SERVICE_NAME = "AirboxAgent"
$SERVICE_PATH = "$BINARY_PATH\airbox.exe"
$SERVICE_LOG_PATH = "$LOG_PATH"
$SERVICE_INTERVAL = $INTERVAL

$serviceParams = @{
    Name             = $SERVICE_NAME
    BinaryPathName   = "$SERVICE_PATH --logpath=$SERVICE_LOG_PATH --interval=$SERVICE_INTERVAL"
    DisplayName      = $SERVICE_NAME
    Description      = "AirBox Agent Service"
    StartupType      = "Automatic"
    Credential       = "LocalSystem"
    DependsOn        = @("tcpip")
    ErrorControl     = "Normal"
    ServiceArguments = @()
}

$service = New-Service @serviceParams

if ($service) {
    Write-Host "Service created successfully."
} else {
    Write-Host "Failed to create the service."
}
