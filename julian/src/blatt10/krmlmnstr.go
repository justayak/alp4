package main
import (
	"fmt"
	"time"
	"yulib"
	."strconv"
)

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// M A I N
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func main(){
	fmt.Println("init")
	keksdose=make( chan bool, KEKSDOSE_MAX)
	
	temp:= make ( chan *Monster, 99)
	
	temp<-NewMonster("Julian")
	temp<-NewMonster("Alexander")
	temp<-NewMonster("Felix")
	temp<-NewMonster("Timo")
	temp<-NewMonster("Memel")
	temp<-NewMonster("Hans")
	temp<-NewMonster("Franz")
	temp<-NewMonster("Fritz")
	temp<-NewMonster("Sid") 
	start(temp)
}

func start (list chan *Monster){
	for {
		monster :=<- list
		go monster.simulate()
	}
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// D E F
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

const (
	KEKSDOSE_MAX = 20
	OFEN_MAX = 10
	MAGEN_MAX = 5
)

var keksdose chan bool 

func PrintKeksdose() {
	n:=len(keksdose)
	result:="("
	for i:=0;i<KEKSDOSE_MAX;i++{
		if i<n{
			result+="[o]"
		}else{
			result+="[ ]"
		}	
	}
	fmt.Println(result)
}

//~~~ Monster  ~~~

type Monster struct {
	name string 
}

// Erzeuge ein Monster-Objekt
func NewMonster(name string ) *Monster {
	result:= new (Monster)
	result.name = name
	return result
}

func (m *Monster) backe() {
	fmt.Println("Monster <" + m.name + "> : starte backen..")
	// Erzeuge zufällig viele Kekse
	kekseAnzahl:=yulib.RandInt(3,OFEN_MAX)
	time.Sleep(time.Second/2) // backen..
	fmt.Println("Monster <" + m.name + "> : backen beendet [" + Itoa(kekseAnzahl) + "]")
	for i:=0;i<kekseAnzahl;i++{
		timeout:=time.After(5*time.Second)
		select {
			case keksdose <- true: // Füge Kekse in die Keksdose ein
				PrintKeksdose() 
			case <-timeout:
				fmt.Println("Monster <" + m.name + "> : wartet zu lange, die Kekse sind kalt...")
				return
		}
		
		
	}	
}

func (m *Monster)iss() {
	n:=yulib.RandInt(1,MAGEN_MAX) 
	fmt.Println("Monster <" + m.name + "> : moechte " + Itoa(n) + " Kekse essen")
	for i:=0;i<n;i++{
		timeout:=time.After(5*time.Second)
		select {
			case <-keksdose:
				time.Sleep(time.Second/4) // iss!
				PrintKeksdose()
			case <-timeout:
				fmt.Println("Monster <" + m.name + "> wartet zu lange.. Bricht ab..") 
				return
		}
		 
	}
	fmt.Println("Monster <" + m.name + "> ist satt") 
}

func (m *Monster) simulate() {
	for {
		s:=yulib.RandInt(0,2)
		switch s{
			case 0:
				m.backe()
			case 1:
				m.iss()
		}
		time.Sleep(time.Second*5)
	}
}
