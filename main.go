package main

import (
	"fmt"
	"github.com/driverzhang/go-image-hosting/entities"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bd"
	app.Usage = "自动化github图床工具"
	app.Version = entities.Version
	app.Commands = []cli.Command{
		{
			Name:    "bd",
			Aliases: []string{"a"},
			Usage:   "upload a picture to image hosting",
			//Flags: []cli.Flag{
			//	cli.StringFlag{
			//		Name:        "p",
			//		Value:       "",
			//		Usage:       "An absolute path",
			//		Destination: &entities.Clip.AP,
			//	},
			//	cli.StringFlag{
			//		Name:        "r",
			//		Value:       "",
			//		Usage:       "github repo name",
			//		Destination: &entities.Clip.Repo,
			//	},
			//},
			Action: entities.RunClipBoard,
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "image hosting version",
			Action: func(c *cli.Context) error {
				fmt.Println(entities.GetVersion())
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
