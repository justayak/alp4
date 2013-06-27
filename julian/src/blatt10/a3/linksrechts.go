package main
import (
	"fmt"
	"yulib"
	"time"
)

func main() {
	fmt.Println("hallo")
	leftPending = createPending(N)
	rightPending = createPending(N)
	crit := NewCriticalRoad()
	render := crit.simulate(leftPending,rightPending)
	go draw(render)
	duration:=time.Duration(10) * time.Second
	time.Sleep(duration)
}

// rendert den krams blabla
func draw(render <-chan string ){
	for {
		o := <- render
		fmt.Println(o)
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// I M P L E M E N T A T I O N
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~

const (
	N = 5
	Left = iota 
	Right = iota 
)

// Warteschlangen für die jeweiligen Seiten
var leftPending, rightPending <-chan bool

// Stellt den kritischen Abschnitt dar
type CriticalRoad struct {
	PreferLeft bool 
	HasValue bool
	Direction int
}

// Erzeugt einen neuen kritischen Abschnitt
func NewCriticalRoad() *CriticalRoad {
	result:= new( CriticalRoad )
	result.PreferLeft = true 
	result.HasValue = false
	return result
}

func (road *CriticalRoad)simulate(left,right <-chan bool )<- chan string {
	result:= make( chan string )
	go func() {
		for {
			l:=""
			m:="[__]"
			r:=""
			if road.HasValue {
				road.HasValue = false 
				road.PreferLeft = !road.PreferLeft
				if road.Direction == Right {
					m = "[_>]"
				} else {
					m = "[<_]"
				}				
			} else {
				if road.PreferLeft {
					if len( left ) > 0 {
						<-left
						road.HasValue = true 
						road.Direction = Right
					} else if len ( right ) > 0 {
						<-right
						road.HasValue = true 
						road.Direction = Left
					}
				}else {
					if len( right ) > 0 {
						<-right
						road.HasValue = true 
						road.Direction = Right
					} else if len ( left ) > 0 {
						<-left
						road.HasValue = true 
						road.Direction = Right
					}
				}
			}
			for i:=0;i<len(left);i++{
				l+="[>]"
			}
			
			for i:=0;i<len(right);i++{
				r+="[<]"
			}	
			rs := l + m + r
			result <- rs
		}
	}()
	return result
}

// Erzeugt die Strasse VOR dem kritischen
// Abschnitt... blabla
func createPending(n uint) <- chan bool {
	pending:= make ( chan bool , n)
	go func() {
		for {
			pending<- true
			rand:=yulib.RandInt(1,13)
			duration:=time.Duration(rand) * time.Second
			time.Sleep(duration)
		}		
	}()
	return pending	
}


