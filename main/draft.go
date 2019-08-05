package main

import "fmt"
import"time"
import"log"
//import"os"
import"io/ioutil"

func main () {





file1 := "d:/Gocode/testproject"
//V := time.Now()
//T := V.Add(-10*time.Second)

 for {

  V := time.Now()

  T := V.Add(-10*time.Second)

    

       files, err := ioutil.ReadDir(file1)
    if err != nil {
        log.Fatal(err)
    }


   for _, h := range files {

  

 //fmt.Println(h.Name())
  long := h.ModTime() 
  diff := long.Sub(T)

  if diff == (time.Duration(0)* time.Second ){
  	fmt.Printf("%s it cant be posiible%s \n",h.Name(),T)
    break
  }

 //fmt.Println(T)
  if diff > (time.Duration(0)* time.Second) {
       fmt.Printf("%s a new file added in direatory at %s \n", h.Name(),T)
       break
                 
	       }
   }
   time.Sleep(7 *time.Second)
}


}