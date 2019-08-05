package main


import (

	"fmt"
	//"io"
	//"log"
	"os"
)


func CopyFile(src, dst string) (err error) {
    sfi, err := os.Stat(src)
    if err != nil {
        return
    }
    if !sfi.Mode().IsRegular() {
        // cannot copy non-regular files (e.g., directories,
        // symlinks, devices, etc.)
        return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
    }
    dfi, err := os.Stat(dst)
    if err != nil {
        if !os.IsNotExist(err) {
            return
        }
    } else {
        if !(dfi.Mode().IsRegular()) {
            return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
        }
        if os.SameFile(sfi, dfi) {
            return
        }
    }
    if err = os.Link(src, dst); err == nil {
        return
    }
   // err = copyFileContents(src, dst)
    return
}


func main(){


	var src string
	var dst string

	src = "D:/rapiscan/input/632DV/where.txt"
	dst = "D:/rapiscan/output/632DV/hie.txt"

	CopyFile(src ,dst)


} 