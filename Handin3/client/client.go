package main

import (
	proto "ChittyChat/grpc"
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strconv"
	

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id int
}

var (
	port = flag.Int("port", 0, "server port number")
	id = flag.Int("id", 0, "client id")
)

func publishMessage(client *Client){
	serverConnection, _ := connectToServer()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
	
		stream, err := serverConnection.Broadcast(context.Background())

		stream.Send(&proto.ClientMessage{
			ClientId: int64(client.id),
			Message: string(input),
			Timestamp: int64(1),
		})

		in, err := stream.Recv()

		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Printf("User %d @%d :%s", in.ClientId, in.Timestamp, in.Message)
		}
	}
	
}

func connectToServer() (proto.ChittyChatClient, error) {
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *port)
	} else {
		log.Printf("Connected to the server at port %d\n", *port)
	}
	return proto.NewChittyChatClient(conn), nil
}

func main() {
	flag.Parse()

	client := &Client{
		id: *id,
	}

	go publishMessage(client)

	for {
		
	}
}