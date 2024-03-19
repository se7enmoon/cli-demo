package case1

import (
	"fmt"
	"github.com/se7enmoon/cli-demo/pkg/utils/xcolor"
	cli "github.com/urfave/cli/v2"
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
