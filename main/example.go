package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	

)





type User struct{
	Userid []int  `json: "userId"`
}



func (e userid ) int () int{
	return fmt.Sprintf(e.userid)
}

func getdata(){
 
    url := "https://jsonplaceholder.typicode.com/todos/1"

	res,err := http.Get(url)
		if err != nil{
			panic(err.Error())
		} 
	defer res.Body.Close()

	bytes,err :=ioutil.ReadAll(res.Body)
	  if err != nil{
	  	panic(err.Error())

	  }

	  fmt.Println(bytes)

	  var s User
	  json.Unmarshal(bytes , &s )
	  fmt.Println(s)

    
}
func main(){


	getdata()


}