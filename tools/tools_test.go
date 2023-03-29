package tools

import (
	"fmt"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestNewTool(t *testing.T) {
	tool := NewTool()
	if tool == nil {
		t.Errorf("Expected tool to be initialized, but got nil")
	}
}

func TestAddFunction(t *testing.T) {
	tool := NewTool()
	func1 := func(ctx *cli.Context) error {
		fmt.Println(ctx.Args().Get('u'))
		return nil
	}
	functionConfig := &FunctionConfig{
		Name:        "TestFunc",
		Function:    func1,
		OptionFlags: "u:p:",
		Description: "-u <uin> -p <operator_id>",
	}
	tool.AddFunction(functionConfig)

	if len(tool.functionConfigs) != 1 {
		t.Errorf("Expected functionConfigs length to be 1, but got %d", len(tool.functionConfigs))
	}

	if tool.functionConfigs[0].Name != "TestFunc" {
		t.Errorf("Expected functionConfig name to be 'TestFunc', but got '%s'", tool.functionConfigs[0].Name)
	}
}

// 其他测试函数...
