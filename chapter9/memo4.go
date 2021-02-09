package chapter9

import "sync"

type entry struct {
	res   result
	ready chan struct{}
}

type Memo3 struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func New2(f Func) *Memo3 {
	return &Memo3{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo3) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
