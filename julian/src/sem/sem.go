package sem 
type empty struct {}
type Semaphore chan empty

var sem = make(Semaphore, 1)

// acquire n resources
func (s Semaphore) P(n int) {
    e := empty{}
    for i := 0; i < n; i++ {
        s <- e
    }
}

// release n resources
func (s Semaphore) V(n int) { 
    for i := 0; i < n; i++ {
        <-s
    }
}