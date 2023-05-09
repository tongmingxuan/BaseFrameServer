// Package Command /*
package Command

import (
	"fmt"
	"github.com/urfave/cli"
)

type DemoCommand struct {
}

func (command DemoCommand) Command() cli.Command {
	return cli.Command{
		Name:  "demo",
		Usage: "run",
		Subcommands: []cli.Command{
			{
				Name:   "start",
				Usage:  "DemoCommand:start",
				Action: command.Run,
			},
		},
	}
}

func (command DemoCommand) Run(c *cli.Context) error {

	fmt.Println("DemoCommand:运行")

	return nil
}
