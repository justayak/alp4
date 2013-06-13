package yulib

import (
	"fmt"
	"math/rand"
	"time"
	"container/heap"
	)

type Node struct{
	Data int 
	Children []Node
}

func (n *Node) Child(c Node){
	temp := make([]Node, len(n.Children)+1)
	for i:= range n.Children {
		temp[i]=n.Children[i]
	}
	temp[len(n.Children)] = c
	n.Children = temp
}

func (n Node) size() int{
	count := 1
	for i:= range n.Children {
		child := n.Children[i]
		count += child.size()
	}	
	return count
}

func (n *Node) Print(){
	fmt.Print(" <(")
	fmt.Print(n.Data)
	fmt.Print(") [")
	for i:= range n.Children {
		child := n.Children[i]
		child.Print()
	}
	fmt.Print("]>  ")
}

type Tree struct{
	Root Node
}

func (t Tree) Size() int{
	return t.Root.size() -1
}

func (t Tree) Print(){
	t.Root.Print()
	fmt.Println("--")
}

var seedSalt int64 = 0
func RandInt(min int, max int) int {	
	rand.Seed( time.Now().UTC().UnixNano() + seedSalt)
	seedSalt += 1
	return  min + rand.Intn(max-min)
}
