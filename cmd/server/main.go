package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"fioincline/internal/server"
)

func main() {
	configPath := findConfig()
	cfg := server.ReadConfig(configPath)
	logPath := filepath.Join(filepath.Dir(configPath), "server.log")

	h := server.NewServer(cfg, logPath)
	addr := fmt.Sprintf("%s:%d", cfg.Address, cfg.Port)

	msg := fmt.Sprintf("SOAP server running at http://%s:%d/soap", cfg.Address, cfg.Port)
	fmt.Println(msg)
	h.Log().Log("", msg)

	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatal(err)
	}
}

// findConfig ищет config.ini: рядом с бинарником или в текущем каталоге.
func findConfig() string {
	if exe, err := os.Executable(); err == nil {
		p := filepath.Join(filepath.Dir(exe), "config.ini")
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return "config.ini"
}