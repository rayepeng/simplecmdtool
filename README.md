# 一个简单的go 命令行工具


## 目的

方便编写命令行工具，在cli的基础上增加了解析参数的功能，只需要定义好函数、参数即可

## 使用

比如我要写一个读文件的函数，做一些操作，比如逆序输出每一行

那么只需要将函数添加到tools即可，自动添加到命令行中

```go
func ReverseFileLines(ctx *cli.Context) error {
	filePath := ctx.String("f")

	data, err := os.ReadFile(filePath)
	//...
	return nil
}

// 添加函数
tool := tools.NewTool([]*tools.FunctionConfig{
    {
        Name:        "ReverseFileLines",
        Function:    ReverseFileLines,
        OptionFlags: "f:",
        Description: "-f <file_path>",
    },
})
err := tool.Run(os.Args) 
if err != nil{
	//....
}
```

命令行输出：

```bash
❯ go run main.go ReverseFileLines -h
NAME:
   main ReverseFileLines

USAGE:
   main ReverseFileLines [command options] [arguments...]

OPTIONS:
   -f value, --file_path value  -f <file_path>
   --help, -h                   show help
```

## 特性

- [ ] 自动解析选项，选项别名
- [ ] 文档