// 代码生成时间: 2025-08-23 18:13:33
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// LogEntry represents a single log entry with a timestamp and a message.
type LogEntry struct {
    Timestamp time.Time
    Message   string
}

// parseLogEntry parses a line from a log file into a LogEntry struct.
func parseLogEntry(line string) (*LogEntry, error) {
    parts := strings.Fields(line)
    if len(parts) < 2 {
        return nil, fmt.Errorf("invalid log entry format")
    }

    timestamp, err := time.Parse("2006-01-02 15:04:05", parts[0]+ " " + parts[1])
    if err != nil {
        return nil, fmt.Errorf("failed to parse timestamp: %w", err)
    }

    message := strings.Join(parts[2:], " ")
    return &LogEntry{Timestamp: timestamp, Message: message}, nil
}

// parseLogFile reads and parses a log file, returning a slice of LogEntry structs.
func parseLogFile(filePath string) ([]LogEntry, error) {
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to read log file: %w", err)
    }

    lines := strings.Split(strings.TrimSpace(string(content)), "\
")
    var entries []LogEntry
    for _, line := range lines {
        entry, err := parseLogEntry(line)
        if err != nil {
            log.Printf("Skipping invalid log entry: %v", err)
            continue
        }
        entries = append(entries, *entry)
    }
    return entries, nil
}

// walkLogs recursively walks through directories and subdirectories to parse log files.
func walkLogs(rootPath string) ([]LogEntry, error) {
    var allEntries []LogEntry
    err := filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
