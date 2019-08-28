package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {

	addrs, err := net.InterfaceAddrs()
	//name, err := os.Hostname()

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	cnt := 1

	for _, get_addrs := range addrs {
		if ipnet, ok := get_addrs.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To16() != nil {
				fmt.Println(strconv.Itoa(cnt) + ":" + ipnet.IP.String())
				cnt = cnt + 1
			}
		}
	}
}
