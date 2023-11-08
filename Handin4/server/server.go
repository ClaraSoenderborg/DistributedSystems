package main

import (
	"context"
	"flag"
	"log"
	proto "mutex/grpc"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
type Node struct {
	proto.UnimplementedMutexServer // Necessary
	name                             string
	port                             int
	ports []proto.MutexClient
}

var port = flag.Int("port", 0, "server port number")

func main() {
	// Get the port from the command line when the server is run
	flag.Parse()

	// Create a server struct
	node := &Node{
		name: "serverName",
		port: *port,
	}

	// Start the server
	go startNode(node)

	// Keep the server running until it is manually quit
	for {

	}
}

func startNode(n *Node) {

	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(n.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", n.port)

	// Register the grpc server and serve its listener
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}

}

func (n *Node)connectToPeer() error {

	for i := 5000; i <= 5005; i++{
		if(*port != i){
			// Dial the server at the specified port.
			conn, err := grpc.Dial("localhost:"+strconv.Itoa(i), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Could not connect to port %d", i)
			} else {
				log.Printf("Connected to the server at port %d\n", i)
				n.ports = append(n.ports, proto.NewMutexClient(conn))
			}
		}
	}

	return nil
} 

func (n *Node) mutex(ctx context.Context, request proto.Request)(*proto.Reply, error){
	
}