package main

import (
    "fmt"
    //"io/ioutil"
    //"log"
    "os"
    "path/filepath"
)

func FilePathWalkDir(root string) ([]string, error) {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    return files, err
}

func main() {
    var (
        root  string
        files []string
        err   error
    )

    root = "d:/Gocode"
    // filepath.Walk
    files, err = FilePathWalkDir(root)
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }
}
