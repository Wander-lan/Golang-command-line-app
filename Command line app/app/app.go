package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// This will return the app ready for being executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line app"
	app.Usage = "Searches for IPs and server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Searches for IPs",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers_names",
			Usage:  "Searches for servers names on the internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

// Searches public IPs of a given website
func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// Searches for servers names from a given website
func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host) // server name
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}
