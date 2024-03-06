package case3

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"gitlab.local.com/voda/cli-demo/pkg/utils/xcolor"
)

// 获取参数Args

var Cmd = &cli.Command{
	Name:    "case3",
	Aliases: []string{"c3"},
	Args:    true,
	Usage:   "Flag参数用例",
	Action:  Case,
}

func Case(cli *cli.Context) (err error) {
	arg := cli.Args().Get(0)
	var ans string
	if cli.String("lang") == "spanish" {
		ans = fmt.Sprintf("Hola %s", arg)
	} else {
		ans = fmt.Sprintf("Hello %s", arg)
	}
	fmt.Println(xcolor.Red(ans))
	return nil
}
