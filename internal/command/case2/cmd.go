package case2

import (
	cli "github.com/urfave/cli/v2"
	"gitlab.local.com/voda/cli-demo/pkg/utils/xcolor"
)

var Cmd = &cli.Command{
	Name:    "case2",
	Aliases: []string{"c2"},
	Usage:   "测试功能2",
	Action:  Case,
}

func Case(cli *cli.Context) (err error) {

	xcolor.Red("Case 2 已执行！")
	return nil
}
