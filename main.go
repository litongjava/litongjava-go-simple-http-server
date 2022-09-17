/*
编写一个简单的 http服务器,返回文件
*/
package main

import (
  "log"
  "net/http"
  "os"
)

func main() {
  port := "3000"
  for i := 1; i < len(os.Args); i += 2 {
    param := os.Args[i]
    if param == "--port" {
      port = os.Args[i+1]
    }
  }
  log.Println("litongjava-go-simple-http-server start on", port)

  http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))
  log.Println("started")
}
