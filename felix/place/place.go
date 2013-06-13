package place

import ("container/list"
)

type Place struct{
	queue *list.List
	nummer uint
}

func (p *Place)GetNummer() uint{
	return p.nummer
}

func (p *Place)Einfügen(objList object[]){
	for i:=0;i<len(objList);i++{
		p.queue.PushBack(objList[i]
	}
}

func (p *Place)Senden(){
	sum := 0
	for sum+p.queue.Front(), <= 4{
		
}
