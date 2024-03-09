package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

// auth wraps an http.Handler with basic auth.
func auth(handler http.Handler, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != username || pass != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
func main() {
	var err error
	port := flag.Int("port", 3000, "server port.")
	dir := flag.String("dir", ".", "wrok path.")
	username := flag.String("username", "", "username")
	password := flag.String("password", "", "password")
	flag.Parse()

	addr := ":" + strconv.Itoa(*port)

	log.Println("http-server start on", addr, *dir)
	log.Println("Please open the browser to visit the following address:")
	log.Println("username and password:", *username, *password)

	printInterfaceUrl(err, port)

	// Wrap the file server handler with the auth middleware.
	fileServerHandler := http.FileServer(http.Dir(*dir))
	if *username != "" && *password != "" {
		fileServerHandler = auth(fileServerHandler, *username, *password)
	}

	err = http.ListenAndServe(addr, fileServerHandler)

	if err != nil {
		log.Fatalln(err)
	}

}

func printInterfaceUrl(err error, port *int) {
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
}
