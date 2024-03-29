package main

import (
  "fmt"
  "io"
  "log"
  "os"
  "io/ioutil"
  "path/filepath"
)

var target_path string
var source_path string

func main() {

  source_path = "D:/rapiscan/input/632DV"
  target_path = "D:/alredyexists"

  err := CopyDir(source_path, target_path)
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Print("copy finish")
  }

}











func CopyFile(src, dst string) (err error) {
  in, err := os.Open(src)
  if err != nil {
    return
  }
  defer in.Close()

  out, err := os.Create(dst)
  if err != nil {
    return
  }
  defer func() {
    if e := out.Close(); e != nil {
      err = e
    }
  }()

  _, err = io.Copy(out, in)
  if err != nil {
    return
  }

  err = out.Sync()
  if err != nil {
    return
  }

  si, err := os.Stat(src)
  if err != nil {
    return
  }
  err = os.Chmod(dst, si.Mode())
  if err != nil {
    return
  }

  return
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string) (err error) {
  src = filepath.Clean(src)
  dst = filepath.Clean(dst)

  si, err := os.Stat(src)
  if err != nil {
    return err
  }
  if !si.IsDir() {
    return fmt.Errorf("source is not a directory")
  }

  _, err = os.Stat(dst)
  if err != nil && !os.IsNotExist(err) {
    return
  }
  if err == nil {
    return fmt.Errorf("destination already exists")
  }

  err = os.MkdirAll(dst, si.Mode())
  if err != nil {
    return
  }

  entries, err := ioutil.ReadDir(src)
  if err != nil {
    return
  }

  for _, entry := range entries {
    srcPath := filepath.Join(src, entry.Name())
    dstPath := filepath.Join(dst, entry.Name())

    if entry.IsDir() {
      err = CopyDir(srcPath, dstPath)
      if err != nil {
        return
      }
    } else {
      // Skip symlinks.
      if entry.Mode()&os.ModeSymlink != 0 {
        continue
      }

      err = CopyFile(srcPath, dstPath)
      if err != nil {
        return
      }
    }
  }

  return
}
