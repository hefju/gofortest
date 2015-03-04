package maintest

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type Config struct {
    mainserver
}

type mainserver struct {
    Ip      string
    Port    int
    Comment string
}

func Jsonconfigtest() {
    ToFile()
    FromFile()
    fmt.Println("cao")
}

func ToFile() {
    c := Config{}
    c.mainserver.Ip = "127.0.0.1"
    c.mainserver.Port = 9800
    c.Comment = "这是主服务器的配置啊"
    // c := new(config)
    data, err := json.Marshal(c)
    if err != nil {
        fmt.Println(err)
    }

    fileName := "conf2.json"
    dstFile, err := os.Create(fileName)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    defer dstFile.Close()
    dstFile.WriteString(string(data))
}
func FromFile() {
    r, err := os.Open("conf2.json")
    if err != nil {
        log.Fatalln(err)
    }
    decoder := json.NewDecoder(r)
    var c Config
    err = decoder.Decode(&c)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(c)
}
