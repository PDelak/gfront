package main

import "gfront/runtime"
import "gfront/parser"

type FileHandle struct
{
  runtime.OwningType
  id int
}


func main() {
  parser.Parse("template")
}
         
