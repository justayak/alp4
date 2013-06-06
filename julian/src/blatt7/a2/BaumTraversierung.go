package main

import "fmt"
import "yulib"

func createTree(flip bool) yulib.Tree {
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

type NodeContent struct{
	Data int 
	Level int
	IsEnd bool
}

func (n NodeContent) equals(other NodeContent) bool {
	return n.Data == other.Data && n.Level == other.Level && n.IsEnd == other.IsEnd
}

// Traversiert einen Baum und schreibt die Kinder
// in den Channel c
func traverse(tree yulib.Tree, c chan NodeContent){	
	traverseRec(tree.Root, c, 0)	
	endSignal := NodeContent{}
	endSignal.IsEnd = true 
	c<-endSignal
	close(c)
}

func traverseRec(node yulib.Node, c chan NodeContent, treeDepth int ){
	current := NodeContent{}
	current.Data = node.Data
	current.Level = treeDepth
	current.IsEnd = false
	c<-current	
	for i:= range node.Children {
		child := node.Children[i]
		traverseRec(child, c, treeDepth + 1)
	}	
}

// Vergleicht zwei Baum-Channels
func compare(a,b chan NodeContent, result chan bool){
	
	for {		
		left := <- a
		right := <- b
		
		if left.equals(right) {
			if left.IsEnd {
				result <- true
				close(result)
				break
			}
		} else {
			result <- false
			close(result)
			break
		}
	
	}
	
	
}

func main() {    
	
	// 2x der gleiche Baum
	tree1 := createTree( true )
	//tree2 := createTree( true ) 
	
	
	
	// 1x ein anderer Baum
	tree3 := createTree( false )
	
	fmt.Println(EqualTrees(tree1,tree3))
	
	fmt.Println("End")
}

// Test-Funktion
func EqualTrees(a yulib.Tree, b yulib.Tree) bool {
	if a.Size() != b.Size() { return false }
	
	left := make(chan NodeContent)
	right := make(chan NodeContent)
	result := make(chan bool)
	
	
	go traverse(a, left)
	go traverse(b, right)
	go compare(left,right,result)
	
	e := <-result	
	return e
	
}