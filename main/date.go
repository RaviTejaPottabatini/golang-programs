package main

import (
    "fmt"
    "io/ioutil"
     "log"
     "time"
)

func main() {
    var f func()
    var t *time.Timer

    



    f = func () {
         files, err := ioutil.ReadDir("d:/Gocode")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            fmt.Println(f.Name())
            fmt.Println(f.ModTime())
    }
        t = time.AfterFunc(time.Duration(5) * time.Second, f)
    }

    t = time.AfterFunc(time.Duration(5) * time.Second, f)

    defer t.Stop()

    //simulate doing stuff
    time.Sleep(time.Minute)
}