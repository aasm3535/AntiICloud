# AntiICloud

AntiICloud is a Windows command-line utility designed to help authorized users diagnose and manage iCloud activation lock workflows on Apple devices.

## Features

- USB device connection check
- Activation lock status diagnostics
- Simulated bypass procedure with real-time progress
- Step-by-step console feedback
- Safe, offline operation — no network traffic

## Requirements

- Windows 10/11
- Apple device connected via USB
- No additional drivers required

## Usage

1. Download `AntiICloud.exe` from the [Releases](../../releases) page.
2. Open a terminal in the download folder.
3. Run the program and follow the on-screen instructions:

```powershell
.\AntiICloud.exe
```

## Building from source

```bash
git clone https://github.com/aasm3535/AntiICloud.git
cd AntiICloud
go build -ldflags "-s -w" -o AntiICloud.exe .
```

## Disclaimer

This tool is intended for use only on devices you own or have explicit permission to manage. The author is not responsible for any misuse.
