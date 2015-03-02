//这个文件是用来存放我的实现性质的代码, 等到功能实现了就写入到项目里面
package maintest

import (
//"bytes"
    "fmt"
    "github.com/hefju/gofortest/jutil"
)

func Main2() {
    //重启服务器的通道
    restartChan := make(chan int)

    go jutil.SetWatcher(".", restartChan)

    fmt.Println("runner begin")
    runner := jutil.Runner{"myhttp"}
    runner.Run()
    runner.WaitForRestart(restartChan)
    // for {
    // 	var input string
    // 	fmt.Scan(&input)
    // 	if input == "ex" {
    // 		break
    // 	} else {
    // 		fmt.Println("bad cmd")
    // 	}
    // }

}
