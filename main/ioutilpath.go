package main

import (
    "fmt"
    "io/ioutil"
    //"log"
    //"os"
   // "path/filepath"
)

func IOReadDir(root string) ([]string, error) {
    var files []string
    fileInfo, err := ioutil.ReadDir(root)
    if err != nil {
        return files, err
    }

    for _, file := range fileInfo {
        files = append(files, file.Name())
    }
    return files, nil
}

func main() {
    var (
        root  string
        files []string
        err   error
    )

    root = "d:/Gocode"

       // ioutil.ReadDir
    files, err = IOReadDir(root)
    if err != nil {
        panic(err)
    }

        for _, file := range files {
        fmt.Println(file)
    }
}