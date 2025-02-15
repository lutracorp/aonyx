package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/lutracorp/aonyx/internal/app"
	"github.com/lutracorp/aonyx/internal/app/controller/authentication"
	"github.com/lutracorp/aonyx/internal/app/controller/user"
	"github.com/lutracorp/aonyx/internal/app/middleware"
	"github.com/lutracorp/aonyx/internal/pkg/database"
	"github.com/lutracorp/aonyx/internal/pkg/server"
	"github.com/matthewhartstonge/argon2"
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

					arc := argon2.DefaultConfig()

					ac := authentication.NewController(&arc)
					ag := api.Group("/auth")
					ag.Post("/login", ac.Login)
					ag.Post("/register", ac.Register)

					uc := user.NewController(&arc)
					ug := api.Group("/users", middleware.User)
					ug.Get("/@me", uc.GetCurrent)
					ug.Patch("/@me", uc.ModifyCurrent)

					return server.Open(cfg.Server)
				},
			},
		},
	}

	if err := ca.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
