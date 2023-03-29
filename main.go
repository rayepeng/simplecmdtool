package main

import (
	"fmt"

	"github.com/rayepeng/simplecmdtool/tools"
	"github.com/urfave/cli/v2"
)

func main() {
	tool := tools.NewTool()

	func1Config := &tools.FunctionConfig{
		Name:        "Func1",
		Function:    Func1,
		OptionFlags: "u:p:",
		Description: "-u <uin> -p <operator_id>",
	}
	tool.AddFunction(func1Config)

	err := tool.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func Func1(ctx *cli.Context) error {
	fmt.Printf("Function: Func1, UIN: %s, OperatorID: %s\n", ctx.Args().Get('u'), ctx.Args().Get('p'))
	return nil
}