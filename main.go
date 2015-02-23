package main

import (
	"fmt"
	"github.com/hefju/gofortest/jutil"//采用gopath+路径, 不知道为什么webstorm不能识别,idea14可以
)

//该项目用来测试go代码的,
func main() {
	jutil.TestIni_configuration()
//	fmt.Println(jutil.IniConfiger["server"]["ip"])
	fmt.Println("end")
}
