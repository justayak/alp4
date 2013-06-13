package queue

type Queue struct{
	queue [] interface{}
}

func NewQueue() *Queue{
	q:=new(Queue)
	q.queue = make([]interface{},0)
	return q
}

func (q *Queue)Enqueue(i interface{}){
	q.queue=append(q.queue,i)
}
func (q *Queue)Dequeu(){
	q.queue = q.queue[1:]
}

func (q *Queue)Head() interface{}{
	return q.queue[0]
}

func  (q *Queue)Empty() bool{
	return len(q.queue)==0
}

func (q *Queue)Liste() []interface{}{
	return q.queue
}
