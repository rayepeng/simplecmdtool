# 一个简单的go 命令行工具


## 目的

方便编写命令行工具，在cli的基础上增加了解析参数的功能，只需要定义好函数、参数即可

## 使用

比如我要写一个读文件的函数，做一些操作，比如逆序输出每一行

那么只需要

```go
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

tool := tools.NewTool([]*tools.FunctionConfig{
    {
        Name:        "ReverseFileLines",
        Function:    ReverseFileLines,
        OptionFlags: "f:",
        Description: "-f <file_path>",
    },
})
tool.Run(os.Args) 
```

即可
