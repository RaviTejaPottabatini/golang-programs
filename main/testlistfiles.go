package main

import (
    "fmt"
    "io/ioutil"
     "log"
    // "time"
)

func main() {
    files, err := ioutil.ReadDir("d:/Gocode")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            fmt.Println(f.Name())
            fmt.Println(f.ModTime())
    }
  
}