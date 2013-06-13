package main 
//import ("math/rand"; "time"; "fmt")
import ("fmt"; ."queue";)

const (
	BIG = 4
	MIDDLE = 2
	SMALL = 1
)

type Package struct {
	Value int
}

type Position struct {
	
}

func main () {
	
	queue := Queue{}
	queue.Enq("lol")
	queue.Enq("dummkopf")
	
	fmt.Printf ("size %d\n", queue.Size()) 
	
	elem := queue.Deq().(string)
	
	fmt.Printf(elem)
}