package main
import(
	"fmt"
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
//~~~~~~~~~~~~~~~~~~~~~
// T A B L E
//~~~~~~~~~~~~~~~~~~~~~

func putOnTable() <-chan int {
	result:=make(chan int)
	go func() {
		for {
			t:=rand.Intn(3)
			switch t{
				case 0:
					result<-Paper
					result<-Tobacco
				case 1:
					result<-Match
					result<-Tobacco
				case 2:
					result<-Match
					result<-Paper
				case 3:
					fmt.Println("lol")
			}
		}
	}()
	return result
}

//~~~~~~~~~~~~~~~~~~~~~~
// R A U C H E R
//~~~~~~~~~~~~~~~~~~~~~~

type Smoker struct {
	stuff int 
	deposit chan int
}


// Erzeugt einen neuen Raucher, der mit seinem jeweiligen 
// Standardwert gefuellt wird
func Raucher(n int, defValue int) *Smoker{
	raucher:=make(chan int, n)
	for i:=0;i<n;i++ {
		raucher <- defValue
	}
	
	result:=new(Smoker)
	result.stuff = defValue
	result.deposit = raucher
	return result
}

func (smoker *Smoker)simulate(table <-chan int){
	for {
		first:=<-table
		second:=<-table
		if smoker.stuff != first && smoker.stuff != second {
			// smoke... ** paff paff **
			<-smoker.deposit
			time.Sleep(time.Second)
			print()
		}
	}
}

//~~~~~~~~~~~~~~~~~~~~~~
// M A I N
//~~~~~~~~~~~~~~~~~~~~~~
var fred, fritz, franz *Smoker 

func main() {
	n:=10
	fred = Raucher(n,Paper)	
	fritz = Raucher(n,Tobacco)	
	franz = Raucher(n,Match)		
	
	table := putOnTable()
	
	fredIn, fritzIn, franzIn := fanOut(table)
	
	print()
	
	go fred.simulate(fredIn)
	go fritz.simulate(fritzIn)
	go franz.simulate(franzIn)
	
	duration:=time.Duration(10) * time.Second
	time.Sleep(duration)
	
	fmt.Println("Der Laden macht jetzt dicht blabla verp*sst euch!")
}

func print(){
	// table.block.Lock()
	r:= "Fred: \t"
	for i:=0;i<len(fred.deposit);i++{
		r+= "[P]"
	}
	
	r+= "\nFritz: \t"
	for i:=0;i<len(fritz.deposit);i++{
		r+= "[T]"
	}
	
	r+= "\nFranz: \t"
	for i:=0;i<len(franz.deposit);i++{
		r+= "[M]"
	}
	
	totalCount := len(fred.deposit) + len(fritz.deposit) + len(franz.deposit)
	r+= "\ntotal: " 
	r+=Itoa(totalCount)
	
	cmd:=exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()		
	fmt.Println(r)
}

//~~~~~~~~~~~~~~~~~~~~~~~
// U T I L S
//~~~~~~~~~~~~~~~~~~~~~~~ 

func FanIn(a,b,c chan int) <- chan int {
	result:=make(chan int)
	go func() {
		for {
			select {
				case v:=<-a:
					result <- v
				case v:=<-b:
					result <- v
				case v:=<-c:
					result <- v	
			}
		}
	}()
	return result
}

func fanOut(d <-chan int) (<-chan int,<-chan int,<-chan int){
	a:= make (chan int,1)
	b:= make (chan int,1)
	c:= make (chan int,1)
	go func() {
		for {
			v:=<-d
			a <- v
			b <- v
			c <- v
		}
	}()
	return a,b,c
}