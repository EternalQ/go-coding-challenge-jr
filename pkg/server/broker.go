package server

type Broker struct {
	StopCh     chan struct{}
	publishCh  chan *Timer
	subCh      chan chan *Timer
	unsubCh    chan chan *Timer
	errorCh    chan error
	subErrCh   chan chan error
	unsubErrCh chan chan error
}

func NewBroker() *Broker {
	return &Broker{
		StopCh:     make(chan struct{}),
		publishCh:  make(chan *Timer, 1),
		subCh:      make(chan chan *Timer, 1),
		unsubCh:    make(chan chan *Timer, 1),
		errorCh:    make(chan error, 1),
		subErrCh:   make(chan chan error, 1),
		unsubErrCh: make(chan chan error, 1),
	}
}

func (b *Broker) Start() {
	subs := map[chan *Timer]struct{}{}
	errSubs := map[chan error]struct{}{}
	for {
		select {
		case <-b.StopCh:
			return

		case msgCh := <-b.subCh:
			subs[msgCh] = struct{}{}
		case errCh := <-b.subErrCh:
			errSubs[errCh] = struct{}{}

		case msgCh := <-b.unsubCh:
			delete(subs, msgCh)
		case errCh := <-b.unsubErrCh:
			delete(errSubs, errCh)

		case msg := <-b.publishCh:
			for msgCh := range subs {
				// msgCh is buffered, use non-blocking send to protect the broker:
				select {
				case msgCh <- msg:
				default:
				}
			}
		case err := <-b.errorCh:
			for errorCh := range errSubs {
				// msgCh is buffered, use non-blocking send to protect the broker:
				select {
				case errorCh <- err:
				default:
				}
			}
		}
	}
}

func (b *Broker) Stop() {
	close(b.StopCh)
}

func (b *Broker) Subscribe() (chan *Timer, chan error) {
	msgCh := make(chan *Timer, 5)
	errCh := make(chan error, 5)

	b.subCh <- msgCh
	b.subErrCh <- errCh

	return msgCh, errCh
}

func (b *Broker) Unsubscribe(msgCh chan *Timer, errCh chan error) {
	b.unsubCh <- msgCh
	b.unsubErrCh <- errCh
}

func (b *Broker) Publish(msg *Timer) {
	b.publishCh <- msg
}

func (b *Broker) Error(err error) {
	b.errorCh <- err
}
