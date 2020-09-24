package xip

import (
	"net"
)

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, adr := range addrs {
		if ip,ok:=adr.(*net.IPNet);ok && !ip.IP.IsLoopback(){
			if ip.IP.To4() !=nil{
				return ip.IP.To4().String()
			}
		}
	}
	return ""
}
