package main 
import ("math/rand"; "time"; "fmt")
//import ("fmt"; "time";)


var ( Counter int; done chan bool;)

var sleepTime = 99999

func v() {
	time.Sleep (time.Duration(rand.Int63n (1e5)))
}
 
func inc (n *int) { 
	Accu:= *n // "LDA n" 
	v()
	Accu++ // "INA" 
	v() 
	*n = Accu // "STA n"
	fmt.Println(*n) 
	v() 
}

func count (p int) { 
	const N = 1 
	for n:= 0; n < N; n++ { 
		inc (&Counter) 
	} 
	done <- true 
}
func main () {
	Counter = 0 
	done = make (chan bool) 
	go count (0); 
	go count (1) 
	go count (2); 
	go count (3) 
	go count (4); 
	go count (5) 
	go count (6); 
	go count (7) 
	go count (8); 
	go count (9) 
	<-done; 
	<-done 
	<-done; 
	<-done 
	<-done; 
	<-done 
	<-done; 
	<-done 
	<-done; 
	<-done 
	fmt.Printf ("Counter = %d\n", Counter) 
}