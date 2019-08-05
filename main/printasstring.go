package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	

)


func main(){
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
		if err != nil{
			fmt.Println(err)
		} 
	
   // defer.res.Body.Close() 
	bytes,erro :=ioutil.ReadAll(res.Body)

	  if erro != nil{
	  	fmt.Println(erro)
	  }
    string_body:= (bytes)
    fmt.Printf("%s",string_body)
}