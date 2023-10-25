package main

import (
	proto "ChittyChat/grpc"
	"flag"
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Server struct {
    name string
	proto.UnimplementedChittyChatServer
	port int
}

var port = flag.Int("port", 0, "server port number")

func (s *Server) Broadcast(msgStream proto.ChittyChat_BroadcastServer) error {
    // get all the messages from the stream
    for {
        msg, err := msgStream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        if err := msgStream.Send(msg); err != nil {
            return err
        }
        
        }
	// every time we get a message from a client, we should immediately broadcast it to all clients (goroutine)

	// how does server know a client has joined? 

    //ack := proto.ServiceBroadcast// make an instance of your return type
    //msgStream.SendAndClose(ack)

    return nil
}

func newServer(server *Server) {
    ChittyChat := grpc.NewServer()

    listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

    if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

    proto.RegisterChittyChatServer(ChittyChat, server)
	serveError := ChittyChat.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func main() {
    flag.Parse()

	server := &Server{
		name: "ChittyChat",
		port: *port,
	}

	go newServer(server)

	for {

	}
}