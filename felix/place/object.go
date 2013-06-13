package place

import ("fmt"
)

type Object struct{
	size uint
}

func New(n uint) *Object{
	if n==1 || n==2|| n==4{
		o:=new(Object)
		o.size = n
		return o
	}else{
		panic(fmt.Sprintf("Größe des Onjektes falsch",n))
		return nil
	}
	return nil
}