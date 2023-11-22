package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"

	proto "handin5/grpc"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id int
	leaderPort int
	backupport1 int
	backupport2 int
}

var (
	leaderPort = flag.Int("port", 0, "server port number")
	id = flag.Int("id", 0, "client id number")
	backupport1 = flag.Int("bp1", 0, "backup-port 1")
	backupport2 = flag.Int("bp2", 0, "backup-port 2")
)

func main() {
	// Parse the flags to get the port for the client
	flag.Parse()

	// Create seperate logfile
	logfile, err := os.OpenFile("clientAuction.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		log.Fatal("Could not open clientAction.log")
	}
	defer logfile.Close()
	mw := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(mw)
	log.SetPrefix(fmt.Sprint("Client ", *id, ": "))


	// Create a client
	client := &Client{
		id: *id,
		leaderPort: *leaderPort,
		backupport1: *backupport1,
		backupport2: *backupport2,
	}

	go Auction(client)

	for {
		
	}
}

func connectToServer() (proto.AuctionClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*leaderPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *leaderPort)
	} else {
		log.Printf("Connected to the server at port %d\n", *leaderPort)
	}
	return proto.NewAuctionClient(conn), nil
}

func Auction(client *Client) {
	// Connect to the server
	serverConnection, connErr := connectToServer()
	if (connErr != nil){
		log.Printf("serverconnection failed")
		log.Fatal(connErr.Error())
		
	}

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		numberMatch, matchErr := regexp.Compile("^[0-9]*$")
		if (matchErr != nil){
			log.Fatal(matchErr.Error())
		}
		if(numberMatch.MatchString(input)) {
			i, _ := strconv.ParseInt(input, 10, 64)
			bidMessage, err := serverConnection.Bid(context.Background(), &proto.BidRequest{
				Clientid: int64(client.id),
				Bid: i,
			})
			if (err != nil){
				warnServer()
			} else {
				log.Printf(bidMessage.Outcome)
			}

		} else if (input == "result") {
			resultMessage, err := serverConnection.Result(context.Background(), &proto.ResultRequest{Clientid: int64(client.id)})
			if (err != nil){
				warnServer()
			} else {
				log.Printf(resultMessage.Message)
			}
			
		} else {
			log.Printf("Input needs to be a bid or a request for a result")
		}

	}
}

func warnServer() {
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*backupport1), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *backupport1)
	} else {
		log.Printf("Connected to the server at port %d\n", *backupport1)
	}
	connection := proto.NewAuctionClient(conn)
	_, erro := connection.DoElection(context.Background(), &proto.ElectionWarning{})

	if erro != nil {
		conn, err := grpc.Dial("localhost:"+strconv.Itoa(*backupport2), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Could not connect to port %d", *backupport2)
		} else {
			log.Printf("Connected to the server at port %d\n", *backupport2)
		}
		connection := proto.NewAuctionClient(conn)
		connection.DoElection(context.Background(), &proto.ElectionWarning{})
	}
}