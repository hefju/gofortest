package main

import (
	"fmt"
	"github.com/hefju/gofortest/jutil"
)

//该项目用来测试go代码的
func main() {
	jutil.ParseIni("conf.ini")
	fmt.Println(jutil.IniConfiger)
	fmt.Println("end")
}
