package main 
import ("fmt"; ."queue"; "yulib"; "sync"; "strings";)

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
	Sender string 
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
	Current chan Package
	hasElement bool
}

func (k *Kanal) listen(doneA chan bool, doneB chan bool){
	for {
		current := <-k.Current
		fmt.Print(current.Sender)
		for _,element := range current.Elements{
			fmt.Print(":")
			fmt.Print(element.Value)
		}
		fmt.Println(":")
	}
	
	<-doneA
	<-doneB
	close(Condition)
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
var cond sync.Cond

func (p *Position) isEmpty() bool {
	x := p.Pending.IsEmpty()
	return x
}

var Condition chan bool = make(chan bool, 1)

// versende die Pending-Objekte
func (p *Position) send(k *Kanal, done chan bool, flag string ){	
	for {
		if p.Pending.IsEmpty() {
			done <- true
			break			
		}
		// build Package:
		
		paket := Package{}		
		sum:= 0
		s := make( []Object,0 )		
		for {
			if p.Pending.IsEmpty() {
				break			
			}
			sum += p.Pending.Peek().(Object).Value					
			if sum > 4 { break }
			s = append(s, p.Pending.Deq().(Object))
		}
		paket.Elements = s
		paket.Sender = flag
		// Paket ist fertig, versuche, zu versenden
				
				
		fmt.Print("A.K:")
		fmt.Print(k.A.IsPriority)
		fmt.Print("  B.K:")
		fmt.Println(k.B.IsPriority)
		
		
		transmit := func (o *Position){					
			k.Current <- paket	
			// cond.Signal()
			Condition <- true			
			// o.IsPriority = true
			// p.IsPriority = false			
			if strings.Contains(flag, "A"){
				k.A.IsPriority = false
				k.B.IsPriority = true
			}else{
				k.B.IsPriority = false
				k.A.IsPriority = true
			}
		}
		
		lock.Lock()			
		other := k.getOther(p)			
		if other.isEmpty() {
			transmit(other)
			lock.Unlock()
		}else{
			if p.IsPriority {
				transmit(other)
				lock.Unlock()
			}else{
				lock.Unlock()
				// cond.Wait()
				<-Condition
				lock.Lock()
				transmit(other)
				lock.Unlock()
			}			
		}
		
	}
	<-done
	close(done)
	
}

func main () {
	doneA := make( chan bool )
	doneB := make( chan bool )
	kanalChan := make (chan Package,1)
	
	A := generatePosition()	
	A.IsPriority = true
	B := generatePosition()
	B.IsPriority = false
	
	kanal := Kanal {
		A : A,
		B : B,
		hasElement : false,
		Current : kanalChan,
	}
	
	go A.send(&kanal, doneA, "A")
	go B.send(&kanal, doneB, "B")
	go kanal.listen(doneA,doneB)
	
	<-doneA
	<-doneB
}