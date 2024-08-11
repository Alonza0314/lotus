package test

import (
	"lotus/security"
	"testing"
)

func TestLoadTLSCertificate(t *testing.T) {
	if _, err := security.LoadTLSCertificate("test.pem"); err != nil {
		t.Fatal(err)
	}
}
