package main

import (
	"fmt"
	"net"
)

func getLocalIP() (ips []string, err error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("get ip interfaces error:", err)
		return
	}

	for _, i := range ifaces {
		addrs, errRet := i.Addrs()
		if errRet != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				if ip.IsGlobalUnicast() {
					ips = append(ips, ip.String())
				}
			}
		}
	}
	return
}

func main() {
	ip, err := getLocalIP()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// ipV4 := ip[2]
	// fmt.Println(ipV4)
	fmt.Println(ip)
}