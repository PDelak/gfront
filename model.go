package main
import "fmt"

type OwningType struct
{
  owner bool
}

type FileHandle struct
{
  OwningType 
  id int
}

type UniqueLifecycle interface
{
  IsOwner() bool
  Destructor()
}

func NewFileHandle(id int) FileHandle {
  fmt.Printf("open file\n")
  return FileHandle{OwningType:OwningType{true}, id:id}
}

func Move(src *FileHandle) FileHandle {
  src.owner = false
  var res = *src
  res.owner = true
  return res
}

func (handle FileHandle) Destructor() {
  handle.id = 0
  fmt.Printf("close file\n")
}

func(handle FileHandle) IsOwner() bool {
  return handle.owner
}


func ScopeExit(res UniqueLifecycle) {
  if(res.IsOwner()) {
    res.Destructor()
  }  
}

func main() {
   {
   src := NewFileHandle(1)
   dst := NewFileHandle(2)
   fmt.Printf("src : %d, %d\n", src.id, src.OwningType.owner)
   fmt.Printf("dst : %d, %d\n", dst.id, dst.OwningType.owner)
   ScopeExit(src)
   ScopeExit(dst)
   }
   {
   src := NewFileHandle(1)
   alias := src
   dst := Move(&src)
   fmt.Printf("alias : %d, %d\n", alias.id, alias.OwningType.owner)
   fmt.Printf("src : %d, %d\n", src.id, src.OwningType.owner)
   fmt.Printf("dst : %d, %d\n", dst.id, dst.OwningType.owner)
   ScopeExit(src)
   ScopeExit(dst)
   }
}
