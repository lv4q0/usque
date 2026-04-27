// Package config handles configuration loading and validation for usque.
// It supports loading configuration from environment variables and config files.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds all configuration for the usque application.
type Config struct {
	// WireGuard settings
	PrivateKey string
	PublicKey  string
	Endpoint   string
	PeerPubKey string

	// WARP / Cloudflare settings
	DeviceID   string
	AccountID  string
	AccessToken string
	LicenseKey  string

	// Tunnel settings
	TunnelAddress  string
	TunnelAddress6 string
	DNS            string
	MTU            int

	// Proxy settings
	ListenAddr string
	ListenPort int

	// Reconnect settings
	ReconnectInterval time.Duration
	MaxRetries        int

	// Logging
	LogLevel string
	LogJSON  bool
}

// DefaultConfig returns a Config populated with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		Endpoint:          "engage.cloudflareclient.com:2408",
		DNS:               "1.1.1.1",
		MTU:               1280,
		ListenAddr:        "127.0.0.1",
		ListenPort:        1080,
		ReconnectInterval: 5 * time.Second,
		MaxRetries:        10,
		LogLevel:          "info",
		LogJSON:           false,
	}
}

// LoadFromEnv populates Config fields from environment variables,
// overriding any previously set values.
func (c *Config) LoadFromEnv() error {
	if v := os.Getenv("USQUE_PRIVATE_KEY"); v != "" {
		c.PrivateKey = v
	}
	if v := os.Getenv("USQUE_PUBLIC_KEY"); v != "" {
		c.PublicKey = v
	}
	if v := os.Getenv("USQUE_ENDPOINT"); v != "" {
		c.Endpoint = v
	}
	if v := os.Getenv("USQUE_PEER_PUB_KEY"); v != "" {
		c.PeerPubKey = v
	}
	if v := os.Getenv("USQUE_DEVICE_ID"); v != "" {
		c.DeviceID = v
	}
	if v := os.Getenv("USQUE_ACCOUNT_ID"); v != "" {
		c.AccountID = v
	}
	if v := os.Getenv("USQUE_ACCESS_TOKEN"); v != "" {
		c.AccessToken = v
	}
	if v := os.Getenv("USQUE_LICENSE_KEY"); v != "" {
		c.LicenseKey = v
	}
	if v := os.Getenv("USQUE_TUNNEL_ADDRESS"); v != "" {
		c.TunnelAddress = v
	}
	if v := os.Getenv("USQUE_TUNNEL_ADDRESS6"); v != "" {
		c.TunnelAddress6 = v
	}
	if v := os.Getenv("USQUE_DNS"); v != "" {
		c.DNS = v
	}
	if v := os.Getenv("USQUE_MTU"); v != "" {
		mtu, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid USQUE_MTU value %q: %w", v, err)
		}
		c.MTU = mtu
	}
	if v := os.Getenv("USQUE_LISTEN_ADDR"); v != "" {
		c.ListenAddr = v
	}
	if v := os.Getenv("USQUE_LISTEN_PORT"); v != "" {
		port, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid USQUE_LISTEN_PORT value %q: %w", v, err)
		}
		c.ListenPort = port
	}
	if v := os.Getenv("USQUE_LOG_LEVEL"); v != "" {
		c.LogLevel = v
	}
	if v := os.Getenv("USQUE_LOG_JSON"); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("invalid USQUE_LOG_JSON value %q: %w", v, err)
		}
		c.LogJSON = b
	}
	return nil
}

// Validate checks that required fields are set and values are within
// acceptable ranges. Returns an error describing the first violation found.
func (c *Config) Validate() error {
	if c.PrivateKey == "" {
		return fmt.Errorf("private key must be set")
	}
	if c.PeerPubKey == "" {
		return fmt.Errorf("peer public key must be set")
	}
	if c.Endpoint == "" {
		return fmt.Errorf("endpoint must be set")
	}
	if c.MTU < 576 || c.MTU > 9000 {
		return fmt.Errorf("MTU %d is out of valid range [576, 9000]", c.MTU)
	}
	if c.ListenPort < 1 || c.ListenPort > 65535 {
		return fmt.Errorf("listen port %d is out of valid range [1, 65535]", c.ListenPort)
	}
	return nil
}
