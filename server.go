package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/reoxey/learn-grpc/chat"
)

func main() {

	l, e := net.Listen("tcp", ":9000")
	if e != nil {
		log.Fatalln(e)
	}

	grpcOb := grpc.NewServer()

	s := chat.Server{}

	chat.RegisterChatServiceServer(grpcOb, &s)

	if e := grpcOb.Serve(l); e != nil {
		log.Fatalln(e)
	}

	log.Println("All is well")
}
