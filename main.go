/*
编写一个简单的 http服务器,返回文件
*/
package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 3000, "server port.")
	dir := flag.String("dir", ".", "wrok path.")
	addr := ":" + strconv.Itoa(*port)
	log.Println("litongjava-go-simple-http-server start on", addr, *dir)
	err := http.ListenAndServe(addr, http.FileServer(http.Dir(*dir)))
	if err != nil {
		log.Fatalln(err)
	}
}
