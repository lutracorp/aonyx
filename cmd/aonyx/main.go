package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/lutracorp/aonyx/internal/app"
	"github.com/lutracorp/aonyx/internal/app/controller/authentication"
	"github.com/lutracorp/aonyx/internal/pkg/database"
	"github.com/lutracorp/aonyx/internal/pkg/server"
	"github.com/urfave/cli/v2"
)

func main() {
	ca := &cli.App{
		Name:  "aonyx",
		Usage: "The Single Sign-On portal for LutraCorp services",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Starts a Aonyx server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "config",
						Value: "config.hcl",
						Usage: "path to a configuration file",
					},
				},
				Action: func(ctx *cli.Context) error {
					var cfg app.Config
					if err := hclsimple.DecodeFile(ctx.String("config"), nil, &cfg); err != nil {
						return err
					}

					if err := database.Open(cfg.Database); err != nil {
						return err
					}

					if err := database.Migrate(); err != nil {
						return err
					}

					api := server.App.Group("/api")

					ac := authentication.NewController()
					ag := api.Group("/auth")
					ag.Post("/login", ac.Login)
					ag.Post("/register", ac.Register)

					return server.Open(cfg.Server)
				},
			},
		},
	}

	if err := ca.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
