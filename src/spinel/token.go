package spinel

import "fmt"
import "crypto/sha256"

type Token struct {
	Permissions string `json:"p"`
	Checksum    string `json:"c"`
	Expires     int32 `json:"e"`
}

func (tok *Token) CalculateChecksum(secret string) {
	tok.Checksum = fmt.Sprintf("%x",sha256.Sum256([]byte(secret +":"+ tok.Permissions)))
}

func (tok *Token) Validate(secret string) bool{
	checksum := fmt.Sprintf("%x",sha256.Sum256([]byte(secret +":"+ tok.Permissions)))
	return checksum == tok.Checksum
}

