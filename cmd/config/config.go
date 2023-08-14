package config

import (
	"os"
)

func ConfigSetup() {
	// Database settings
	os.Setenv("DB_USERNAME", "vgudza")
	os.Setenv("DB_PASSWORD", "vgudza")
	os.Setenv("DB_HOST", "***input your DB host here***")
	os.Setenv("DB_NAME", "vgudza_shop")

	os.Setenv("DB_POOL_MAXCONN", "5")
	os.Setenv("DB_POOL_MAXCONN_LIFETIME", "300")

	// NATS-Streaming settings
	os.Setenv("NATS_HOSTS", "***input your NATS host here***")
	os.Setenv("NATS_CLUSTER_ID", "world-nats-stage")
	os.Setenv("NATS_CLIENT_ID", "vgudza")
	os.Setenv("NATS_SUBJECT", "go.test-gudza")
	os.Setenv("NATS_DURABLE_NAME", "Replica-1")
	os.Setenv("NATS_ACK_WAIT_SECONDS", "30")

	// Cache settings
	os.Setenv("CACHE_SIZE", "10")
	os.Setenv("APP_KEY", "WB-1")
}
