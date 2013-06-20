package main

import (
	"fmt"
	."sync"
	"math/rand"
)

const (
	None = iota
	Paper = iota
	Tobacco = iota
	Match = iota
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// T A B L E (monitor)
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Table struct {
	block RWMutex
	items [2]int
}

// Erzeugt neuen Table
func NewTable() *Table {
	result:=new(Table)
	result.items[0] = None
	result.items[1] = None
	return result
}

// Monitor - Function
func (t *Table) putOnTable(valueA int, valueB int) {
	t.block.Lock()
	t.items[0] = valueA
	t.items[1] = valueB
	t.block.Unlock()
}

// Monitor - Function
func (t *Table) use(smoker chan int){
	t.block.Lock()
	<-smoker
	t.items[0] = None
	t.items[1] = None
	// smoke... **paff paff paff**
	t.block.Unlock()
}

// Monitor - Function
func (t *Table) check(smoker chan int) bool{
	t.block.RLock()
	take:=<-smoker // Nimm, um zu pruefen..	
	if t.items[0] != None && t.items[0] != take {
		if t.items[1] != None && t.items[1] != take {
			t.block.RUnlock()
			return true
		}
	}	
	smoker <- take // Nimms sein item wieder auf..
	t.block.RUnlock()
	return false
}



// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// P E R S O N S
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

var fred, fritz, franz chan int 
var table *Table

func putNextStuffOnTable(table *Table, put chan bool){
	t:=rand.Intn(3)
	switch t{
		case 0:
			table.putOnTable(Paper, Tobacco)
		case 1:
			table.putOnTable(Match, Tobacco)
		case 2:
			table.putOnTable(Match, Paper)
	}
	put <- true
}

// Erzeugt einen neuen Raucher, der mit seinem jeweiligen 
// Standardwert gefuellt wird
func Raucher(n int, defValue int) chan int{
	raucher:=make(chan int, n)
	for i:=0;i<n;i++ {
		raucher <- defValue
	}
	return raucher
}

func main(){
	n:=20
	fred = Raucher(n,Paper)	
	fritz = Raucher(n,Tobacco)	
	franz = Raucher(n,Match)		
	// die Raucher sind bereit
	
	table = NewTable()
	
	fmt.Println("hallo")
}

