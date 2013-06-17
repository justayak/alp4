package main
import (
	"fmt"
	//"sem"
	"queue"
	"os/exec"
	"time"
	"os"
	"strconv"
)

var n int = 15
var count int = 0
var queueIn queue.Queue
var queueOut queue.Queue
var aufzugUnten bool = true

var sekunde = 1000000000

func render(rendering chan string ){
	
	for {
		cmd:=exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		s := <- rendering
		fmt.Println(s)
		
	}	
	
}

func fill(done chan bool, rendering chan string ){
	for i:=0;i<10;i++ {		
		result := "lol" + strconv.Itoa(i)
		rendering<- result
	}
	done <- true
}

func main(){
	queueIn = queue.New()
	queueOut = queue.New()
	
	done := make(chan bool, 0)
	rendering := make(chan string, 1)
	
	rendering <- "Hallo"
	go render(rendering)
	go fill(done,rendering)
	
	//close(rendering)
	
	<-done
	time.Sleep(time.Duration(sekunde))
	//close (done)
	
	// for i:=0; i < 10000;i++{
		// cmd:=exec.Command("clear")
		// cmd.Stdout = os.Stdout
		// cmd.Run()
		// fmt.Println("aa", i)
		// time.Sleep(time.Duration(sekunde))
	// }
}