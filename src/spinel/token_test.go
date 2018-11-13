package spinel

import "testing"

func TestToken(t *testing.T){
  a := &Token{Permissions: "*", Expires: 12345}
	a.CalculateChecksum("secret")
  if ! a.Validate("secret") {
		t.Error("failed to validate the checksum")
	}
  if a.Validate("wrongsecret") {
		t.Error("validate with wrong secret passed")
	}
}
