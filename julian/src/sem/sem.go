package sem
type Semaphore struct{
	c chan bool
}

func New (n uint) *Semaphore{
	x:=new (Semaphore)
	x.c = make(chan bool, n)
	for i:= uint(0);i<n;i++ {x.c<- true}
	return x
}

func (x *Semaphore) P() {<-x.c}
func (x *Semaphore) V() {x.c<-true}