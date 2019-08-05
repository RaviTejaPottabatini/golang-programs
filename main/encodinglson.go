package main


import(
	"encoding/json"
	"fmt"
	//"net/http"
	//"io/ioutil"
)

type Bird struct{
	Species string `json : "birdType"`
	Description string `json: "what is does"`

}

func main(){
	pigeon := &Bird{
		Species : "pigeon",
		Description :"likes to eat seed ",

	}

	data,_ := json.Marshal([]*Bird{pigeon})

	fmt.Println(string(data))
}