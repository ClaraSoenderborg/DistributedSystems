package main

import (
	"flag"
	proto "handin5/grpc"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Node struct {
	proto.UnimplementedAuctionServer
	id                             	int
	port                            int
	leader							bool
	ports							[]int
}

var port = flag.Int("port", 0, "node port number")
var id = flag.Int("id", 0, "node id number")
var leaderport = flag.Int("leaderport", 0, "leader port number")
var backupport1 = flag.Int("backupport1", 0, "backup 1 port number")
var backupport2 = flag.Int("backupport2", 0, "backup 2 port number")

func main(){
	flag.Parse()


	// Create a server struct
	node := &Node{
		id: *id,
		port: *port,
		leader: *leaderport == *port,
		ports: []int{*leaderport,*backupport1,*backupport2},
	}

	go startNode(node)

	for{

	}
}

func startNode(node *Node) {
	grpcNode := grpc.NewServer()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(node.port))

	if err != nil{
		log.Fatalf("could not create the node %v", err)
	}
	log.Printf("started node at port: %d \n", node.port)

	proto.RegisterAuctionServer(grpcNode, node)
	serveError:= grpcNode.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func (node *Node) connectToPeers() error {
	for _, currentport:= range node.ports{
		if(node.port != currentport){
			// Dial the server at the specified port.
			conn, err := grpc.Dial("localhost:"+strconv.Itoa(i), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Printf("Could not connect to port %d", i)
			} else {
				log.Printf("Connected to the server at port %d\n", i)
				
			}
		}
	}

	return nil
}

