package tools

import (
	"fmt"
	"os"
	"regexp"

	"github.com/urfave/cli/v2"
)

type Function func(ctx *cli.Context) error

type FunctionConfig struct {
	Name        string
	Function    Function
	OptionFlags string
	Description string
}

type Tool struct {
	app             *cli.App
	functionConfigs []*FunctionConfig
}

func NewTool() *Tool {
	return &Tool{
		app: cli.NewApp(),
	}
}

func (t *Tool) AddFunction(config *FunctionConfig) {
	t.functionConfigs = append(t.functionConfigs, config)
}

func parseFlags(flagString string, description string) ([]cli.Flag, error) {
	flags := []cli.Flag{}

	// 使用正则表达式从 flagString 中匹配字母和冒号组合
	flagRegex := regexp.MustCompile(`\w:`)
	flagPairs := flagRegex.FindAllString(flagString, -1)

	// 使用正则表达式从 Description 中提取选项的使用说明
	usageRegex := regexp.MustCompile(`-(\w) <.*?>`)
	usageMatches := usageRegex.FindAllStringSubmatch(description, -1)

	if len(flagPairs) != len(usageMatches) {
		return nil, fmt.Errorf("number of option descriptors and usages do not match")
	}

	usageMap := make(map[string]string)
	for _, match := range usageMatches {
		usageMap[match[1]] = match[0]
	}

	for _, pair := range flagPairs {
		flagName := string(pair[0])
		usage, ok := usageMap[flagName]
		if !ok {
			return nil, fmt.Errorf("flag %s not found in description", flagName)
		}

		flag := &cli.StringFlag{
			Name:  flagName,
			Usage: usage,
		}

		flags = append(flags, flag)
	}

	return flags, nil
}

func (t *Tool) Run() error {

	for _, config := range t.functionConfigs {
		flags, err := parseFlags(config.OptionFlags, config.Description)
		if err != nil {
			return err
		}
		t.app.Commands = append(t.app.Commands, &cli.Command{
			Name: config.Name,
			Action: func(c *cli.Context) error {
				return config.Function(c)
			},
			Flags: flags,
		})

	}

	return t.app.Run(os.Args)
}