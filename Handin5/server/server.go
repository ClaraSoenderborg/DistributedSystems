package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	proto "handin5/grpc"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Node struct {
	proto.UnimplementedAuctionServer
	id                             	int
	port                            int
	leader							bool
	ports							[]int
	connections						[]proto.AuctionClient
	highestBid						int
	highestBidder					int
	active							bool
}

var port = flag.Int("port", 0, "node port number")
var id = flag.Int("id", 0, "node id number")
var leaderport = flag.Int("lp", 0, "leader port number")
var backupport1 = flag.Int("bp1", 0, "backup 1 port number")
var backupport2 = flag.Int("bp2", 0, "backup 2 port number")

func main(){
	flag.Parse()

	// Create seperate logfile
	logfile, err := os.OpenFile("serverAuction.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		log.Fatal("Could not open serverAction.log")
	}
	defer logfile.Close()
	mw := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(mw)

	// Create a server struct
	node := &Node{
		id: *id,
		port: *port,
		leader: *leaderport == *port,
		ports: []int{*leaderport,*backupport1,*backupport2},
		active: true,
		highestBid: 0,
		highestBidder: -1,
	}

	go startNode(node)
	go runNode(node)
	time.Sleep(time.Duration(60*time.Second))
	node.active = false
	log.Printf("Auction is over!")

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

func runNode (node *Node) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		input := scanner.Text()
		log.Printf(input)
		if (input == "start"){
			node.connectToPeers()
		}
	}
}

func (node *Node) connectToPeers() error {
	for _, currentport:= range node.ports{
		if(node.port != currentport){
			// Dial the server at the specified port.
			conn, err := grpc.Dial("localhost:"+strconv.Itoa(currentport), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Printf("Could not connect to port %d", currentport)
			} else {
				log.Printf("Connected to the server at port %d\n", currentport)
				node.connections = append(node.connections, proto.NewAuctionClient(conn))
			}
		}
	}
	return nil
}

func (node *Node) Bid(ctx context.Context, in *proto.BidRequest) (*proto.BidAck, error) {
	// Exception
	if(!node.active) {
		return &proto.BidAck{Outcome: string("The auction has closed")}, nil
	}
	
	// Failure
	if(in.Bid <= int64(node.highestBid)) {
		return &proto.BidAck{Outcome: string("Your bid is not high enough")}, nil
	}
	node.highestBid = int(in.Bid)
	node.highestBidder = int(in.Clientid)

	for _, conn:= range node.connections{
		ack, _ := conn.InternalBid(ctx, in)
		if(ack.Outcome != "success") {
			node.connections = delete(node.connections, conn)
		}
	}

	// Success
	returnString := fmt.Sprintf("Your bid on %d has been accepted!", in.Bid)
	return &proto.BidAck{Outcome: string(returnString)}, nil
}

func delete(slice []proto.AuctionClient, deletion proto.AuctionClient) []proto.AuctionClient{
    var index int 
    for i, element := range slice{
        if (element == deletion){
            index = i
        }
    }
    return append(slice[:index], slice[index+1:]...)
}

func (node *Node) InternalBid(ctx context.Context, in *proto.BidRequest) (*proto.BidAck, error) {
	node.highestBid = int(in.Bid)
	node.highestBidder = int(in.Clientid)

	return &proto.BidAck{Outcome: string("success")}, nil
}

func (node *Node) Result(ctx context.Context, res *proto.ResultRequest) (*proto.OutcomeResponse, error) {
	if(node.active) {
		activeResult := fmt.Sprintf("The auction is still running and the highest bid is %d", node.highestBid)
		return &proto.OutcomeResponse{Message: string(activeResult)}, nil
	} 
	doneResult := fmt.Sprintf("The auction is over and client %d won with the highest bid %d", node.highestBidder, node.highestBid)
		return &proto.OutcomeResponse{Message: string(doneResult)}, nil
}
