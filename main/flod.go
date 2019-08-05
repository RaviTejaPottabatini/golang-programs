package main 


import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
)


func copy(src, dst string) (int64, error) {
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


func main(){


	  sourceFile:= ""
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