package deepcolor

type QueueChannel struct {
	Chan  chan interface{}
	input chan int
	size  int
	cache *Queue
}

func NewQueueChannel(size int) *QueueChannel {
	qc := &QueueChannel{
		Chan:  make(chan interface{}, size),
		input: make(chan int, 1),
		size:  size,
		cache: NewQueue(),
	}
	go qc.magic()
	return qc
}

func (q *QueueChannel) Size() int {
	return q.size
}

func (q *QueueChannel) Pop() {

}

func (q *QueueChannel) Push(elem interface{}) {
	q.cache.Push(elem)
	select {
	case q.input <- 1:
	default:
	}
}

func (q *QueueChannel) magic() {
	for q.Chan != nil {
		if q.cache.Count() > 0 {
			q.Chan <- q.cache.Pop()
		} else {
			<-q.input
		}

	}
}
