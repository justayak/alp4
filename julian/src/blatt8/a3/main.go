package main 
import ("fmt"; ."queue"; "yulib"; "sync";)

const (
	BIG = 4
	MIDDLE = 2
	SMALL = 1
)

type Object struct {
	Value int
}

// Paket zum versenden der Elemente
type Package struct {
	Elements [] Object
}

type Position struct {
	id int
	Pending Queue
	IsPriority bool 
}

var currentPositionID = 0

// Erzeugt Position
func generatePosition() Position{
	queue := Queue{}	
	for i:= 0; i < 30; i++ {
		queue.Enq(generateObject())
	}
	result := Position{}	
	result.Pending = queue
	result.id = currentPositionID
	currentPositionID += 1
	return result
}


func (p *Position) hasPending() bool {
	return !p.Pending.IsEmpty()
}

type Kanal struct{
	A Position
	B Position
	Current Package
	hasElement bool
}

// Liefert das Andere Zielobjekt (umständliche Methode :( )
func (k *Kanal) getOther(p *Position) *Position{
	if p.id == k.A.id {
		return &k.B
	}
	return &k.A
}

// Erzeugt zufälliges Objekt
func generateObject() Object{
	s := yulib.RandInt(0, 10)
	v := 0
	switch {
		case s < 5:  v = SMALL
		case s >= 5 && s < 8: v = MIDDLE
		case s >= 8: v = BIG
	}
	result := Object{
		Value : v,
	}
	return result
}

var lock sync.Mutex

// versende die Pending-Objekte
func (p *Position) send(k *Kanal, done chan bool ){
	for {
		if p.Pending.IsEmpty() {
			done <- true
			break			
		}
		//other := k.getOther(p)		
		// build Package:
		paket := Package{}		
		sum:= 0
		s := make( []Object,0 )
		for {
			sum += p.Pending.Peek().(Object).Value		
			if sum > 4 { break }
			t := make ([]Object, len(s), (cap(s)+1))
			copy(t,s)
			s = t
			s[len(s)] = p.Pending.Deq().(Object)
			
		}
		paket.Elements = s
		// Paket ist fertig, versuche, zu versenden
		
		lock.Lock()
		// Kanal belegt?
		isBelegt := k.hasElement		
		lock.Unlock()
		
		
	}
	
}

func main () {
	//done := make( chan bool )
	
	
	A := generatePosition()	
	B := generatePosition()
	
	kanal := Kanal {
		A : A,
		B : B,
		hasElement : false,
	}
	
	fmt.Printf ("size A %d\n", A.Pending.Size()) 
	fmt.Printf ("size B %d\n", B.Pending.Size()) 
	fmt.Printf ("size B %d\n", kanal.hasElement) 
	
	elem := A.Pending.Deq().(Object)
	
	fmt.Printf("v %d\n", elem.Value)
}