package main

import (
    "fmt"
    "net/http"   
    "io/ioutil"
    "os"
    "encoding/json"
    )


    type Givenjson struct{
    	Userid int `json: "userId"`
    	Id int `json: "id"`
    	Title string `json: "title"`
    	Completed  bool `json: "completed"`
    }



    func main(){

    	response,  err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
        if err != nil {
        fmt.Printf("%s", err)
        os.Exit(0)
    }


       defer response.Body.Close()
        realdata, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(0)
        }

    fmt.Println(string(realdata))

     var data Givenjson
     json.Unmarshal([]byte(realdata), &data)
     fmt.Println(data.Title)


     for _, v := range data.Title {

        x := string(v)
     fmt.Printf("%s\n", x)
     }
    
  
 }   