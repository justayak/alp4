package queue

type Queue struct {
	queue [] interface{}
}

func New() Queue{
	queue:=make([]interface{}, 0)
	
	result := Queue{
		queue : queue,
	}
	
	return result
}

func (q *Queue) Enq(i interface{}){
	q.queue=append(q.queue,i)
}

func (q *Queue) Deq() interface{}{
	result := q.queue[0]
	q.queue= (q.queue)[1:]
	return result
}

func (q *Queue) Size() int{
	return len(q.queue)
}

func (q *Queue) Peek() interface{}{
	return q.queue[0]
}

func (q *Queue) IsEmpty() bool{
	return len(q.queue) == 0
}