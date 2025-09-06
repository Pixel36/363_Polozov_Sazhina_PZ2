package main

import "fmt"

type LogEntry struct {
 IP        string
 HTTPCode  int
 Timestamp string
}

func filterErrorLogs(logs []LogEntry) []LogEntry {
 out := []LogEntry{}
 for _, l := range logs {
  if l.HTTPCode >= 400 && l.HTTPCode < 600 {
   out = append(out, l)
  }
 }
 return out
}

func main() {
 logs := []LogEntry{
  {IP: "1.1.1.1", HTTPCode: 200, Timestamp: "2025-09-01T10:00:00Z"},
  {IP: "2.2.2.2", HTTPCode: 404, Timestamp: "2025-09-01T11:00:00Z"},
  {IP: "3.3.3.3", HTTPCode: 500, Timestamp: "2025-09-01T12:00:00Z"},
 }
 errLogs := filterErrorLogs(logs)
 fmt.Println("Error logs:", errLogs)
}
