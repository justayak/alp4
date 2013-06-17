package main
import (
	"fmt"
	//"sem"
	"queue"
	"os/exec"
	"time"
	"os"
	//"strconv"
	"sync"
)

var n int = 15
var count int = 0
var queueIn queue.Queue
var queueOut queue.Queue
var aufzugUnten bool = true
var elevator sync.Mutex 

var sekunde = 1000000000

func render(rendering chan string ){	
	for {
		s := <- rendering
		cmd:=exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()		
		fmt.Println(s)		
	}		
}

func drawScene(carsIn int , carsOut int , carsParked int ,n int ) string {
	result := ""
	for i:=0;i<carsIn;i++ {
		result += "UU>"
	}	
	if aufzugUnten {
		result += " [ v ] "
	}else{
		result += " [ ^ ] "
	}	
	for i:=0;i<carsOut;i++ {
		result += "<U"
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
	
}

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
	
	<-done
	time.Sleep(time.Duration(sekunde))
}