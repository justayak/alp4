package main

import (
	"fmt"
	."sync"
	"math/rand"
	"os/exec"
	"time"
	"os"
	."strconv"
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
func (t *Table) check(smoker chan int) (bool, int){
	t.block.RLock()
	take:=<-smoker // Nimm, um zu pruefen..	
	if t.items[0] != None && t.items[0] != take {
		if t.items[1] != None && t.items[1] != take {
			t.block.RUnlock()
			return true, take
		}
	}	
	smoker <- take // Nimms sein item wieder auf..
	t.block.RUnlock()
	return false, take
}



// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// P E R S O N S
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

var fred, fritz, franz chan int 
var fredDone, fritzDone, franzDone chan bool 
var table *Table

func putNextStuffOnTable(table *Table, put chan bool, smoked <-chan bool){	
	for {
		t:=rand.Intn(3)
		switch t{
			case 0:
				fmt.Println("Put [P] & [T]")
				table.putOnTable(Paper, Tobacco)
			case 1:
				fmt.Println("Put [M] & [T]")
				table.putOnTable(Match, Tobacco)
			case 2:
				fmt.Println("Put [M] & [P]")
				table.putOnTable(Match, Paper)
			case 3:
				fmt.Println("lol")
		}
		put <- true // Um alle 3 zu informieren
		<-smoked
		time.Sleep(time.Second)
	}
	
}

// Erzeugt einen neuen Raucher, der mit seinem jeweiligen 
// Standardwert gefuellt wird
func Raucher(n int, defValue int) (chan int, chan bool){
	raucher:=make(chan int, n)
	smokerDone:=make(chan bool)
	for i:=0;i<n;i++ {
		raucher <- defValue
	}
	return raucher, smokerDone
}

func FanIn(a,b,c chan bool) <- chan bool {
	result:=make(chan bool)
	go func() {
		for {
			select {
				case <-a:
					result <- true
				case <-b:
					result <- true
				case <-c:
					result <- true	
			}
		}
	}()
	return result
}

func fanOut(d chan bool) (chan bool,chan bool,chan bool){
	a:= make (chan bool,1)
	b:= make (chan bool,1)
	c:= make (chan bool,1)
	go func() {
		for {
			<-d
			a <- true
			b <- true
			c <- true
		}
	}()
	return a,b,c
}

func SimulateSmoker(smoker chan int, smoked chan bool, table *Table, smokedSignal chan bool, done chan bool){	
	timeout:= time.After(50*time.Second)
	for {
		if len(smoker) == 0 {
			fmt.Println("Ich hab nix mehr!")
			done<-true
			break
		}
		select {
			case <-smokedSignal:	
				repair:=len(smoker)
				isChecked, taken := table.check(smoker)
				if isChecked{
					if len(smoker) != repair {
						smoker <- taken // scheiss-hack..
					}
					table.use(smoker)
					print()
					smoked<- true
				}				
				
			case <-timeout:
				fmt.Println("Ich warte zu lange!")
				done<-true
		}
	}	
}

// rendered den krams
func print(){
	// table.block.Lock()
	r:= "Fred: \t"
	for i:=0;i<len(fred);i++{
		r+= "[P]"
	}
	
	r+= "\nFritz: \t"
	for i:=0;i<len(fritz);i++{
		r+= "[T]"
	}
	
	r+= "\nFranz: \t"
	for i:=0;i<len(franz);i++{
		r+= "[M]"
	}
	
	totalCount := len(fred) + len(fritz) + len(franz)
	r+= "\ntotal: " 
	r+=Itoa(totalCount)
	
	cmd:=exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()		
	fmt.Println(r)
	// table.block.Unlock()
}

func main(){
	n:=10
	fred, fredDone = Raucher(n,Paper)	
	fritz, fritzDone = Raucher(n,Tobacco)	
	franz, franzDone = Raucher(n,Match)		
	// die Raucher sind bereit
	
	print()
	
	putOnTableSignal := make(chan bool,1)
	
	table = NewTable()
	
	go putNextStuffOnTable(table,putOnTableSignal, FanIn(fredDone,fritzDone,franzDone))
	
	fredNoti,fritzNoti,franzNoti := fanOut(putOnTableSignal)
	
	done:=make(chan bool)
	
	go SimulateSmoker(fred, fredDone, table, fredNoti, done)
	go SimulateSmoker(fritz,fritzDone, table, fritzNoti, done)
	go SimulateSmoker(franz,franzDone, table, franzNoti, done)
	
	<-done
	fmt.Println("Wir gehen raus..")
}

