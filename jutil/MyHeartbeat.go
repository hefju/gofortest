package jutil
import (
    "time"
    "fmt"
)

//用来正数和倒数的
func Myheartbeat(count int) {
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
