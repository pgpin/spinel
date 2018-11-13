package spinel

import "testing"
import "log"

func TestCidrAllowed(t *testing.T){
	ranges := make([]string,1) 
	ip := "10.0.0.0/16"
	ranges[0] = ip
	cidrs := CidrsParse(ranges)
	log.Println(cidrs)
	if ! CidrsContains( &cidrs, "10.0.0.1") {
		t.Error("did not find ip address in cidr blocks")
	}

	if CidrsContains( &cidrs, "127.0.0.1" ) {
		t.Error("found incorrect IP in cidr block")
	}
}
