package main 
//import ("math/rand"; "time"; "fmt")
import ("fmt"; ."queue"; "yulib";)

const (
	BIG = 4
	MIDDLE = 2
	SMALL = 1
)

type Object struct {
	Value int
}

// zum versenden der Elemente
type Package struct {
	Elements [] Object
}

type Position struct {
	
}

// Erzeugt zufälliges Objekt
func generateObject() Object{
	s := yulib.RandInt(0, 10)
	v := 0
	switch {
		case s < 5:  v = SMALL
		case s >= 5 && s < 8: v = MIDDLE
		case s >= 8: v = BIG
	}
	result := Object{
		Value : v,
	}
	return result
}

func main () {
	
	queue := Queue{}
	
	for i:= 0; i < 10; i++ {
		queue.Enq(generateObject())
	}
	
	fmt.Printf ("size %d\n", queue.Size()) 
	
	elem := queue.Deq().(Object)
	
	fmt.Printf("v %d\n", elem.Value)
}