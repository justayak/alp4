package main
import (
	"buffer"
	"fmt"
	"time"
)

func main(){
	
	buffer := buffer.NewChannel(10)
	
	buffer.Insert("Hallo")
	buffer.Insert("Welt")
	
	x:=buffer.GetAsync2()
	y:=buffer.GetAsync2()
	
	go testAsync(x, "x")
	go testAsync(y, "y")
	
	
	duration:=time.Duration(10) * time.Second
	time.Sleep(duration)
	fmt.Println("ende")
}

func testAsync(msg <-chan string, z string){
	for {
		fmt.Println(<-msg + " | " + z)
	}
}