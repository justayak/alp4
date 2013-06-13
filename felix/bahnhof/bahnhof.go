package bahnhof

import ("semaphore"
		"fmt"
)


type Bahnhof struct{
	gleise uint
	sem *semaphore.Imp
}

func New(n uint) *Bahnhof{
	m:=new (Bahnhof)
	m.gleise = 0
	m.sem = semaphore.New(n)
	return m
}

func (b *Bahnhof)Einfahren(n uint,ch chan bool){
	fmt.Println("Zug ",n," will einfahrt!")
	b.sem.P()
	b.gleise ++
	fmt.Println("Zug ",n," fährt ein!")
	fmt.Println("Zug ",n," will abfahrt!")
	b.gleise --
	fmt.Println("Zug ",n," fährt ab!")
	b.sem.V()	
	ch<-true
}

func (b *Bahnhof)Leer() bool{
	return b.gleise==0
}
