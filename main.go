package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/diniboy1123/usque/internal/config"
	"github.com/diniboy1123/usque/internal/proxy"
)

const (
	appName    = "usque"
	appVersion = "0.1.0"
)

func main() {
	var (
		configFile  = flag.String("config", "config.json", "Path to configuration file")
		showVersion = flag.Bool("version", false, "Print version information and exit")
		verbose     = flag.Bool("verbose", false, "Enable verbose logging")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", appName)
		fmt.Fprintf(os.Stderr, "A WireGuard to MASQUE/HTTP3 proxy tunnel.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *showVersion {
		fmt.Printf("%s version %s\n", appName, appVersion)
		os.Exit(0)
	}

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Verbose logging enabled")
	}

	// Load configuration from file or environment variables
	cfg, err := config.Load(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if *verbose {
		log.Printf("Configuration loaded from %s", *configFile)
		log.Printf("Endpoint: %s", cfg.Endpoint)
		log.Printf("Listen address: %s", cfg.ListenAddr)
	}

	// Initialize and start the proxy
	p, err := proxy.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize proxy: %v", err)
	}

	log.Printf("%s v%s starting...", appName, appVersion)
	log.Printf("Listening on %s", cfg.ListenAddr)

	if err := p.Run(); err != nil {
		log.Fatalf("Proxy exited with error: %v", err)
	}
}
