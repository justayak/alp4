package sem

type Channel struct{
Botschaft string
	s *Semaphore
	e *Semaphore
	mutex *Semaphore
}

func Init() *Channel{
	x:=new(Channel)
	x.s = New(0)
	x.e = New(0)
	x.mutex = New(1)
	return x
}

func (x *Channel) send(c string){ 
	x.mutex.P()
	x.Botschaft=c
	x.s.V()
	x.e.P()
}

func (x *Channel) recv(c string){
	x.s.P()
	c = x.Botschaft
	x.e.V()
	x.mutex.V()
}
