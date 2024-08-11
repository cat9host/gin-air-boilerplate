package config

import (
	"os"
	"testing"
)

func TestConfigure(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("METRICS_PORT", "10000")
	os.Setenv("HC_PORT", "10001")
	os.Setenv("SECRET_KEY", "kawabanga")
	os.Setenv("MYSQL_DSN", "splinter:turtle@(192.168.0.29:3333)/data?charset=utf8mb4&parseTime=True&loc=Local&tls=skip-verify")

	Configure()

	if AppPort != "8080" {
		t.Errorf("Expected 8080; got %s", AppPort)
	}

	// Add more assertions for other configuration fields
}
