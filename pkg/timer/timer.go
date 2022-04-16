package timer

type Timer struct {
	Name      string
	Seconds   int64
	Frequency int64
}

var (
	timersMap = make(map[string]chan *Timer, 0)
)

func New(name string, seconds int64, frequency int64) *Timer {
	return &Timer{
		Name:      name,
		Seconds:   seconds,
		Frequency: frequency,
	}
}

// Return channel with timer values if timer exists
// Create new otherwise
func (timer *Timer) Serve() chan *Timer {
	timerCh, ok := timersMap[timer.Name]
	if !ok {
		// Create new timer channel
		timer.StartTimer()
	}

	return timerCh
}

// Start timer goroutine
func (timer *Timer) StartTimer()  {
	
}
