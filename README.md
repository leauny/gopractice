# gopractice


# Go的项目管理

1. 使用`go mod init github.com/leauny/gopractice`来初始化项目
2. 同目录下pacakge名称需要想同，同一个package内可以自由相互调用（无需大写）。
   同一个目录下所有go文件合并为同一个文件是等价的。
3. 使用`go get github.com/bytedance/sonic`来获取第三方库, 或使用`go mod tidy`
4. 每个包中可以有多个init函数，如果不调用包中的内容但需要执行其中的init，则带下划线书写。如: `import _ "fmt"`
