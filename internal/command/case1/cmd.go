package case1

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"gitlab.local.com/voda/cli-demo/pkg/utils/xcolor"
)

var Cmd = &cli.Command{
	Name:    "case1",
	Aliases: []string{"c1"},
	Usage:   "基本用例",
	Action:  Case,
}

func Case(cli *cli.Context) (err error) {
	fmt.Println(xcolor.Green("Case 1 已执行！"))
	return nil
}
