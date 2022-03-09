package main

import "fmt"
import "gfront/runtime"
import "gfront/parser"

type FileHandle struct {
	runtime.OwningType
	id int
}

type Node struct {
	runtime.OwningType
	value int
	next  *Node
}

func pushFront(value int, node *Node) *Node {
	newHead := NewNodeHeap(value, node)
	return newHead
}

func passNodeUniquely(node *Node) Node {
	newNode := MoveNode(node)
	return newNode
}

func visit(node *Node) {
	var curr *Node
	curr = node
	for {
		if curr == nil {
			break
		}
		fmt.Printf("\n%d", curr.value)
		curr = curr.next
	}
	fmt.Printf("\n")
}

func testGen() {
	{
		src := NewFileHandle(1)
		dst := NewFileHandle(2)
		fmt.Printf("src : %d, %d\n", src.id, src.OwningType.Owner)
		fmt.Printf("dst : %d, %d\n", dst.id, dst.OwningType.Owner)
		ScopeExit(src)
		ScopeExit(dst)
	}
	{
		src := NewFileHandle(1)
		alias := src
		dst := MoveFileHandle(&src)
		fmt.Printf("alias : %d, %d\n", alias.id, alias.OwningType.Owner)
		fmt.Printf("src : %d, %d\n", src.id, src.OwningType.Owner)
		fmt.Printf("dst : %d, %d\n", dst.id, dst.OwningType.Owner)
		ScopeExit(src)
		ScopeExit(dst)
	}
	{
		src := NewNode(1, nil)
		alias := src
		dst := MoveNode(&src)
		fmt.Printf("alias : %d, %d\n", alias.value, alias.OwningType.Owner)
		fmt.Printf("src : %d, %d\n", src.value, src.OwningType.Owner)
		fmt.Printf("dst : %d, %d\n", dst.value, dst.OwningType.Owner)
		ScopeExit(src)
		ScopeExit(dst)

	}
	{
		src := NewNode(1, nil)
		dst := MoveNode(&src)
		anotherNode := passNodeUniquely(&dst)
		fmt.Printf("src : %d, %d\n", src.value, src.OwningType.Owner)
		fmt.Printf("dst : %d, %d\n", dst.value, dst.OwningType.Owner)
		fmt.Printf("anotherNode : %d, %d\n", anotherNode.value, anotherNode.OwningType.Owner)
		ScopeExit(src)
		ScopeExit(dst)
	}
	{
		head1 := pushFront(1, nil)
		head2 := pushFront(2, head1)
		head3 := pushFront(3, head2)
		visit(head3)
		ScopeExit(head3)
		ScopeExit(head2)
		ScopeExit(head1)
	}
	{
		src := NewNodeHeap(5, nil)
		dst := MoveNode(src)
		fmt.Printf("src : %d, %d\n", src.value, src.OwningType.Owner)
		fmt.Printf("dst : %d, %d\n", dst.value, dst.OwningType.Owner)
		ScopeExit(src)
		ScopeExit(dst)
	}
}

func main() {
	testGen()
	parser.Parse("template")
}
