<h1>How to run the program </h1>

In seperate terminals, run the followng:
1. Run the node: `go run server/server.go -port 5001 -id 1`
2. Run the node: `go run server/server.go -port 5002 -id 2`
3. Run the node: `go run server/server.go -port 5003 -id 3`

The nodes are now listening on their respective ports.

In each terminal, write `start` and execute. 
This will make it connect to its peer nodes. It will then generate a random number 0-9, and if it's greater than 2, it will decide that it wants to enter critical state and start sending requests to its peers. 
The `start` command can be run at any time, as long as the other nodes are alive. 
