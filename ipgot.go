package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"log"
)

func get_local_ip() string {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func get_public_ip() string {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	pub_ip := string(body)
	return pub_ip
}


func main() {
	localip := get_local_ip()
	publicip := get_public_ip()

	fmt.Println(localip)
	fmt.Println(publicip)
}
