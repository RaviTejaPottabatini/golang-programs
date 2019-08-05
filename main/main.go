package main

  import (
      "log"
      "fmt"
      "os"
      "io/ioutil"
      "encoding/json"
      "github.com/fsnotify/fsnotify"
  )


type Config struct{
        Model string `json:"model"`
        Src string `json:"src"`
        Destination string `json : "destination"`
        //Support string `json : "support"`
    } 

 /* func jsonconfig(file string) Config{

  var jsondata Config
  //json_file := ("link of the file ")

  file, err := ioutil.ReadFile(filename)
   if err != nil {
        fmt.Printf("%s", err)
        os.Exit(0)
    }

    defer file.Body.Close()

    filejson ,err  :=ioutil.ReadAll(file.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(0)
        } 


    //json.Unmarshal([]byte(realdata), &data)
    json.Unmarshal([]byte(filejson), &jsondata)
    fmt.Println(jsondata)
    return 
}*/



func CopyFile(src, dst string) (err error) {
    sfi, err := os.Stat(src)
    if err != nil {
        return
    }
    if !sfi.Mode().IsRegular() {
        // cannot copy non-regular files (e.g., directories,
        // symlinks, devices, etc.)
        return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
    }
    dfi, err := os.Stat(dst)
    if err != nil {
        if !os.IsNotExist(err) {
            return
        }
    } else {
        if !(dfi.Mode().IsRegular()) {
            return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
        }
        if os.SameFile(sfi, dfi) {
            return
        }
    }
    if err = os.Link(src, dst); err == nil {
        return
    }
   // err = copyFileContents(src, dst)
    return
}



  func main() {

     // var src string
      //var dst string
      var jsondata Config
  //json_file := ("link of the file ")

  file, err := os.Open("D:/Gocode/testproject/src/main/configration.json")
   if err != nil {
        fmt.Printf("%s", err)
        os.Exit(0)
    }

    defer file.Close()

    filejson ,err  :=ioutil.ReadAll(file)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(0)
        } 


        json.Unmarshal(filejson, &jsondata)




    fmt.Println(jsondata.Src)
    fmt.Println(jsondata.Destination)
    //src = "D:\\Gocode\\data"
    //dst = "d:/destination"
    //jsondata.Destination = dst

    

      watcher, err := fsnotify.NewWatcher()
      if err != nil {
          log.Fatal(err)
      }
            err = watcher.Add(jsondata.Src)
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
                      //fmt.Printf("%T", event.Name)
                      CopyFile(event.Name, jsondata.Destination)
                      fmt.Println("file copied")

                  }
              case err := <-watcher.Errors:
                  fmt.Println("error:", err)
              }
          }
      }()

     
  }



