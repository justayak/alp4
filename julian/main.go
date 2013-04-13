package main

import (
	"fmt"
	"runtime"
)

func add(x int, y int) int {
    return x + y
}

type Vertex struct {
    X int
    Y int
}

func main() {

    const x = 10
    
    var create = func()[x]int{
        return [x]int{1,2,3,4,5}
    }
	
	var os = runtime.GOOS
	
	fmt.Println(os)
	
	
	m := make(map[string]int)
	m["Answer"] = 42
    
    
    var liste = create()
    
    for i:=0; i < 5; i++ {
    	fmt.Println(liste[i])
    }
	
	fmt.Println(Vertex{1, 2})
    
}
