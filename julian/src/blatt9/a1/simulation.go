package main
import (
	"fmt"
	."elevator"
	//"math/rand"
	"os/exec"
	"time"
	"os"
	)
	
func main(){
	
	rendering := make(chan string,1)
	done := make(chan bool)
	
	deck := NewParkingDeck(10)
	
	go Rein(deck, rendering,done)
	go Raus(deck, rendering,done)
	go render(rendering)
	go fill(deck, rendering, done)
	<-done
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~
// M E T H O D S
// ~~~~~~~~~~~~~~~~~~~~~~~~~

func Rein(d *ParkingDeck, rendering chan string,done chan bool){	
	for {
		// fill(d, rendering,done)
		d.TryEnter()	
		// fill(d, rendering,done)		
		time.Sleep(time.Duration(time.Second))
	}
}

func Raus(d *ParkingDeck, rendering chan string,done chan bool){
	for {		
		// fill(d, rendering,done)
		time.Sleep(time.Second * 4)
		d.TryLeave()		
		// fill(d, rendering,done)
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~
// R E N D E R I N G
// ~~~~~~~~~~~~~~~~~~~~~~~~~

func fill(deck *ParkingDeck, rendering chan string, done chan bool ){	
	for i:=0; i<50;i++ {
		<-deck.RenderNotifier		
		e:=deck.Elevator()
		aufzugUnten := e.IsFull
		rendering <- drawScene(len(e.In), len(e.Out), deck.ParkedCars(), deck.N(),aufzugUnten, e.State)
		// time.Sleep(time.Duration(time.Second/5))
	}
}

// rendert die simulation auf die Console (unter Linux wird die console 
// jedes mal ge"clear"t - funzt leider nicht unter Win)
func render(rendering chan string ){	
	for {
		s := <- rendering
		cmd:=exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()		
		fmt.Println(s)		
	}
}

// Zeichnet jeweils die Scene
func drawScene(carsIn,carsOut,carsParked,n int, aufzugUnten bool, state int) string {
	result := ""
	for i:=0;i<carsIn;i++ {
		result += " U>"
	}	
	// if aufzugUnten {
		// result += " [___] "
	// }else{
		// result += " [---] "
	// }	
	switch state {
		case In:
			result += " [U>_] "
		case Out:
			result += " [_<U] "
		case None:
			if aufzugUnten {
			result += " [___] "
			}else{
				result += " [---] "
			}	
	}
	for i:=0;i<carsOut;i++ {
		result += "<U "
	}	
	result += ":::::"	
	for i:=0;i<carsParked;i++ {
		result += "[<U]"
	}	
	for i:=0;i<(n-carsParked);i++ {
		result += "[_]"
	}
	return result
}