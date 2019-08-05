package main

import (
    "fmt"
    //"net/http"   
    //"io/ioutil"
   // "os"
    "encoding/json"
    )



/*{
  "birdType": "pigeon",
  "what it does": "likes to perch on rocks"
}*/
 type Bird struct{
      Species string `json:"birdType"`
      Description string `json:"what it does"`
      
    } 

     
/*func link(){
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

       var s User
       json.Unmarshal([]byte(contents) , &s)  
       fmt.Printf("Id: %d", s.Id)
}*/

/*func (l Bird ) String() string{
	return fmt.Sprintf(l.Bird)
}*/

func main() {
  birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
  var birds Bird
  json.Unmarshal([]byte(birdJson), &birds)
  fmt.Println(birds)

 

}