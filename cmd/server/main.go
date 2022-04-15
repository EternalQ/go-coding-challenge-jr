package main

import (
	"challenge/pkg/proto"
	"challenge/pkg/server"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

//Configuring env variables
func configureViper() error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := configureViper(); err != nil {
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()
	srv := &server.ChallengeServer{}
	proto.RegisterChallengeServiceServer(s, srv)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
