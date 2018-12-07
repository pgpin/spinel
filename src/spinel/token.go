package spinel

import "fmt"
import "crypto/sha256"

type Token struct {
	Permissions string `json:"p"`
	Checksum    string `json:"c"`
	Expires     int64 `json:"e"`
}

func (tok *Token) CalculateChecksum(secret string) {
	tok.Checksum = fmt.Sprintf("%x",sha256.Sum256([]byte(secret +":"+ tok.Permissions +":"+ string(tok.Expires))))
}

func (tok *Token) Validate(secret string) bool{
	checksum := fmt.Sprintf("%x",sha256.Sum256([]byte(secret +":"+ tok.Permissions +":"+ string(tok.Expires))))
	return checksum == tok.Checksum
}

