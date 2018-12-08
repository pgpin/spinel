package spinel

import "fmt"
import "crypto/sha256"
//import "gopkg.in/yaml.v2"
//import "encoding/base64"
import "encoding/json"

type Token struct {
	Permissions string `json:"p"`
	Checksum    string `json:"c"`
	Expires     int64 `json:"e"`
}

func NewToken( secret string, permissions string, expires int64 ) *Token{
	tok := &Token{Permissions: permissions, Expires: expires}
	tok.CalculateChecksum(secret)
	return tok
}

func (tok *Token) CalculateChecksum(secret string) {
	tok.Checksum = fmt.Sprintf("%x",sha256.Sum256([]byte(secret +":"+ tok.Permissions +":"+ string(tok.Expires))))
}

func (tok *Token) Validate(secret string) bool{
	checksum := fmt.Sprintf("%x",sha256.Sum256([]byte(secret +":"+ tok.Permissions +":"+ string(tok.Expires))))
	return checksum == tok.Checksum
}

func ( tok *Token) AsJsonString() string{
	bytes, _ := json.Marshal(tok)
	//return  base64.StdEncoding.EncodeToString(bytes)
	return  string(bytes)
}
