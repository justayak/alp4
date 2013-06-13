package moveObject

import ("place"
		"sync"
		"time"
)

type Mover struct{
	mu sync.Mutex
	favorit uint
}

func New() *Mover{
	m:=new(Mover)
	m.favorit = 0
	return m
}

func (m *Mover)makePackage(a *place.Place) []*place.Object{
	var sum uint = 0
	var list []*place.Object
	list=make([]*place.Object,0)
	if !a.Empty(){
	for sum+a.FirstElem().GetSize()<=4{
		sum +=a.FirstElem().GetSize()
		list=append(list,a.Remove())
		if a.Empty(){
			break;
		} 	
	}}
	return list
}

//Von a nach b senden
func (m *Mover)Send(a,b *place.Place){
	p:=m.makePackage(a)
	b.InsertPack(p)
	if m.favorit==0{
		m.favorit=1
	}else{
		m.favorit=0
	}
}

func (m *Mover)transmitHelper(a,b *place.Place){
	m.mu.Lock()
	m.Send(a,b)
	m.mu.Unlock()
}

//transmit von a nach b
func (m *Mover)Transmit(a,b *place.Place, ch chan bool){
	for a.GetNumber()!=m.favorit{
		time.Sleep(1)	
		if b.Empty(){
			break;
		}
	}
	m.transmitHelper(a,b)
	ch<-true
}