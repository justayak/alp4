package main
import (
	"fmt"
	"yulib"
	"time"
)

func main() {
	fmt.Println("hallo")
	leftPending = createPending(N)
	rightPending = createPending(N)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// I M P L E M E N T A T I O N
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~

const (
	N = 5
)

// Warteschlangen für die jeweiligen Seiten
var leftPending, rightPending <-chan bool

// Erzeugt die Strasse VOR dem kritischen
// Abschnitt... blabla
func createPending(n uint) <- chan bool {
	pending:= make ( chan bool , n)
	go func() {
		for {
			pending<- true
			rand:=yulib.RandInt(1,13)
			duration:=time.Duration(rand) * time.Second
			time.Sleep(duration)
		}		
	}()
	return pending	
}