package maintest
import (
    "fmt"
    "github.com/hefju/gofortest/jutil"
    "os"
)
func HotCompile(filename string){

    gofile:=filename+".go"
    _, err := os.Stat(gofile)//"filename.txt"
    if err != nil {
        fmt.Println("boot failed!\nerror:no file:",gofile)
        return
    }


    //重启服务器的通道
    restartChan := make(chan int)

    go jutil.SetWatcher(".", restartChan)

    fmt.Println("runner begin")
    runner := jutil.Runner{Filename: gofile}//"myhttp"}
    runner.Run()
    runner.WaitForRestart(restartChan)
}
