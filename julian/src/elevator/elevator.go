package elevator
import (
	"time"
)

// ~~~~~~~~~~~~~~~~~~~~~~~
// E L E V A T O R
// ~~~~~~~~~~~~~~~~~~~~~~~~

type Elevator struct{
	In chan bool  // Autos, die rein wollen
	Out chan bool // Autos, die raus wollen
	parkingDeck *ParkingDeck
	IsDown bool 		
}

func newE(n int) *Elevator {
	result:= new( Elevator )
	result.In = make(chan bool,n) // damit wir das auch schoen anzeigen koennen..
	result.Out = make(chan bool,n)
	result.IsDown = true
	return result
}

func (e *Elevator) simulate() {
	go func() {
		for {	
			if (e.parkingDeck.IsFull()){
				// warte, bis einer raus will
				<-e.Out
			} else {
				// parkdeck ist immer noch nicht voll,
				// da wir NUR hier in "p.cars" schreiben
				select {
					case <- e.In:						
						e.parkingDeck.cars <- true
					case <-e.Out:
						// einer will raus 
				}
			}
			time.Sleep(time.Second)
		}
	}()
}


// ~~~~~~~~~~~~~~~~~~~~~~~
// P A R K I N G D E C K
// ~~~~~~~~~~~~~~~~~~~~~~~~

type ParkingDeck struct {
	cars chan bool 
	n int
	elevator *Elevator
}

func (e *ParkingDeck) N() int {
	return e.n
}

func NewParkingDeck(n int) *ParkingDeck {
	result:=new(ParkingDeck)
	result.cars = make(chan bool, n)
	result.n = n	
	e:=newE(n)
	e.parkingDeck = result
	result.elevator = e	
	e.simulate()
	return result
}

func (e *ParkingDeck) Elevator() *Elevator {
	return e.elevator
}

func (e *ParkingDeck) TryLeave(){
	if len(e.cars) > 0 {
		e.elevator.Out <- <- e.cars
	}
}

func (e *ParkingDeck) TryEnter(){
	e.elevator.In <- true
}

func (e *ParkingDeck) ParkedCars() int {
	return len(e.cars)
}

func (p *ParkingDeck) IsFull() bool {
	return p.ParkedCars() >= p.n
}

