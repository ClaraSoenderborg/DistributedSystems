<h1>How to run the program </h1>

1. Run the server: `go run server/server.go -port 5454`
2. In a different terminal, run a client: `go run client/client.go -port 5454 -id 1`
3. To add more clients, open a new terminal and run with a new id value: `go run client/client.go -port 5454 -id 2`

Now clients can write messages in ChittyChat 
