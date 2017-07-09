package main

import (
	"os"
	"net"
	"time"
	"io/ioutil"
	"github.com/urfave/cli"
	"fmt"
	"github.com/asaskevich/govalidator"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-whois"
	app.Usage = "a WHOIS parser tool"
	app.Version = "0.1"
	app.Action = run

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	if len(c.Args()) > 0 && checkArgsIsIP(c.Args()){
		for i := 0; i < len(c.Args()); i++ {
			ip := c.Args().Get(i)
			ret, _ := GetWhoIS(ip, "whois.apnic.net")
			fmt.Println(ret)
		}
	} else {
		fmt.Println("go-whois: try 'go-whois --help' or 'go-whois -h' for more information ")
	}
	return nil
}

func checkArgsIsIP(args cli.Args) bool{
	for i := 0; i < len(args); i++ {
		ip := args.Get(i)
		if !govalidator.IsIP(ip) {
			return false
		}
	}
	return true
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
