package case1

import (
	cli "github.com/urfave/cli/v2"
	"gitlab.local.com/voda/cli-demo/pkg/utils/xcolor"
)

var Cmd = &cli.Command{
	Name:    "case1",
	Aliases: []string{"c1"},
	Usage:   "测试功能1",
	Action:  Case,
}

func Case(cli *cli.Context) (err error) {

	xcolor.Green("Case 1 已执行！")
	return nil
}
