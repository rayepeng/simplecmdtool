package tools_test

import (
	"fmt"
	"testing"

	"github.com/rayepeng/simplecmdtool/tools"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func Func1(ctx *cli.Context) error {
	uin := ctx.String("u")
	operatorID := ctx.String("p")

	if uin == "" || operatorID == "" {
		return fmt.Errorf("both uin and operator_id must be provided")
	}

	fmt.Printf("Func1 called with uin: %s and operator_id: %s\n", uin, operatorID)
	return nil
}

func TestToolRunWithFunc1(t *testing.T) {
	tool := tools.NewTool()

	func1Config := &tools.FunctionConfig{
		Name:        "Func1",
		Function:    Func1,
		OptionFlags: "u:p:",
		Description: "-u <uin> -p <operator_id>",
	}

	tool.AddFunction(func1Config)

	// Prepare the arguments for the app to run with.
	args := []string{"appName", "Func1", "-u", "123", "-p", "456"}

	// Run the tool with the provided arguments.
	err := tool.Run(args)

	// Check if Func1 was executed without any error.
	assert.NoError(t, err)
}
