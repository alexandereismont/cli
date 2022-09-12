package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func main() {
	info()
	cliConfig()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "Simple GOLANG Cli"
	app.Usage = "Commands"
	app.Version = "1.0.0"
}

func cliConfig() {
	app.EnableBashCompletion = true
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:    "greeting",
			Aliases: []string{"g"},
			Usage:   "prints welcome message",
			Action: func(cCtx *cli.Context) error {
				if cCtx.String("username") != "" {
					fmt.Println("Hello user: ", cCtx.String("username"))
				}
				fmt.Println("Welcome to this cli ")
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "username",
					Aliases: []string{"u"},
					Usage:   "prints a username with a hello message",
				},
			},
		},
	}
}
