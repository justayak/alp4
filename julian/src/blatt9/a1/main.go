package main
import (
	"fmt"
	//"sem"
	"queue"
	"os/exec"
	"time"
	"os"
	//"strconv"
	"yulib"
	"sync"
)

var n int = 15
var count int = 0
var queueIn queue.Queue
var queueOut queue.Queue
var aufzugUnten bool = true
var elevator sync.Mutex 
var in sync.Mutex
var out sync.Mutex

var sekunde = 1000000000

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
func drawScene(carsIn int , carsOut int , carsParked int ,n int ) string {
	result := ""
	for i:=0;i<carsIn;i++ {
		result += " U>"
	}	
	if aufzugUnten {
		result += " [ v ] "
	}else{
		result += " [ ^ ] "
	}	
	for i:=0;i<carsOut;i++ {
		result += "<U "
	}	
	result += "____"	
	for i:=0;i<carsParked;i++ {
		result += "[<U]"
	}	
	for i:=0;i<(n-carsParked);i++ {
		result += "[_]"
	}
	return result
}

func fill(done chan bool, rendering chan string ){
	for i:=0;i<10;i++ {		
		result := drawScene(queueIn.Size(), queueOut.Size(), count, n)		
		rendering<- result
		time.Sleep(time.Duration(sekunde))
	}
	done <- true
}

func simulate(){
	for {
		time.Sleep(time.Duration(sekunde/2))
		selector := yulib.RandInt(0,10)
		switch selector { 
			case 0, 1,2,3,4,5:
				autoIn()
			case 6,7,8:
				autoOut()				
			// 9,10 sind "leer"
		}
	}
}

// Ein Auto moechte auf das Parkdeck fahren
func autoIn(){
	in.Lock()
	queueIn.Enq("car")
	in.Unlock()
}

// Ein Auto (falls eins vorhanden ist) moechte das Parkdeck verlassen
func autoOut(){
	if count > 0 {
		out.Lock()
		queueOut.Enq("car")
		count--
		out.Unlock()
	}
}

func queueInSimulate(){
	for {
		time.Sleep(time.Duration(sekunde))		
		for {			
			if queueIn.Size() == 0 {
				break
			}else{		
				in.Lock()			
				elem := queueIn.Deq()
				fmt.Println(elem)
				in.Unlock()
			}
			
		}
		
	}
}

func queueOutSimulate(){
	for {
		time.Sleep(time.Duration(sekunde))	
		for {
						
		}
	}
}

// sorgt dafuer, dass der Fahrstuhl nicht verklemmt
// blabla nicht die beste Loesung blabla
func simulateElevator(){
	time.Sleep(time.Duration(sekunde))
	elevator.Lock()	
	if aufzugUnten {
		if queueOut.Size() > 0 && queueIn.Size() == 0 {
			aufzugUnten = false
		}
	}else{
		if queueIn.Size() > 0 && queueOut.Size() == 0 {
			aufzugUnten = true
		}
	}	
	elevator.Unlock()
}


// ==============================
// M A I N
// ==============================
func main(){
	queueIn = queue.New()
	queueOut = queue.New()
	
	//in = sem.New(0)
	//out = sem.New(0)
	//elevator = make( sem.Semaphore, 0)
	
	done := make(chan bool, 0)
	rendering := make(chan string, 1)
	
	go render(rendering)
	go fill(done,rendering)
	go simulateElevator()
	go simulate()
	go queueInSimulate()
	//go queueOutSimulate()
	
	<-done
	time.Sleep(time.Duration(sekunde))
}