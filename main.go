package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rayepeng/simplecmdtool/tools"
	"github.com/urfave/cli/v2"
)

func main() {
	tool := tools.NewTool([]*tools.FunctionConfig{
		{
			Name:        "Func1",
			Function:    Func1,
			OptionFlags: "u:p:",
			Description: "-u <uin> -p <operator_id>",
		},
		{
			Name:        "ReverseFileLines",
			Function:    ReverseFileLines,
			OptionFlags: "f:",
			Description: "-f <file_path>",
		},
	})
	err := tool.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func Func1(ctx *cli.Context) error {
	fmt.Printf("Function: Func1, UIN: %s, OperatorID: %s\n", ctx.String("u"), ctx.String("p"))
	return nil
}

func ReverseFileLines(ctx *cli.Context) error {
	filePath := ctx.String("f")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Println(lines[i])
	}

	return nil
}
