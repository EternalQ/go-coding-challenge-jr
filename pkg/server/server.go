package server

import (
	"challenge/pkg/proto"
	"challenge/pkg/shortener"
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
)

type ChallengeServer struct {
	proto.UnimplementedChallengeServiceServer
	Brokers map[string]*Broker
}

func (s *ChallengeServer) getBroker(ptimer *proto.Timer) *Broker {
	t := New(ptimer.Name, ptimer.Seconds, ptimer.Frequency)

	b, ok := s.Brokers[t.Name]
	println(b, ok)
	if !ok {
		b = NewBroker()
		s.Brokers[ptimer.Name] = b
		go b.Start()
		t.StartTimer(b, s)
	}

	return b
}

func realeseTimer(name string, brokers map[string]*Broker) {
	delete(brokers, name)
}

func (s *ChallengeServer) MakeShortLink(ctx context.Context, link *proto.Link) (*proto.Link, error) {
	shortLink, err := shortener.GetBitlyShorten(link.Data)
	if err != nil {
		return nil, err
	}
	return &proto.Link{Data: shortLink}, nil
}

func (s *ChallengeServer) StartTimer(ptimer *proto.Timer, stream proto.ChallengeService_StartTimerServer) error {

	b := s.getBroker(ptimer)
	// println("sub", b)
	timers, errors := b.Subscribe()
	defer b.Unsubscribe(timers, errors)

	for {
		select {
		case <-b.StopCh:
			return nil
		case timer := <-timers:
			stream.Send(&proto.Timer{
				Name:    timer.Name,
				Seconds: timer.Seconds,
			})
		case err := <-errors:
			return err
		}
	}
}

func (s *ChallengeServer) ReadMetadata(ctx context.Context, placeholder *proto.Placeholder) (*proto.Placeholder, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata recived")
	}

	// Where from i should get key
	v := md.Get("somekey")

	return &proto.Placeholder{Data: strings.Join(v, ",")}, nil
}
