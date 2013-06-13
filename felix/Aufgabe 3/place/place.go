package place

import ("queue"
		"fmt"
)

type Place struct{
	objectQueue *queue.Queue
	nummer uint
}

func NewPlace(n uint) *Place{
	p:=new(Place)
	p.objectQueue = queue.NewQueue()
	p.nummer=n
	return p
}

func (p *Place)Insert(o *Object){
	p.objectQueue.Enqueue(o)	
}

func (p *Place)Remove() *Object{
	i:=p.objectQueue.Head()
	p.objectQueue.Dequeu()
	return i.(*Object)
}

func (p *Place)InsertPack(o []*Object){
	for i:=0;i<len(o);i++{
		p.objectQueue.Enqueue(o[i])
	}	
}

func (p *Place)FirstElem() *Object{
	return p.objectQueue.Head().(*Object)
}

func (p *Place)GetNumber() uint{
	return p.nummer
}

func (p *Place)Empty() bool{
	return p.objectQueue.Empty()
}

func (p *Place)Print(){
	l:=p.objectQueue.Liste()
	for i:=0;i<len(l);i++{
		fmt.Print(l[i].(*Object).GetSize(),", ")
	}
	fmt.Println()
}