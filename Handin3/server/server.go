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
    Timestamp int64
}

var port = flag.Int("port", 0, "server port number")

func (s *Server) Broadcast(msgStream proto.ChittyChat_BroadcastServer) error {
    
    s.streams = append(s.streams, msgStream)
    
    _, clientid := s.logOn(msgStream)
    go s.receive(msgStream)

    <-msgStream.Context().Done()
    
    s.logOff(clientid, msgStream)
    
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
        if(msg.Timestamp > s.Timestamp){
            s.Timestamp = msg.Timestamp+1
        }else{
            s.Timestamp += 1
        }
        log.Printf("Serverlog: User %d @%d :%s", msg.ClientId, s.Timestamp, msg.Message)
        go s.broadcastStreams(msg)
    }
    return nil
    
}

func (s *Server) broadcastStreams(msg *proto.ClientMessage) error {
        for _, stream := range s.streams {
            s.Timestamp += 1
            if err := stream.Send(&proto.ClientMessage{
                ClientId: int64(msg.GetClientId()),
                Message: string(msg.GetMessage()),
                Timestamp: int64(s.Timestamp),
            }); err != nil {
                return err
            }
            
            log.Printf("vi sender %s fra user %d", msg.GetMessage(), msg.GetClientId())
        }
    return nil
}

func (s *Server) logOn(recv proto.ChittyChat_BroadcastServer) (error, int){
    log.Printf("start logon")
    msg, err := recv.Recv()
    if err == io.EOF {
        return err, -1
    }

    if(msg.Timestamp > s.Timestamp){
        s.Timestamp = msg.Timestamp+1
    }else{
        s.Timestamp += 1
    }


    for _, stream := range s.streams {
        s.Timestamp += 1
        stream.Send(&proto.ClientMessage{
			ClientId: int64(msg.GetClientId()),
			Message: string("has joined"),
			Timestamp: int64(s.Timestamp),
		})
    }
    

 return nil, int(msg.ClientId)
    
}

func (s *Server) logOff(clientid int, stream proto.ChittyChat_BroadcastServer){
    s.streams = delete(s.streams, stream)
    for _, stream := range s.streams {
        s.Timestamp += 1
        stream.Send(&proto.ClientMessage{
			ClientId: int64(clientid),
			Message: string("has left"),
			Timestamp: int64(s.Timestamp),
		})
    }
}

func delete(slice []proto.ChittyChat_BroadcastServer, stream proto.ChittyChat_BroadcastServer) []proto.ChittyChat_BroadcastServer{
    var index int 
    for i, element := range slice{
        if (element == stream){
            index = i
        }
    }

    return append(slice[:index], slice[index+1:]...)

}


func newServer(server *Server) {
    ChittyChat := grpc.NewServer()

    listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))
    server.Timestamp = 0
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