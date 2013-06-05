package main

import "fmt"
import "yulib"

func makeTree(flip bool) yulib.Tree {
	tree := yulib.Tree{}	
	child1 := yulib.Node{}
	child1.Data = 1	
	child2 := yulib.Node{}
	child2.Data = 2	
	child3 := yulib.Node{}
	child3.Data = 3	
	child4 := yulib.Node{}
	child4.Data = 4	
	child5 := yulib.Node{}
	child5.Data = 5	
	child6 := yulib.Node{}
	if flip {
		child6.Data = 6	
	}else{
		child6.Data = 8
	}	
	child7 := yulib.Node{}
	child7.Data = 7	
	
	child3.Child(child7)
	child5.Child(child6)
	child2.Child(child4)
	child2.Child(child5)	
	tree.Root.Child(child1)
	tree.Root.Child(child2)
	tree.Root.Child(child3)
	return tree
}

func main() {    
	
	// 2x der gleiche Baum
	tree1 := makeTree(true)
	tree2 := makeTree(true) 
	
	// 1x ein anderer Baum
	tree3 := makeTree(false)
	
	
	
	tree1.Print()
	tree2.Print()
	tree3.Print()
	
	fmt.Println("End")
}

// Test-Funktion
func EqualTrees(a yulib.Tree, b yulib.Tree) bool {
	if a.Size() != b.Size() { return false }
	
	
	
	return false
	
}