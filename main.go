/*
编写一个简单的 http服务器,返回文件
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	var err error
	port := flag.Int("port", 3000, "server port.")
	dir := flag.String("dir", ".", "wrok path.")
	flag.Parse()

	addr := ":" + strconv.Itoa(*port)

	log.Println("http-server start on", addr, *dir)
	log.Println("Please open the browser to visit the following address:")

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Unable to obtain network interface list:", err)
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Println("Unable to obtain IP address list:", err)
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && ipNet.IP.To4() != nil {
				fmt.Printf("http://%v:%v\n", ipNet.IP, *port)
			}
		}
	}

	err = http.ListenAndServe(addr, http.FileServer(http.Dir(*dir)))

	if err != nil {
		log.Fatalln(err)
	}

}
