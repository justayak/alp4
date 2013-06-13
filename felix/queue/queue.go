package queue

type Queue [] int

func NewQueue()Queue{
	queue:= make([]int,0)
	return queue
}

func Enqueue(q *Queue, i int){
	*q=append(*q,i)
}
func Dequeu(q *Queue){
	*q= (*q)[1:]
}

func Head(q Queue) int{
	return q[0]
}

func Empty (q Queue) bool{
	return len(q)==0
}
