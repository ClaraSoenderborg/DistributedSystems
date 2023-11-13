package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"math/rand"
	proto "mutex/grpc"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
type Node struct {
	proto.UnimplementedMutexServer // Necessary
	id                              int
	port                            int
	ports 							[]proto.MutexClient
	state							string
	timestamp						int
	queue							chan int
}

var port = flag.Int("port", 0, "server port number")
var id = flag.Int("id", 0, "node id")

func main() {
	// Get the port from the command line when the server is run
	flag.Parse()

	// Create a server struct
	node := &Node{
		id: *id,
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

	n.state = "RELEASED"
	n.timestamp = 0
	n.queue = make(chan int)

	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan(){
		input := scanner.Text()
		if (input == "start"){
			n.connectToPeer()

			y:=rand.Intn(10)
			if(y>2){
				log.Printf("Random number: %d",y)
				n.enter()
			}
		}
	}


	
	

}

func (n *Node)connectToPeer() error {
	log.Printf("Hey girls, i want to connect")
	for i := 5000; i <= 5005; i++{
		if(*port != i){
			// Dial the server at the specified port.
			conn, err := grpc.Dial("localhost:"+strconv.Itoa(i), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Printf("Could not connect to port %d", i)
			} else {
				log.Printf("Connected to the server at port %d\n", i)
				n.ports = append(n.ports, proto.NewMutexClient(conn))
			}
		}
	}

	return nil
} 

//This is the recieving function
func (n *Node) mutex(ctx context.Context, request proto.Request)(*proto.Reply, error){
	n.timestamp =+1

	if (n.state == "HELD" || (n.state == "WANTED" && (n.timestamp < int(request.Timestamp)))){
		<- n.queue

	} 
	return &proto.Reply{
		Timestamp: int64(n.timestamp),
		ClientID: int64(n.id),
	}, nil
}

func (n *Node) enter(){
	waittime := time.Duration(rand.Intn(2000))*time.Millisecond
	time.Sleep(waittime)
	n.state = "WANTED"
	n.timestamp =+1

	for _, v := range n.ports{
		v.Mutex(context.Background(), &proto.Request{
			Timestamp: int64(n.timestamp),
			ClientID: int64(n.id),
		})
	}

	n.state = "HELD"
	n.exit()
}

func (n *Node) exit(){
	log.Printf("I am node nr.%d and i am doing a critical thing", n.id)
	n.state = "RELEASED"
	n.timestamp =+1

	n.queue <- 1
}

