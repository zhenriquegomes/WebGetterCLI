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
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search for IP addresses on web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "gooogle.com.br",
				},
			},
			Action: searchIP,
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
