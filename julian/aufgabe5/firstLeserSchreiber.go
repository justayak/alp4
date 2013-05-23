package main


import ("fmt")
import "sync"

var (nR, nW uint; m sync.Mutex; c *sync.Cond = sync.NewCond (&m))



func ReaderIn(){
	m.Lock()
	for nW > 0{
		c.Wait()	
	}
	
	nR++
	c.Signal()
	m.Unlock()
}

func ReaderOut(){
	m.Lock()
	nR--
	c.Signal()
	m.Unlock()
}

func WriterIn(){
	m.Lock()
	for nR > 0 || nW > 0 {
		c.Wait()	
	}	
	nW = 1
	m.Unlock()
}

func WriterOut(){
	m.Lock()
	nW = 0
	c.Signal()
	m.Unlock()
}

func main () {



   fmt.Println("Start")
}