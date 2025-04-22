![Go](https://img.shields.io/badge/Go-1.19+-blue.svg)\
![License](https://img.shields.io/badge/License-MIT-green.svg)

A robust Go utility for cleaning up old files in specified directories with configurable age thresholds and logging capabilities.

Features
--------

-   🗑️ Delete files older than N days

-   📝 Comprehensive logging (console and/or file)

-   🧪 Dry-run mode for testing

-   🛠️ Configurable via command-line arguments

-   📊 Helpful usage documentation

Installation
------------

### Prerequisites

-   Go 1.19 or higher

-   Linux/Unix system (tested on Rocky Linux)

### Installation Methods

#### 1\. From Source (Recommended)

bash

Copy

Download

# Clone the repository
git clone https://github.com/rouve/go-cleanup.git
cd file-cleanup

# Build the binary
go build -o cleanup

# Install system-wide (optional)
sudo mv cleanup /usr/local/bin/

#### 2\. Using go install

bash

Copy

Download

go install github.com/rouve/go-cleanup@latest

#### 3\. Download Pre-built Binary

Check the [Releases](https://github.com/rouve/go-cleanup/releases) page for pre-built binaries.

Usage
-----

bash

Copy

Download

cleanup --dir=/path/to/directory [options]

### Basic Examples

bash

Copy

Download

# Clean files older than 7 days (default)
cleanup --dir=/tmp

# Clean files older than 30 days
cleanup --dir=/var/log --days=30

# Dry run (test without deleting)
cleanup --dir=/backups --days=90 --dry-run

# With logging to file
cleanup --dir=/data --days=14 --log=/var/log/cleanup.log

### All Options

| Option | Description | Default |
| --- | --- | --- |
| `--dir` | Directory to clean (required) | - |
| `--days` | Delete files older than N days | 7 |
| `--dry-run` | Simulate without deleting | false |
| `--log` | Path to log file | stdout |
| `--help` | Show help message | - |

Cron Job Setup
--------------

Example cron entry to run daily at 2 AM:

bash

Copy

Download

0 2 * * * /usr/local/bin/cleanup --dir=/tmp --days=7 --log=/var/log/cleanup.log

For Webmin:

1.  Go to "System" → "Scheduled Cron Jobs"

2.  Add new command: `/usr/local/bin/cleanup --dir=/path --days=N`

3.  Set desired schedule

Logging
-------

Logs include timestamps and operation details. Example log output:

Copy

Download

2023-07-20 14:30:45 - Starting cleanup in /tmp for files older than 7 days (2023-07-13 14:30:45)
2023-07-20 14:30:45 - Found candidate: tempfile.txt (modified: 2023-07-10 08:15:22)
2023-07-20 14:30:45 - Deleted: tempfile.txt
2023-07-20 14:30:45 - Cleanup complete. 1 files deleted (dry-run: false)

License
-------

MIT License. See [LICENSE](https://license/) file for details.

Contributing
------------

Pull requests are welcome! Please ensure:

1.  Go code is properly formatted (`go fmt`)

2.  New features include tests

3.  Documentation is updated
