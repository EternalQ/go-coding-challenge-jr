package timer

import (
	"time"
)

type Timer struct {
	Name      string
	Seconds   int64
	Frequency int64
}

type timerPipe struct {
	Timers chan *Timer
	Errors chan error
}

var (
	timersMap = make(map[string]*timerPipe)
)

func New(name string, seconds int64, frequency int64) *Timer {
	return &Timer{
		Name:      name,
		Seconds:   seconds,
		Frequency: frequency,
	}
}

// Return channel with timer values if exists
// Create new otherwise
func (timer *Timer) GetPipe() *timerPipe {
	pipe, ok := timersMap[timer.Name]
	if !ok {
		pipe = timer.StartTimer()
	}

	return pipe
}

// Start timer goroutine
func (timer *Timer) StartTimer() *timerPipe {
	StartTimerAPI(timer.Name, timer.Seconds)

	pipe := &timerPipe{
		Timers: make(chan *Timer),
		Errors: make(chan error),
	}
	timersMap[timer.Name] = pipe

	// General timer
	stop := make(chan bool)
	go func() {
		time.Sleep(time.Second * time.Duration(timer.Seconds))
		stop <- true
	}()

	ticker := time.NewTicker(time.Second * time.Duration(timer.Frequency))

	// Updating timer info
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				timerResp, err := CheckTimerAPI(timer.Name)
				if err != nil {
					pipe.Errors <- err
					continue
				}

				timer.Name = timerResp.Name
				timer.Seconds = int64(timerResp.Seconds)

				pipe.Timers <- timer
			case <-stop:
				close(pipe.Errors)
				close(pipe.Timers)
				delete(timersMap, timer.Name)
				return
			}
		}
	}()

	return pipe
}
