package place


type Object struct{
	size uint
}

func New(n uint) *Object{
	o:=new(Object)
	o.size = n
	return o
}

func (o *Object)GetSize() uint{
	return o.size
}