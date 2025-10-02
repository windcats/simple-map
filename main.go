package main

import (
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	Version     string
	BuildTime   string
	Commit      string
	configPath  = flag.String("c", "config.yaml", "Path to the configuration file")
	versionFile = flag.Bool("v", false, "the app version")
)

func main() {
	flag.Parse()
	if *versionFile {
		fmt.Printf("Version: %s\nBuildTime: %s\nCommit: %s\n", Version, BuildTime, Commit)
		return
	}

	cfg := LoadConfig(*configPath)

	provider := newTileProvider(cfg.Sources)

	handler := NewTileHandler(provider)

	go Serve(cfg.Port, handler)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	// 这里可以添加优雅关闭逻辑
}
