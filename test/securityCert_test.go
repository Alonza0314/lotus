package test

import (
	"github.com/Alonza0314/lotus/security"
	"testing"
)

func TestLoadTLSCertificate(t *testing.T) {
	if _, err := security.LoadTLSCertificate("test.pem"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
