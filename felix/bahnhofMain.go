package main 

import ("bahnhof"
		"fmt"
)

func main() {
	var anzahlGleise uint = 2
	b:=bahnhof.New(anzahlGleise)
	ch:=make(chan bool)
	ch1:=make(chan bool)
	ch2:=make(chan bool)
	go b.Einfahren(1,ch)
	go b.Einfahren(2,ch1)
	go b.Einfahren(3,ch2)
	<-ch
	<-ch1
	<-ch2
	if b.Leer(){
		fmt.Println("Fertig")
	}
	
}

