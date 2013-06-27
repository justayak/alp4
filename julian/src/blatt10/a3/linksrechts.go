package main
import (
	"fmt"
	"time"
	"os"
	"os/exec"
)

func main() {
	leftPending := createPending(N)
	rightPending := createPending(N)
	crit := NewCriticalRoad()
	render := crit.simulate(leftPending,rightPending, N)
	go draw(render)
	duration:=time.Duration(10) * time.Second
	time.Sleep(duration)
}

// rendert den krams blabla
func draw(render <-chan string ){
	for {
		o := <- render
		cmd:=exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()	
		fmt.Println(o)
		time.Sleep(time.Second)
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

// Stellt den kritischen Abschnitt dar
type CriticalRoad struct {
	PreferLeft bool 
	HasValue bool
	HasExited bool
	Direction int	
}

// Erzeugt einen neuen kritischen Abschnitt
func NewCriticalRoad() *CriticalRoad {
	result:= new( CriticalRoad )
	result.PreferLeft = true 
	result.HasValue = false
	result.HasExited = false
	return result
}

func (road *CriticalRoad)simulate(left,right <-chan bool, n int )<- chan string {
	result:= make( chan string )
	changed:= true 
	go func() {
		for {
			changed = false
			l:=""
			m:="[__]"
			r:=""
			if road.HasExited {
				road.HasExited = false
				changed = true		
			}else if road.HasValue {
				road.HasValue = false 
				road.HasExited = true
				road.PreferLeft = !road.PreferLeft
				if road.Direction == Right {
					m = "[_>]"
				} else {
					m = "[<_]"
				}	
				changed = true				
			} else {
				if road.PreferLeft {
					if len( left ) > 0 {
						<-left
						road.HasValue = true 
						m = "[>_]"
						road.Direction = Right
						changed = true
					} else if len ( right ) > 0 {
						<-right
						road.HasValue = true 
						m = "[_<]"
						road.Direction = Left
						changed = true
					}
				}else {
					if len( right ) > 0 {
						<-right
						road.HasValue = true 
						m = "[_<]"
						road.Direction = Left
						changed = true
					} else if len ( left ) > 0 {
						<-left
						road.HasValue = true 
						m = "[>_]"
						road.Direction = Right		
						changed = true						
					}					
				}
			}
			if changed  {
				for i:=n;i>0;i--{
					if i < len(left) {
						l+="[>]"
					}else {
						l+="[ ]"
					}					
				}
				
				for i:=0;i<n;i++{
					if i < len(right) {
						r+="[<]"
					}else {
						r+="[ ]"
					}					
				}
				
				result <- l + "___" + m + "___" + r
			}
			
			changed = false 
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
		}		
	}()
	return pending	
}


