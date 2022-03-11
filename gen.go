package main

import "fmt"
import "gfront/runtime"

func NewFileHandle(id int) FileHandle {
	fmt.Printf("open file\n")
	return FileHandle{runtime.OwningType{true}, id}
}

func MoveFileHandle(src *FileHandle, dst *FileHandle) {
	src.Owner = false
	dst.Owner = true
}

func (handle *FileHandle) Destructor() {
	handle.id = 0
	fmt.Printf("close file\n")
}

func (handle *FileHandle) IsOwner() bool {
	return handle.Owner
}

func ScopeExit(res runtime.UniqueLifecycle) {
	if res.IsOwner() {
		res.Destructor()
	}
}

func NewNode(value int, node *Node) Node {
	fmt.Printf("Node %d\n", value)
	return Node{runtime.OwningType{true}, value, node}
}

func NewNodeHeap(value int, node *Node) *Node {
	fmt.Printf("Node %d\n", value)
	return &Node{runtime.OwningType{true}, value, node}
}

func MoveNode(src *Node, dst *Node) {
	src.Owner = false
	dst.Owner = true
}

func (handle *Node) Destructor() {
	handle = nil
	fmt.Printf("release node %d %d\n", handle.value, handle.Owner)
}

func (handle Node) IsOwner() bool {
	return handle.Owner
}
