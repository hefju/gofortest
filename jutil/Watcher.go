package jutil
import (
	"fmt"
	"io/ioutil"

	"github.com/go-fsnotify/fsnotify"
	"log"
	"path"
)

var lstFolder []string

func SetWatcher(root string) {
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
					fmt.Println(event.Name, "--", event.Op, "--", event)
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
			log.Fatal(err)
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
