package main

import (
	"fmt"
)

const (
	Paper = iota
	Tabacco = iota
	Match = iota
)

var fred, fritz, franz <-chan int 
var wirtinErna chan int

// Erzeugt einen neuen Raucher, der mit seinem jeweiligen 
// Standardwert gefuellt wird
func Raucher(n int, defValue int) <-chan int{
	raucher:=make(chan int, n)
	for i:=0;i<n;i++ {
		raucher <- defValue
	}
	return raucher
}

func main(){
	n:=20
	fred = Raucher(n,Paper)	
	fritz = Raucher(n,Tabacco)	
	franz = Raucher(n,Match)	
	//wirtinErna = make(chan int, 0)
	
	//wirtinErna <-1
	
	// die Raucher sind bereit
	
	fmt.Println("hallo")
}
