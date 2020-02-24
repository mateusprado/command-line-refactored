package cli

import (
	"fmt"
	"time"

	"github.com/mateusprado/command-line-refactored/types"
	"github.com/urfave/cli/v2"
)

var (
	dryRunFlags types.DryRunFlags
)

func Builder() *cli.App {

	app := &cli.App{
		Name:     "sre-cli",
		Compiled: time.Now(),
		Commands: []*cli.Command{
			{
				Name: "dry-run",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "url",
						Destination: &dryRunFlags.Url,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "user",
						Destination: &dryRunFlags.User,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "token",
						Destination: &dryRunFlags.Token,
						Required:    true,
					},
				},
				Action: DryRun,
			},
		},
	}

	return app

}

func DryRun(c *cli.Context) error {
	fmt.Printf("user %s, login in %s with auth token %s\n", dryRunFlags.User, dryRunFlags.Url, dryRunFlags.Token)
	return nil
}
