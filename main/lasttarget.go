package main 


import (

    "fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"os"
	"github.com/fsnotify/fsnotify"

)



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
}

func Copy(src, dst string) (int64, error) {
        sourceFileStat, err := os.Stat(src)
        if err != nil {
                return 0, err
        }

        if !sourceFileStat.Mode().IsRegular() {
                return 0, fmt.Errorf("%s is not a regular file", src)
        }

        source, err := os.Open(src)
        if err != nil {
                return 0, err
        }
        defer source.Close()

        destination, err := os.Create(dst)
        if err != nil {
                return 0, err
        }
        defer destination.Close()
        nBytes, err := io.Copy(destination, source)
        return nBytes, err
}



func main() {

	

     json_data , _ := LoadConfiguration("vndgh.json")


            
      watcher, err := fsnotify.NewWatcher()
      if err != nil {
          log.Fatal(err)
      }

       err = watcher.Add(json_data)
      if err != nil {
          log.Fatal(err)
        }



      defer watcher.Close()

      //done := make(chan bool)
       func() {
          for {
              select {
              case Event := <-watcher.Events:
                  fmt.Println("event: in", Event)
                  if Event.Op&fsnotify.Write == fsnotify.Write {
                      fmt.Println("modified file:", Event.Name)
                  }
              case err := <-watcher.Errors:
                  fmt.Println("error:", err)
              }
          }
      }()    
     
      //<-done
 
      sourceFile:= Event.Name
	  destinationFile:=""

	  input, err := ioutil.ReadFile(sourceFile)
        if err != nil {
                fmt.Println(err)
                return
        }

        err = ioutil.WriteFile(destinationFile, input, 0644)
        if err != nil {
                fmt.Println("Error creating", destinationFile)
                fmt.Println(err)
                return
        }
 



 
}














