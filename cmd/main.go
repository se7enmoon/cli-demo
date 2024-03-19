package main

import (
	"github.com/se7enmoon/cli-demo/internal/command/case1"
	"github.com/se7enmoon/cli-demo/internal/command/case2"
	"github.com/se7enmoon/cli-demo/internal/command/case3"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

const Version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "demo"
	app.Usage = "demo"
	app.Version = Version
	app.Flags = []cli.Flag{ // 全局参数，优先级最高
		&cli.StringFlag{
			Name:    "lang",
			Aliases: []string{"l"}, // 使用时需要加 - 前缀 如：-l
			Value:   "english",
			Usage:   "language for the greeting",
		},
	}
	app.Commands = []*cli.Command{
		case1.Cmd,
		case2.Cmd,
		case3.Cmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
