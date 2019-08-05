package main 


import (

    "fmt"
  "log"
  //"encoding/json"
  //"io/ioutil"
  //"io"
  "os"
  "github.com/fsnotify/fsnotify"

)

/*

type Config struct{
        Model string `json:"model"`
        Src string `json:"src"`
        Destination string `json : "destination"`
        Support string `json : "support"`

      
    } 


/*func LoadConfiguration(file string) Config {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}*/

/*

func jsonconfig(file string) Config{

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
    return jsondata
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

  var dst string
  var src string

  

    // json_data , _ := jsonconfig("vndgh.json")


            
      watcher, err := fsnotify.NewWatcher()
      if err != nil {
          log.Fatal(err)
      }

       err = watcher.Add("D:\\rapiscan\\input\\632DV")
      if err != nil {
          log.Fatal(err)
        }

//mailto:rashi.sweet8810@yahoo.com

      defer watcher.Close()

      //done := make(chan bool)
       func() {
          for {
              select {
              case Event := <-watcher.Events:
                  fmt.Println("event: in", Event)
                  if Event.Op&fsnotify.Write == fsnotify.Write {
                      fmt.Println("modified file:", Event.Name)
                      
                        src = Event.Name
                        copy(src , dst)
                        fmt.Println("file is copied")
                      
                  }
              case err := <-watcher.Errors:
                  fmt.Println("error:", err)
              }
          }
      }()    
     
      //<-done
 



 
}














