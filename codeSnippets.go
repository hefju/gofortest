//这个文件是用来存放我的实现性质的代码, 等到功能实现了就写入到项目里面
package main

import (
//"bytes"
    "fmt"
    "github.com/go-fsnotify/fsnotify"
    "io/ioutil"
    "log"
    "os/exec"
    "path"
    "time"
)

func main2() {
    //重启服务器的通道
    restartChan := make(chan int)

    go SetWatcher(".", restartChan)

    fmt.Println("runner begin")
    runner := runner{"myhttp"}
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

type runner struct {
    filename string
}

//重启服务器
func (x runner) WaitForRestart(order chan int) {
    for {
        <-order
        x.Kill()
        x.Run()
    }
}

//执行命令
func (x runner) Run() {
    goname := x.filename + ".go"
    fmt.Println("Run:" + goname)
    c := exec.Command("go", "run", goname) // "myserver.go")

    c.Start()
}

//终结进程.
func (x runner) Kill() {
    taskname := x.filename + ".exe"
    fmt.Println("Kill:" + taskname)
    c := exec.Command("taskkill.exe", "/f", "/im", taskname) //"myserver.exe")
    err := c.Start()
    if err != nil {
        fmt.Println(err)
    }
}

func (x runner) Loop() {
    for {
        fmt.Println("Loop:" + x.filename)
        x.Run()
        go myheartbeat(10)
        time.Sleep(time.Second * 10)

        x.Kill()
        go myheartbeat(-5)
        time.Sleep(time.Second * 5)
    }
}

func myheartbeat(count int) {
    max := 0
    heartbeat := 0
    if count > 0 {
        max = count
    } else {
        max = 0 - 1
        heartbeat = count - 1
    }

    ticker := time.NewTicker(time.Second)
    for {
        <-ticker.C
        heartbeat++
        if heartbeat > max {
            break
        }
        out := heartbeat
        if out < 0 {
            out = 0 - out
        }
        fmt.Println(out)
    }
}

var lstFolder []string

func SetWatcher(root string, order chan int) {
    lstFolder = make([]string, 0)
    lstFolder = append(lstFolder, root)
    GetFoleder(root)

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    done := make(chan bool)
    go func() {
        for {
            select {
            case event := <-watcher.Events:
            //如果有go文件修改,才触发更新
                if path.Ext(event.Name) == ".go" && event.Op&fsnotify.Write == fsnotify.Write {
                    //fmt.Println(event.Name, "--", event.Op, "--", event)
                    order <- 0
                }

            case err := <-watcher.Errors:
                log.Println("error:", err)
            }
        }
    }()
    fmt.Println(".watcher.") //开始监视

    for _, f := range lstFolder {
        err = watcher.Add(f)
        if err != nil {
            log.Fatal("138:", err)
        }
    }

    <-done
}

func GetFoleder(root string) {
    files, _ := ioutil.ReadDir(root)
    for _, fi := range files {
        if fi.IsDir() {
            if fi.Name()[0] == '.' { //不处理以.开头的文件夹
                continue
            }
            mypath := path.Join(root, fi.Name())
            lstFolder = append(lstFolder, mypath)
            GetFoleder(mypath)
            //fmt.Println(mypath)
        }
    }
}

