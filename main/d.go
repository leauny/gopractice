package main

import (
	"fmt"
	"github.com/leauny/gopractice/util"
	// 导入math目录下的maths包
	"github.com/leauny/gopractice/util/math"
	// 将maths包重命名为math
	math "github.com/leauny/gopractice/util/math"
	// 第三方库（字节跳动的json库）
	"github.com/bytedance/sonic"
	// 不调用但执行包下面的init函数
	_ "github.com/bytedance/sonic"
)

func main() {
	a, b, c := 1, 2, 3
	fmt.Println(util.Name) // 使用其他包内的变量需要加上包名
	fmt.Println(util.Add(a, b))
	fmt.Println(maths.Add(a, b, c))
	fmt.Println(math.Add(a, b, c)) // 别名调用
	bytes, _ := sonic.Marshal("hello")
	fmt.Println(string(bytes))
}
