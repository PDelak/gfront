package main

import "fmt"
import "gfront/runtime"


func NewFileHandle(id int) FileHandle {
  fmt.Printf("open file\n")
  return FileHandle{runtime.OwningType{true}, id}
}

func Move(src *FileHandle) FileHandle {
  src.Owner = false
  var res = *src
  res.Owner = true
  return res
}

func (handle FileHandle) Destructor() {
  handle.id = 0
  fmt.Printf("close file\n")
}

func(handle FileHandle) IsOwner() bool {
  return handle.Owner
}


func ScopeExit(res runtime.UniqueLifecycle) {
  if(res.IsOwner()) {
    res.Destructor()
  }
}

