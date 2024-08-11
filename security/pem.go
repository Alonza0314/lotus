package security

import (
	"encoding/pem"
	"errors"
)

func parsePemBlocks(pemData []byte) []*pem.Block {
	blocks := make([]*pem.Block, 0)
	for {
		block, rest := pem.Decode(pemData)
		if block == nil {
			break
		}
		blocks, pemData = append(blocks, block), rest
	}
	return blocks
}

func parsePemFile(pemData []byte) ([]byte, []byte, error) {
	blocks := parsePemBlocks(pemData)

	var certBlock, keyBlock *pem.Block
	for _, block := range blocks {
		switch block.Type {
		case "CERTIFICATE":
			certBlock = block
		case "PRIVATE KEY":
			keyBlock = block
		}
	}

	if certBlock == nil || keyBlock == nil {
		return nil, nil, errors.New("certificate or key missing in Pem file")
	}
	return pem.EncodeToMemory(certBlock), pem.EncodeToMemory(keyBlock), nil
}