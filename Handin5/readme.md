<h1>How to run the program </h1>

First run 3 servers, in different terminals, where the first will be the primary replica
1. Run the primary replica: `go run server/server.go -port 5000 -id 1 -lp 5000 -bp1 5001 -bp2 5002`
2. Run the backup replica: `go run server/server.go -port 5001 -id 2 -lp 5000 -bp1 5001 -bp2 5002`
3. Run the backup replica: `go run server/server.go -port 5002 -id 1 -lp 5000 -bp1 5001 -bp2 5002`



Then run 2 clients in different terminals
1. Run client `go run client/client.go -port 5000 -id 1`
2. Run client `go run client/client.go -port 5000 -id 2`

