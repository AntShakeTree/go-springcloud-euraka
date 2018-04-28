package tools

import (
	"net"
	"os"
	"strings"
)

func Hostname() string {
	host, _ := os.Hostname()
	return host
}

func IpAddr() string {
	mac_address := make([]string, 10)
	interfaces, err := net.Interfaces()
	if err != nil {
		os.Exit(-1)
	}
	var index int = int(^uint(0) >> 1)
	var strs = make([]string, 1)
	for _, inter := range interfaces {
		flags := inter.Flags.String()
		if strings.Contains(flags, "up") && strings.Contains(flags, "broadcast") {
			mac_address = append(mac_address, inter.HardwareAddr.String())
			if index >= inter.Index {
				index = inter.Index
				ars, _ := inter.Addrs()
				if len(ars) > 1 {
					ips := strings.Split(ars[1].String(), "/")
					strs[0] = ips[0]

				} else if len(ars) == 1 {
					if strings.Contains(ars[0].String(), "/") {
						ips := strings.Split(ars[0].String(), "/")
						strs[0] = ips[0]
					} else {
						strs[0] = ars[0].String()
					}

				}
			}
		}
	}
	if len(strs) == 0 {
		return getLocalIP();
	}
	return strs[0]
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
