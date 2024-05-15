package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar, vai retornar a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "youtube.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs",
			Flags:  flags,
			Action: buscarIP,
		},
		{
			Name:   "server",
			Usage:  "Busca servidores",
			Flags:  flags,
			Action: buscarServer,
		},
	}

	return app
}

func buscarIP(c *cli.Context) {
	salvaHost := c.String("host")

	//net
	ips, erro := net.LookupIP(salvaHost)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServer(c *cli.Context) {
	salvarHost := c.String("host")

	servidores, erro := net.LookupNS(salvarHost) //name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
