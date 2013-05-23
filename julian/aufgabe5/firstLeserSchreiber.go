package main


import ("fmt")
import "sync"



var busy bool
var waiting int
var mu sync.Mutex
var OKtoread = sync.NewCond(&mu)
var OKtowrite = sync.NewCond(&mu)
var readercount int

func ReaderStart(){
	mu.Lock()
	if busy {
		OKtoread.Wait()
	}
	readercount++
	OKtoread.Signal()
	
	mu.Unlock()
}

func ReaderEnd(){
	mu.Lock()
	readercount--
	if readercount == 0 {
		OKtowrite.Signal()
	}
	mu.Unlock()
}

func WriterStart(){
	mu.Lock()
	waiting++
	if busy || readercount != 0{
		OKtowrite.Wait()
	}
	busy = true	
	waiting--
	mu.Unlock()
}

func WriterEnd(){
	mu.Lock()
	busy = false
	if waiting > 0 {
		OKtoread.Signal()
	}else{
		OKtowrite.Signal()
	}
	mu.Unlock()
}

func main () {



   fmt.Println("Start")
}