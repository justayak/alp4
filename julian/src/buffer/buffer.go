// ps: der src von Hrn Maurer ist ein scheiss
// kann man seine Variablen BITTE mal vernuenftig 
// benennen? OMG arghhh
package buffer

type Channel struct { 
	In, Out chan string 
	asyncListenerStarted bool
	n uint
	async [] chan string 
	
	asyncChannels chan chan string
}

func NewChannel (n uint) *Channel {
	result:=new( Channel )
	result.n = n
	result.asyncListenerStarted = false
	result.asyncChannels = make(chan chan string, n)
	result.In, result.Out = make(chan string), make(chan string)
	go func(){
		buffer:=make([]string, n)
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

func (x *Channel) Insert(b string) { x.In <- b }
func (x *Channel) Get() string { return <-x.Out }

// Die frage ist irgendwie komisch: ist's so gemeint?
func (x *Channel) GetAsync2() <-chan string {
	result:=make(chan string, x.n)
	x.asyncChannels<-result
	
	if !x.asyncListenerStarted {
		x.asyncListenerStarted = true
		go func() {
			for {
				value:=<-x.Out
				go func() {
					current:=<-x.asyncChannels
					current<-value
					x.asyncChannels<-current
				}()
			}
		}()
	}
	
	return result
}

func (x *Channel) GetAsync() <-chan string {
	result:=make(chan string, x.n)
	x.async = append(x.async, result)
	if !x.asyncListenerStarted {
		x.asyncListenerStarted = true
		go func() {
			for {
				value:=<-x.Out
				for i:=0; i<len(x.async);i++{					
					go func() {
						x.async[i]<-value
					}()
				}
			}
		}()
	}
	return result
}