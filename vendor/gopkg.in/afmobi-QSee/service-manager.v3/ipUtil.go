package service_manager

import (
	"net"
	"fmt"
	"os"
)

func GetIP() string{

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	myip := ""
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		ipnet, ok := address.(*net.IPNet);
		if  ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				myip = ipnet.IP.String()
				fmt.Println(ipnet.IP.String())
			}
		}
	}
	return myip
}
