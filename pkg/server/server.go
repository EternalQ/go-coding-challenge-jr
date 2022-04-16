package server

import (
	"challenge/pkg/proto"
	"challenge/pkg/shortener"
	"challenge/pkg/timer"
	"context"
)

type ChallengeServer struct {
	proto.UnimplementedChallengeServiceServer
}

func (s *ChallengeServer) MakeShortLink(ctx context.Context, link *proto.Link) (*proto.Link, error) {
	shortLink, err := shortener.GetBitlyShorten(link.Data)
	if err != nil {
		return nil, err
	}
	return &proto.Link{Data: shortLink}, nil
}

func (s *ChallengeServer) StartTimer(ptimer *proto.Timer, stream proto.ChallengeService_StartTimerServer) error {
	timer := timer.New(ptimer.Name, ptimer.Seconds, ptimer.Frequency)
	timerChan := timer.Serve()
	for v := range timerChan {
			stream.Send(&proto.Timer{
				Name: v.Name,
				Seconds: v.Seconds,
				Frequency: v.Frequency,
			})
	}
	
	return nil
}

func (s *ChallengeServer) ReadMetadata(ctx context.Context, placeholder *proto.Placeholder) (*proto.Placeholder, error) {
	return &proto.Placeholder{}, nil
}
