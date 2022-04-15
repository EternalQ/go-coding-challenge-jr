package server

import (
	"challenge/pkg/proto"
	"challenge/pkg/shortener"
	"context"
)

type ChallengeServer struct {
	proto.UnimplementedChallengeServiceServer
}

func (s *ChallengeServer) MakeShortLink(ctx context.Context, link *proto.Link) (*proto.Link, error) {
	shortLink, err := shortener.BitlyShortener(link.Data)
	if err != nil {
		return nil, err
	}
	return &proto.Link{Data: shortLink}, nil
}

func (s *ChallengeServer) StartTimer(timer *proto.Timer, stream proto.ChallengeService_StartTimerServer) error {
	return nil
}

func (s *ChallengeServer) ReadMetadata(ctx context.Context, placeholder *proto.Placeholder) (*proto.Placeholder, error) {
	return &proto.Placeholder{}, nil
}
