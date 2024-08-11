package test

import (
	"lotus/security"
	"testing"
)

func TestLoadTLSCertificate(t *testing.T) {
	if tlsCert := security.LoadTLSCertificate("test.pem"); tlsCert == nil {
		t.Fatal("Expected non-nil TLS certificate")
	}
}
