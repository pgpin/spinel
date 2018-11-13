package spinel

import "net"

func CidrsParse(cidrs []string) []*net.IPNet{
	ipnets := make([]*net.IPNet, len(cidrs))
	for i,cidr := range cidrs{
		_, ipnets[i], _ = net.ParseCIDR(cidr)
	}
	return ipnets
}

func CidrsContains(cidrs *[]*net.IPNet, ip string) bool{
	nip := net.ParseIP(ip)
	for _, ipnet := range *cidrs{
		if ipnet.Contains(nip){
			return true
		}
	}
	return false 
}
