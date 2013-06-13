package main

import (
	."sem"
	"fmt"
	"time"
	)
	
// =======================
// M A I N
// =======================

func main(){
	
	bahnhof := New(5)
	
	go bahnhof.TrainEnters()
	go bahnhof.TrainEnters()
	go bahnhof.TrainEnters()
	go bahnhof.TrainEnters()
	go bahnhof.TrainEnters()
	go bahnhof.TrainEnters()	
	go bahnhof.TrainEnters()
	
	go bahnhof.TrainLeaves()
	go bahnhof.TrainLeaves()
	go bahnhof.TrainLeaves()
	
	time.Sleep(2000) // damit auch angezeigt wird..
	fmt.Println("Ende")
}

// =======================
// I M P L
// =======================

type TrainStation struct {
	Tracks Semaphore
}

func New(tracks uint) *TrainStation{
	t:= new(TrainStation)
	t.Tracks = make(Semaphore, tracks)
	return t
}

func (t *TrainStation)TrainEnters(){	
	t.Tracks.P(1)		
	fmt.Println("Train enters..")
}

func (t *TrainStation)TrainLeaves(){	
	t.Tracks.V(1)
	fmt.Println("Train leaves..")
}