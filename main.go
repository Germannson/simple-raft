// Simple implementation of the raft concensus algorithm using golang rpc
// Cody Malick and Jacob Broderick
package main

import (
	"flag"
)

// Reference: https://golang.org/pkg/net/rpc/
const (
	numServers = 2
)

// Leader Election: Choose a leader
// Log replication: Make sure all systems have the same view of the log
// Safety: Make sure a leader who is behind cannot be elected

func main() {
	cmdID := flag.Int("id", 0, "Usage: -id <id>")
	cmdPort := flag.String("port", ":50000", "Usage: -port <portnumber>")
	cmdState := flag.Int("state", 0, "Usage: -state <start-state>")
	flag.Parse()

	server := CreateServer(*cmdID, *cmdPort, *cmdState, make(chan bool))
	exit := make(chan bool)


	// Server 0 is our test leader
	server0 := CreateServer(0, ":50000", 0, nil)
	server1 := CreateServer(1, ":50001", 0, nil)
	server2 := CreateServer(2, ":50002", 0, nil)
	server3 := CreateServer(3, ":50003", 0, nil)
	server4 := CreateServer(4, ":50004", 0, nil)

	servers := []*Server{server0,server1,server2,server3,server4}

	server.Servers = servers

	// Spawn goroutine
	go Run(server)


	// Wait for anything on the exit channel to quit
	<- exit
}
