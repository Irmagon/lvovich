package server

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Logger пишет строки вида "[YYYY-MM-DD HH:MM:SS.mmm] ip msg" в server.log.
type Logger struct {
	mu     sync.Mutex
	path   string
	lastIP string
}

// NewLogger создаёт логгер с указанным файлом (обычно server.log в корне репо).
func NewLogger(path string) *Logger {
	return &Logger{path: path}
}

// SetLastIP запоминает IP последнего запроса (аналог _lastIp в оригинале).
func (l *Logger) SetLastIP(ip string) {
	l.mu.Lock()
	l.lastIP = ip
	l.mu.Unlock()
}

// Log пишет строку с IP последнего запроса (или '-').
func (l *Logger) Log(ip, msg string) {
	l.mu.Lock()
	addr := ip
	if addr == "" {
		addr = l.lastIP
	}
	if addr == "" {
		addr = "-"
	}
	line := logLine(addr, msg)
	l.lastIP = ip
	l.mu.Unlock()
	l.write(line)
}

func logLine(addr, msg string) string {
	now := time.Now()
	ts := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%03d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1e6)
	return fmt.Sprintf("[%s] %s %s\n", ts, addr, msg)
}

func (l *Logger) write(line string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	f, err := os.OpenFile(l.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return
	}
	defer f.Close()
	_, _ = f.WriteString(line)
}
