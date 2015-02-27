//这个文件是用来存放我的实现性质的代码, 等到功能实现了就写入到项目里面
package main

import (
	"fmt"
	"os/exec"
	"time"
)
//这是一个独立的程序, 放在myexec.go文件里面. gobuild之后myexec.exe放在myserver.go同一层文件夹下, 就可以启动myserver了
func main() {

	done := make(chan bool)
	ticker := time.NewTicker(time.Second)
	heartbeat := 0

	cmd := exec.Command("go", "run", "myserver.go")
	cmd.Start()

	go func() {
		time.Sleep(time.Second * 20)
		c := exec.Command("taskkill.exe", "/f", "/im", "myserver.exe")
		err := c.Start()
		if err != nil {
			fmt.Println(err)
		}
		done <- true
	}()

	go func() {
		for {
			<-ticker.C
			heartbeat++
			fmt.Println(heartbeat)
		}
	}()
	<-done
	fmt.Println("myexec.go cao")
}
