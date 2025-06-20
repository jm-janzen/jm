package main

import (
	"context"
	"fmt"
	"log"
	"os"

	apiRoutes "jm/domains/api.jmjanzen.com/routes"
	comRoutes "jm/domains/jmjanzen.com/routes"
	cvRoutes "jm/domains/jmjanzen.cv/routes"

	"github.com/urfave/cli/v3"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cmd := &cli.Command{
		Action: func(ctx context.Context, cmd *cli.Command) error {
			domainName := cmd.Args().First()
			switch domainName {
			case "jmjanzen.com":
				comRoutes.Launch()
			case "api.jmjanzen.com":
				apiRoutes.Launch()
			case "jmjanzen.cv":
				cvRoutes.Launch()
			default:
				message := fmt.Sprintf("Not a known domain name '%s'", domainName)
				return cli.Exit(message, 1)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
