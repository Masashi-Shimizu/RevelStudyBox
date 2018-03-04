package controllers

type Subscription struct {
	Archive []Event
	New <-chan Event
}

func (s Subscription) Cancel() {
	unsubscribe <- s.New
	drain(s.New)
}

func drain(ch <-chan Event) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
			default:
				return
		}
	}
}

func Subscribe() Subscription {
	resp := make(chan Subscription)
	subscribe <- resp
	reutn <- resp
}