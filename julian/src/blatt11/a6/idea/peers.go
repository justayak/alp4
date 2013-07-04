package main
import (
	"net"
	"fmt"
	"time"
)
// ~~~~~~~~~~~~~~~~~~~~~
func main(){
	fmt.Println("select your port:")
	var port string
	
	p1 := NewPeer("1555")
	p2 := NewPeer("1556")
	
	
	
	
	fmt.Scanf("%s",&port)
	fmt.Println(p1.port + " | " + p2.port)
	
	
	p1.talkTo(p2)
	
	
	
	time.Sleep(time.Minute)
	
}
// ~~~~~~~~~~~~~~~~~~~~~
type Peer struct {
	connection net.Listener
	port string
	ip string
}
func NewPeer(port string ) *Peer {
	ln, _ := net.Listen("tcp", ":"+port)
	go func() {
		stop := false
		for !stop {
			_, err:= ln.Accept()
			if err != nil {
				stop = true 
			}
			fmt.Println("accepted")
		}
	}()
	return &Peer{ln,port,"localhost"}
}
func (p *Peer) talkTo(other *Peer) {
	conn,_:= net.Dial("tcp", other.port)
	conn.Write([]byte("Hallo"))
}