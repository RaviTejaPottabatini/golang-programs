package main

  import (
      "log"
      "fmt"
      "os"
      "io"
      "path/filepath"
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
  in, err := os.Open(src)
  if err != nil {
    return
  }
  defer in.Close()

  out, err := os.Create(dst)
  if err != nil {
    return
  }
  defer func() {
    if e := out.Close(); e != nil {
      err = e
    }
  }()

  _, err = io.Copy(out, in)
  if err != nil {
    return
  }

  err = out.Sync()
  if err != nil {
    return
  }

  si, err := os.Stat(src)
  if err != nil {
    return
  }
  err = os.Chmod(dst, si.Mode())
  if err != nil {
    return
  }

  return
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string) (err error) {
  src = filepath.Clean(src)
  dst = filepath.Clean(dst)

  si, err := os.Stat(src)
  if err != nil {
    return err
  }
  if !si.IsDir() {
    return fmt.Errorf("source is not a directory")
  }

  _, err = os.Stat(dst)
  if err != nil && !os.IsNotExist(err) {
    return
  }
  if err == nil {
    return fmt.Errorf("destination already exists")
  }

  err = os.MkdirAll(dst, si.Mode())
  if err != nil {
    return
  }

  entries, err := ioutil.ReadDir(src)
  if err != nil {
    return
  }

  for _, entry := range entries {
    srcPath := filepath.Join(src, entry.Name())
    dstPath := filepath.Join(dst, entry.Name())

    if entry.IsDir() {
      err = CopyDir(srcPath, dstPath)
      if err != nil {
        return
      }
    } else {
      // Skip symlinks.
      if entry.Mode()&os.ModeSymlink != 0 {
        continue
      }

      err = CopyFile(srcPath, dstPath)
      if err != nil {
        return
      }
    }
  }

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
                      CopyDir(jsondata.Src, jsondata.Destination)
                      fmt.Println("file copied")

                  }
              case err := <-watcher.Errors:
                  fmt.Println("error:", err)
              }
          }
      }()

     
  }



