package main

import (
	"os"
	"fmt"
	"net"
	"time"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Command syntax: go-whois [ip]")
		return
	}

	d, err := GetWhoIS(os.Args[1], "whois.apnic.net")
	if err != nil {
		fmt.Println("GetWhoIS failed:", err)
	} else {
		fmt.Println(d)
	}
}

func GetWhoIS(ip string, server string)(string, error){
	var retError error
	var ret string

	connection, err := net.DialTimeout("tcp", net.JoinHostPort(server, "43"), time.Second * 5)
	if err != nil {
		retError = err
		return "", retError
	}

	defer connection.Close()

	connection.Write([]byte(ip + "\r\n"))
	buffer, err := ioutil.ReadAll(connection)
	if err != nil {
		retError = err
		return "", retError
	}
	ret = string(buffer)

	return ret, retError
}
