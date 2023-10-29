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
	Timestamp int64
}

var (
	port = flag.Int("port", 0, "server port number")
	id = flag.Int("id", 0, "client id")
)

func publishMessage(client *Client){
	
	serverConnection, _ := connectToServer()
	stream, _ := serverConnection.Broadcast(context.Background())
	client.Timestamp=1 //timestamp is set to 1 as it is the first action we do
	stream.Send(&proto.ClientMessage{
		ClientId: int64(client.id),
		Message: string(""),
		Timestamp: int64(client.Timestamp), 
	})

	go client.receive(stream)
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		input := scanner.Text()	
		client.Timestamp += 1
		stream.Send(&proto.ClientMessage{
			ClientId: int64(client.id),
			Message: string(input),
			Timestamp: int64(client.Timestamp),
		})

	}

}

func (client *Client) receive(stream proto.ChittyChat_BroadcastClient){
	for{
	in, err := stream.Recv()

		if err != nil {
			log.Printf(err.Error())
		} else {
			if(in.Timestamp > client.Timestamp){
				client.Timestamp = in.Timestamp+1
			}else{
				client.Timestamp += 1
			}
			log.Printf("User %d @%d :%s", in.ClientId, client.Timestamp, in.Message)
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