package case2

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"gitlab.local.com/voda/cli-demo/pkg/utils/xcolor"
)

// 获取参数Args

var Cmd = &cli.Command{
	Name:    "case2",
	Aliases: []string{"c2"},
	Args:    true,
	Usage:   "获取参数Args用例",
	Action:  Case,
}

func Case(cli *cli.Context) (err error) {
	arg := cli.Args().Get(0)

	fmt.Println(xcolor.Red(fmt.Sprintf("Case 2 已执行, 参数为%s！", arg)))
	return nil
}
