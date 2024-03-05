package main

import (
	cli "github.com/urfave/cli/v2"
	"gitlab.local.com/voda/cli-demo/internal/command/case1"
	"gitlab.local.com/voda/cli-demo/internal/command/case2"
	"log"
	"os"
)

const Version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "Cli Demo"
	app.Usage = "Cli Demo"
	app.Version = Version
	app.Commands = []*cli.Command{
		case1.Cmd,
		case2.Cmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
