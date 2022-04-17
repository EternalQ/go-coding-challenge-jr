package server

import (
	"time"
)

type Timer struct {
	Name      string
	Seconds   int64
	Frequency int64
}

func New(name string, seconds int64, frequency int64) *Timer {
	return &Timer{
		Name:      name,
		Seconds:   seconds,
		Frequency: frequency,
	}
}

// Start timer goroutine
func (timer *Timer) StartTimer(b *Broker, s *ChallengeServer) {
	StartTimerAPI(timer.Name, timer.Seconds)

	// General timer
	stop := make(chan bool)
	go func() {
		time.Sleep(time.Second * time.Duration(timer.Seconds))
		stop <- true
	}()

	ticker := time.NewTicker(time.Second * time.Duration(timer.Frequency))

	// Timer updater
	go func() {
		defer ticker.Stop()
		defer delete(s.Brokers, timer.Name)
		for {
			select {
			case <-ticker.C:
				timerResp, err := CheckTimerAPI(timer.Name)
				if err != nil {
					b.Error(err)
					continue
				}

				timer.Name = timerResp.Name
				timer.Seconds = int64(timerResp.Seconds)

				b.Publish(timer)
			case <-stop:
				b.Stop()
				return
			}
		}
	}()
}
