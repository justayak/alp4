// ps: der src von Hrn Maurer ist ein scheiss
// kann man seine Variablen BITTE mal vernuenftig 
// benennen? OMG arghhh
package buffer

type Channel struct { 
	In, Out chan byte 
}

func NewChannel (n uint) *Channel {
	result:=new( Channel )
	result.In, result.Out = make(chan byte), make(chan byte)
	go func(){
		buffer:=make([]byte, n)
		var count, in, out	uint
		for {
			if count == 0 {
				buffer[in] = <-result.In
				in = (in + 1)% n
				count = 1
			}else if count == n{
				result.Out <- buffer[out]
				out = (out + 1)%n
				count = n - 1				
			}else{
				select {
					case buffer[in] = <-result.In:
						in = (in + 1) % n
						count++
					case result.Out<-buffer[out]:
						out = (out + 1)%n
						count--
				}
			}	
		}
	}()
	return result
}

func (x *Channel) Insert(b byte) { x.In <- b }
func (x *Channel) Get() byte { return <-x.Out }
