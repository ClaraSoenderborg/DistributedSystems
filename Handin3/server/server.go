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
    streams []proto.ChittyChat_BroadcastServer
}

var port = flag.Int("port", 0, "server port number")

func (s *Server) Broadcast(msgStream proto.ChittyChat_BroadcastServer) error {
    
    s.streams = append(s.streams, msgStream)
    
    s.logOn(msgStream)
    go s.receive(msgStream)
    
    for {
        
    }
    return nil
}

func (s *Server) receive(stream proto.ChittyChat_BroadcastServer) error {
    for {
        msg, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        go s.broadcastStreams(msg)
    }
    return nil
    
}

func (s *Server) broadcastStreams(msg *proto.ClientMessage) error {
        for _, stream := range s.streams {
            if err := stream.Send(msg); err != nil {
                return err
            }
            log.Printf("vi sender %s fra user %d", msg.GetMessage(), msg.GetClientId())
            
        }
    return nil
}

func (s *Server) logOn(recv proto.ChittyChat_BroadcastServer) error{
    log.Printf("start logon")
    msg, err := recv.Recv()
    if err == io.EOF {
        return err
    }
    
    for _, stream := range s.streams {
        stream.Send(&proto.ClientMessage{
			ClientId: int64(msg.GetClientId()),
			Message: string("has joined"),
			Timestamp: int64(msg.GetTimestamp()),
		})
    }
    

 return nil
    
}

func (s *Server) logOff(msg proto.ClientMessage){
    for _, stream := range s.streams {
        stream.Send(&proto.ClientMessage{
			ClientId: int64(msg.GetClientId()),
			Message: string("has left"),
			Timestamp: int64(msg.GetTimestamp()),
		})
    }
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