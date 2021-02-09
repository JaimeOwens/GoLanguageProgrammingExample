package chapter9

type request struct {
	key      string
	response chan<- result
}

type Memo4 struct {
	requests chan request
}

func New3(f Func) *Memo4 {
	memo := &Memo4{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo4) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo4) Close() {
	close(memo.requests)
}

func (memo *Memo4) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
