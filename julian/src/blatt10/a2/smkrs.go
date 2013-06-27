package main
import(
	"fmt"
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


//~~~~~~~~~~~~~~~~~~~~~~
// R A U C H E R
//~~~~~~~~~~~~~~~~~~~~~~
// Erzeugt einen neuen Raucher, der mit seinem jeweiligen 
// Standardwert gefuellt wird
func Raucher(n int, defValue int) ( chan int , chan bool ){
	raucher:=make(chan int, n)
	smokerDone:=make(chan bool)
	for i:=0;i<n;i++ {
		raucher <- defValue
	}
	return raucher, smokerDone
}


//~~~~~~~~~~~~~~~~~~~~~~
// M A I N
//~~~~~~~~~~~~~~~~~~~~~~
var fred, fritz, franz chan int 
var fredDone, fritzDone, franzDone chan bool 

func main() {
	n:=10
	fred, fredDone = Raucher(n,Paper)	
	fritz, fritzDone = Raucher(n,Tobacco)	
	franz, franzDone = Raucher(n,Match)		
	
	fmt.Println("q")
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

func fanOut(d chan int) (<-chan int,<-chan int,<-chan int){
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