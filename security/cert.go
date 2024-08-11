package security

import (
	"crypto/tls"
	"errors"
	"os"
)

func LoadTLSCertificate(pemPath string) (*tls.Certificate, error) {
	pemData, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, errors.New("failed to read PEM file:\n\t" + err.Error())
	}

	cert, key, err := parsePemFile(pemData)
	if err != nil {
		return nil, errors.New("failed to parse Pem file:\n\t" + err.Error())
	}

	tlsCert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return nil, errors.New("failed to load certificate:\n\t" + err.Error())
	}

	return &tlsCert, nil
}
