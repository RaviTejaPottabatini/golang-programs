package main

import (
    "fmt"
    "io/ioutil"
     //"log"
     "time"
     //"strings"
     //"reflect"
)


func diff(){

   // var f func()
    var h func()
    var t *time.Timer

    
 h = func () {
         files, err := ioutil.ReadDir("d:/Gocode/testproject")
    if err != nil {
        panic(err)
    }





    for _, h := range files {
           // fmt.Println(h.Name())
            str := h.ModTime()

            T := time.Now()
           
           t := T.Add(-10*time.Second)
           
            
            //Strarray := strings.Fields(str)
           if str != t {
            fmt.Println("ok")
           }



            //fmt.Println(h.ModTime())
    }
        t = time.AfterFunc(time.Duration(1) * time.Second, h)
    }


  /*  f = func () {
         files, err := ioutil.ReadDir("d:/Gocode/testproject")
    if err != nil {
        panic(err)
    }
    fmt.Println(files)
    


    for _, f := range files {
            //fmt.Println(f.Name())
             var F2 = f.ModTime()
           // var F1 = h
            f2 := F2
           // f1 := &F1
           fmt.Println(f2)
            //fmt.Println(f.ModTime())
    }
        t = time.AfterFunc(time.Duration(5) * time.Second, f)
    
    */      
  //}

    t = time.AfterFunc(time.Duration(1)* time.Second,h)
  //  t = time.AfterFunc(time.Duration(1) * time.Second, f)

 
    defer t.Stop()
    //simulate doing stuff
    time.Sleep(2* time.Second)
}



func main() {
    
   diff()
  

}