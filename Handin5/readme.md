<h1>How to run the program </h1>

First run 3 servers, in different terminals, where the first will be the primary replica
1. Run the primary replica: `go run server/server.go -port 5003 -id 1 -lp 5003 -bp1 5001 -bp2 5002`
2. Run the backup replica: `go run server/server.go -port 5001 -id 2 -lp 5003 -bp1 5001 -bp2 5002`
3. Run the backup replica: `go run server/server.go -port 5002 -id 1 -lp 5003 -bp1 5001 -bp2 5002`



Then run 2 clients in different terminals
1. Run client `go run client/client.go -port 5003 -id 1 -bp1 5001 -bp2 5002`
2. Run client `go run client/client.go -port 5003 -id 2 -bp1 5001 -bp2 5002`

To simulate a crash of a node, press `ctrl` + `c` in the first terminal, that is running the primary replica. Next call from a client will now start an election process. 

OBS: send the call from client again, when connected to the new server to finalize the bid or request. 
