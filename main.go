package main

import (

	// "github.com/hefju/gofortest/jutil" //采用gopath+路径, 不知道为什么webstorm不能识别,idea14可以
	//"time"
	//"github.com/Unknwon/macaron"
	//"path/filepath"

    "github.com/hefju/gofortest/maintest"
)

//该项目用来测试go代码的,
func main() {
	// jutil.TestIni_configuration()
	// fmt.Println(jutil.IniConfiger["server"]["ip"])

//2
//	m := macaron.Classic()
//	m.Get("/", func() string {
//		return "Hello world!"
//	})
//	m.Run()
//	fmt.Println("end")

	//3.
	//jutil.TraverseFiles(".")
	//jutil.SetWatcher(".")
   // maintest.Main2();
    maintest.HotCompile("myoutput")
    // maintest.Jsonconfigtest()
}
