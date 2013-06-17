package main
import (
	"fmt"
	//"sem"
	"queue"
	//"os/exec"
	//"time"
	//"os"
)

var n int = 15
var count int = 0
var queueIn queue.Queue
var queueOut queue.Queue
var aufzugUnten bool = true

var sekunde = 1000000000

func render(done chan bool, rendering chan string ){
	for {
		s := <- rendering
		fmt.Println(s)
	}
	done <- true
}

func fill(rendering chan string ){
	for {
		rendering<-"lol"
	}
}

func main(){
	queueIn = queue.New()
	queueOut = queue.New()
	
	done := make(chan bool, 0)
	rendering := make(chan string, 0)
	
	//rendering <- "Hallo"
	go render(done,rendering)
	go fill(rendering)
	
	close(rendering)
	
	<-done
	
	//sema := sem.New(5)
	//sema.P()
	//fmt.Println("Start")
	//sema.V()
}