package main

import (
	proto "ChittyChat/grpc"
	"flag"
	"io"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChittyChatServer
	port int
}

var port = flag.Int("port", 0, "server port number")

func (s *Server) Publish(msgStream proto.ChittyChat_PublishServer) error {
    // get all the messages from the stream
    for {
        msg, err := msgStream.Recv()
        if err == io.EOF {
            break
        }
    }
	// every time we get a message from a client, we should immediately broadcast it to all clients (goroutine)

	// how does server know a client has joined? 

    ack := proto.ServiceBroadcast// make an instance of your return type
    msgStream.SendAndClose(ack)

    return nil
}