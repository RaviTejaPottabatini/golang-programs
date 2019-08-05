package main

import (
	"log"
	"os"
	"io/ioutil"
	"fmt"
	"io"
)

func slowcp(src ,dst string)error{
	srcfile ,err := os.Open(src)
	if err != nil{
		return fmt.Errorf("error %v",err)
	}

	defer srcfile.Close()

	ds ,err := ioutil.Readll(srcfile)
	if err != nil{
		return fmt.Errorf("error %v",err)

	}

	//

	dstfile ,err := os.Create(dst)
	if err != nil{
		return fmt.Errorf("error %v",err)

	}
	defer dstfile.Close()

	_,err = dstfile.Write(ds)
	if err != nil{
		return fmt.Errorf("error %v",err)

	}



}


func  cp(src ,dst string)error{
	srcfile ,err := os.Open(src)
	if err != nil{
		return fmt.Errorf("error %v",err)
	}

	defer srcfile.Close()


	//

	dstfile ,err := os.Create(dst)
	if err != nil{
		return fmt.Errorf("error %v",err)

	}
	defer dstfile.Close()

	_,err := io.Copy(dstfile,srcfile)
	if err!= nil{
		return fmt.Errorf("error", err)
	}

	return nil

}

func main(){

	src := os.Args[1]
	dst := os.Args[2]
	
	//


	err := cp(src,dst)
	if er!= nil{
		log.Fatalln(err)
	}

}