package main

  import (
      "log"
      "fmt"

      "github.com/fsnotify/fsnotify"
  )

  func main() {


    

      watcher, err := fsnotify.NewWatcher()
      if err != nil {
          log.Fatal(err)
      }
            err = watcher.Add("D:\\rapiscan\\input\\632DV")
      if err != nil {
          log.Fatal(err)
      }

      defer watcher.Close()

       func() {
          for {
              select {
              case event := <-watcher.Events:
                  fmt.Println("event: in", event)
                  if event.Op&fsnotify.Write == fsnotify.Write {
                      fmt.Println("modified file:", event.Name)
                  }
              case err := <-watcher.Errors:
                  fmt.Println("error:", err)
              }
          }
      }()

     
  }