package main

import (
	"fmt"
	"os"

	reqflags "github.com/shouji-kazuo/cli-reqflags"
	"github.com/shouji-kazuo/ts3-card-payment/ts3"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/urfave/cli.v2"
)

var ts3URL = "https://ts3card.com/"

func main() {
	app := &cli.App{
		Name:      "ts3-card-payment",
		Usage:     "",
		ArgsUsage: " ",
		Version:   "v1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "set user name to login ts3card.com",
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "set password to to login ts3card.com",
			},
		},
		Action: func(c *cli.Context) error {
			username := ""
			password := ""

			onMissingUsername := func() error {
				fmt.Print("Enter username: ")
				if _, err := fmt.Fscanln(os.Stdin, &username); err != nil {
					return err
				}
				return nil
			}
			onMissingPassword := func() error {
				fmt.Print("Enter password: ")
				bytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
				fmt.Println()
				if err != nil {
					return err
				}
				password = string(bytes)
				return nil
			}

			err := reqflags.Recover(
				c,
				map[string]func() error{
					"username": onMissingUsername,
					"password": onMissingPassword,
				},
			)
			if err != nil {
				return err
			}

			ts3.Navigate(&ts3.Config{
				URL:      ts3URL,
				Username: "",
				Password: "hoge",
			})

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
