package security

import (
	"crypto/tls"
	"log"
	"os"
)

func LoadTLSCertificate(pemPath string) *tls.Certificate {
	pemData, err := os.ReadFile(pemPath)
	if err != nil {
		log.Fatalln("Failed to read PEM file:", err)
	}

	cert, key, err := parsePemFile(pemData)
	if err != nil {
		log.Fatalln("Failed to parse Pem file:", err)
	}

	tlsCert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Fatalln("Failed to load certificate:", err)
	}

	return &tlsCert
}