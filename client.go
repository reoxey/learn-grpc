package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/reoxey/learn-grpc/chat"
)

func main() {

	ctx := context.Background()

	conn, e := grpc.DialContext(ctx, ":9000", grpc.WithInsecure())
	if e != nil {
		log.Fatalln(e)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	res, e := c.SayHello(ctx, &chat.Message{Body: "Hello There!"})
	if e != nil {
		log.Fatalln(e)
	}
	log.Println(res.Body)

}
