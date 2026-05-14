package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

func Generate()  *cli.App {
	app := cli.NewApp()
	app.Name = "Command-line application"
	app.Usage = "Search IPs in the server"

	flags := []cli.Flag {
				&cli.StringFlag{
					Name: "host",
					Value: "jlcodetech.vercel.app",
				},
	}

	app.Commands = []*cli.Command{
		{
			Name: "ip",
			Usage: "Searching address to the web",
			Flags: flags,
			Action: fetchIps,
		},
		{
			Name: "servers",
			Usage: "Fecth servers in the web",
			Flags: flags,
			Action: fetchServers,
		},
	}

	return app
}	


func fetchIps( c *cli.Context ) error{
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

    for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}

func fetchServers(c *cli.Context) error{
	host := c.String("host")

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
	
	return nil
}