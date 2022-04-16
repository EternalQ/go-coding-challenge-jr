package server

import (
	"challenge/pkg/proto"
	"challenge/pkg/shortener"
	"challenge/pkg/timer"
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
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

	pipe := timer.GetPipe()

	for {
		select {
		case timer = <-pipe.Timers:
			stream.Send(&proto.Timer{
				Name:    timer.Name,
				Seconds: timer.Seconds,
			})
		case err := <-pipe.Errors:
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
