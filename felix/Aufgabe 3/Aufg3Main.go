package main

import ("place"
		"moveObject"
)

func main(){
	A:=place.NewPlace(0)
	B:=place.NewPlace(1)
	ch,ch2:=make(chan bool),make(chan bool)
	M:=moveObject.New()
	o1:=place.New(1)
	o2:=place.New(2)
	o3:=place.New(4)
	A.Insert(o1)
	A.Insert(o2)
	B.Insert(o3)
	A.Print()
	B.Print()
	
	
		
	go M.Transmit(A,B,ch)
	go M.Transmit(B,A,ch2)
	
	<-ch;<-ch2
	
	A.Print()
	B.Print()
}
