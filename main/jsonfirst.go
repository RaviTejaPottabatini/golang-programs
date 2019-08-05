package main

import (
    "fmt"
    "net/http"   
    "io/ioutil"
    "os"
    "encoding/json"
    )


   type User struct{
       Userid int
       Id int
       Title string
       Completed bool
    } 
     

func main() {
    response,  err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } 

        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
       fmt.Println(string(contents))
       var s User
       json.Unmarshal([]byte(contents) , &s)

       fmt.Printf("user Id: %d\n", s.Userid)
       fmt.Printf("id : %d\n" , s.Id) 
       fmt.Printf("title : %s\n",s.Title )
       fmt.Println(s)

        
    

}