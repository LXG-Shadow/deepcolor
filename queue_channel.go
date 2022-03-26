package deepcolor

type QueueChannel struct {
	Chan  chan interface{}
	size  int
	cache *Queue
}

func NewQueueChannel(size int) *QueueChannel {
	qc := &QueueChannel{
		Chan:  make(chan interface{}, size),
		size:  size,
		cache: NewQueue(),
	}
	go qc.magic()
	return qc
}

func (q *QueueChannel) Size() int {
	return q.size
}

func (q *QueueChannel) Push(elem interface{}) {
	q.cache.Push(elem)
}

func (q *QueueChannel) magic() {
	for q.Chan != nil {
		if q.cache.Count() > 0 {
			select {
			case q.Chan <- q.cache.Front():
				q.cache.Pop()
			}
		}

	}
}
