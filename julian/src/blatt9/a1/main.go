package main
import (
	"fmt"
	//"sem"
	//"queue"
	"os/exec"
	"time"
	"os"
	//"strconv"
)

var n int = 15
var count int = 0
//var queueIn queue.Queue
//var queueOut queue.Queue
var aufzugUnten bool = true

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
		result += "<UU"
	}
	
	result += "__"
	
	for i:=0;i<carsParked;i++ {
		result += "[<UU]"
	}
	
	for i:=0;i<(n-carsParked);i++ {
		result += "[__]"
	}
	return result
}

func fill(done chan bool, rendering chan string ){
	for i:=0;i<10;i++ {		
		result := drawScene(i, 2, 4, 20)
		rendering<- result
	}
	done <- true
}

func main(){
	//queueIn = queue.New()
	//queueOut = queue.New()
	
	done := make(chan bool, 0)
	rendering := make(chan string, 1)
	
	go render(rendering)
	go fill(done,rendering)
	
	<-done
	time.Sleep(time.Duration(sekunde))
}