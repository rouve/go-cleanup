package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const helpMessage = `File Cleanup Utility (v1.0)

Usage:
  cleanup [options]

Options:
  --dir string     Directory path to clean (required)
  --days int       Delete files older than this many days (default 7)
  --dry-run        Simulate cleanup without actually deleting files
  --log string     Path to log file (default stdout)
  --help           Show this help message

Examples:
  cleanup --dir=/tmp --days=30
  cleanup --dir=/var/log --days=7 --log=/var/log/cleanup.log
  cleanup --dir=/backups --days=90 --dry-run
`

func main() {
	// Set PATH in case running from cron
	os.Setenv("PATH", "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin")

	// Define command-line flags
	var (
		dirPath string
		daysOld int
		dryRun  bool
		logFile string
    	help bool
	)

	flag.StringVar(&dirPath, "dir", "", "Directory path to clean (required)")
	flag.IntVar(&daysOld, "days", 7, "Delete files older than this many days")
	flag.BoolVar(&dryRun, "dry-run", false, "Simulate cleanup without actually deleting files")
	flag.StringVar(&logFile, "log", "", "Path to log file (default: stdout only)")
    flag.BoolVar(&help, "help", false, "Show help message")
	flag.Parse()

	// Show help if requested
	if help {
		fmt.Print(helpMessage)
		os.Exit(0)
	}

	// Validate required arguments
	if dirPath == "" {
		fmt.Println("Error: --dir argument is required\n")
		fmt.Print(helpMessage)
		os.Exit(1)
	}

	// Set up logging
	var logWriter = os.Stdout
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening log file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		logWriter = f
	}

	log := func(format string, args ...interface{}) {
		msg := fmt.Sprintf(format, args...)
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(logWriter, "%s - %s\n", timestamp, msg)
	}

	// Calculate cutoff time
	cutoffTime := time.Now().AddDate(0, 0, -daysOld)
	log("Starting cleanup in %s for files older than %d days (%v)", dirPath, daysOld, cutoffTime)
	if dryRun {
		log("DRY RUN MODE - no files will actually be deleted")
	}

	// Read directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log("Error reading directory: %v", err)
		os.Exit(1)
	}

	// Process files
	deletedCount := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if file.ModTime().Before(cutoffTime) {
			filePath := filepath.Join(dirPath, file.Name())
			log("Found candidate: %s (modified: %v)", file.Name(), file.ModTime())

			if !dryRun {
				err := os.Remove(filePath)
				if err != nil {
					log("Failed to delete %s: %v", filePath, err)
				} else {
					log("Deleted: %s", file.Name())
					deletedCount++
				}
			}
		}
	}

	log("Cleanup complete. %d files deleted (dry-run: %v)", deletedCount, dryRun)
}
