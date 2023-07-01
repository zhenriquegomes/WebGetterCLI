package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Returns the CLI app ready to execution
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "WebGetterCLI"
	app.Usage = "Search for IP addresses and servers"
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IP addresses on web",
			Flags:  flags,
			Action: searchIP,
		},
		{
			Name:   "server",
			Usage:  "Search for servers on web",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIP(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
