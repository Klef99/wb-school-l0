package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/urfave/cli/v3"

	"github.com/klef99/wb-school-l0/cmd/commands"
	"github.com/klef99/wb-school-l0/cmd/di"
)

func main() {
	app := &cli.Command{
		Name:     "assets",
		Flags:    []cli.Flag{},
		Commands: commands.Commands(),
		Action: func(context context.Context, cli *cli.Command) error {
			_, shutdown, err := di.InitializeDependencies()
			if err != nil {
				return fmt.Errorf("error initialize application. %w", err)
			}

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, os.Kill)

			defer signal.Stop(quit)

			<-quit

			shutdown()

			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal("error run application. ", err)
	}
}
