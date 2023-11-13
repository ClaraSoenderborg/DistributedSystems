<h1>How to run the program </h1>

In seperate terminals, run the followng:
1. Run the node: `go run server/server.go -port 5001 -id 1`
2. Run the node: `go run server/server.go -port 5002 -id 2`
3. Run the node: `go run server/server.go -port 5003 -id 3`

In each terminal, write `start` and execute. 
This will make each node connect to one another. They will each wait 10 minutes before starting to run the algorithm. 
