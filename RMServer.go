package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	// this has to be the same as the go.mod module,
	// followed by the path to the folder the proto file is in.
	distIncrement "github.com/suvihanninen/DistIncrement.git/grpc"
	"google.golang.org/grpc"
)

type RMServer struct {
	distIncrement.UnimplementedIncrementValueServer
	id        int32
	peers     map[int32]distIncrement.IncrementValueClient
	isPrimary bool
	ctx       context.Context
	time      time.Time
	primary   distIncrement.IncrementValueClient
	value     int32
}

func main() {

	portInput, _ := strconv.ParseInt(os.Args[1], 10, 32) //Takes arguments 0, 1 and 2, see comment X
	primary, _ := strconv.ParseBool(os.Args[2])
	ownPort := int32(portInput) + 5001
	println("primary?: "+strconv.FormatBool(primary)+" port?: ", ownPort)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rmServer := &RMServer{
		id:        ownPort,
		peers:     make(map[int32]distIncrement.IncrementValueClient),
		ctx:       ctx,
		isPrimary: primary,
		value:     -1,
		primary:   nil,
	}

	//log to file instead of console
	f := setLogRMServer()
	defer f.Close()

	//Primary needs to listen so that replica managers can ask if it's alive
	//Replica managers need to listen for incoming data to be replicated
	println("Making connection in order to listen other replicas")
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	println("Registered the new server")
	grpcServer := grpc.NewServer()
	distIncrement.RegisterIncrementValueServer(grpcServer, rmServer)
	println("Server listening on port: ", ownPort)
	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()
	println("Connecting to other peers")
	for i := 0; i < 3; i++ {
		port := int32(5001) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		log.Printf("RMServer %v: Trying to dial: %v\n", rmServer.id, port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock()) //This is going to wait until it receives the connection
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := distIncrement.NewIncrementValueClient(conn)
		rmServer.peers[port] = c
		if port == 5001 {
			log.Printf("We have assigned a primary replica")
			rmServer.primary = c
		}
	}
	println("Connected to the peers")
	// If server is primary, dial to all other replica managers
	if rmServer.isPrimary {
		println("Waiting for heartbeat request")
	} else {
		println("Sending heartbeat requests to PR")
		go func() {
			for {
				time.Sleep(2 * time.Second)
				heartbeatMsg := &distIncrement.BeatRequest{Message: "alive?"}
				_, err := rmServer.primary.GetHeartBeat(rmServer.ctx, heartbeatMsg)
				if err != nil {
					log.Printf("RMServer %v: Something went wrong while sending heartbeat", rmServer.id)
					log.Printf("RMServer %v: Error:", rmServer.id, err)
					log.Printf("RMServer %v: Exception, We did not get heartbeat back from Primary Replica with port %v. It has died, ", rmServer.id)
					delete(rmServer.peers, 5001)
					rmServer.ElectLeader()
				}
				//comment out if you wanna see the logging of heartbeat
				//log.Printf("We got a heart beat from %s", response)

				if rmServer.isPrimary {
					break
				}
			}
		}()
		for {

		}
	}
	for {

	}

}

func (RM *RMServer) ElectLeader() {
	log.Printf("RMServer %v: Leader election started with Bully Algorithm", RM.id)
	var min int32
	min = RM.id
	for id := range RM.peers {
		if min > id {
			min = id
		}
	}

	if RM.id == min {
		RM.isPrimary = true
	} else {
		RM.primary = RM.peers[min]
	}
	log.Printf("RMServer %v: New Primary Replica has port %v ", RM.id, min)
}

func (RM *RMServer) GetHeartBeat(ctx context.Context, Heartbeat *distIncrement.BeatRequest) (*distIncrement.BeatAck, error) {
	return &distIncrement.BeatAck{Port: fmt.Sprint(RM.id)}, nil
}

func (RM *RMServer) Increment(ctx context.Context, AddRequest *distIncrement.AddRequest) (*distIncrement.Response, error) {
	result, err := RM.addValueToRM(AddRequest.GetValue())
	if err != nil {
		log.Fatalf("Adding value to Replica Managers failed inside RMServer: %s", err)
	}

	return &distIncrement.Response{Response: result}, nil
}

func (rm *RMServer) addValueToRM(value int32) (int32, error) {
	//Update Primary Replica
	rm.value = rm.value + 1
	addValue := &distIncrement.AddRequest{Value: value}

	//Broadcasting the Increment request to all replica managers
	for id, server := range rm.peers {
		_, err := server.AddToValue(rm.ctx, addValue)
		if err != nil {
			log.Printf("RMServer %v: Something went wrong when adding value to %v", rm.id, id)
			log.Printf("RMServer %v: Exception, Replica Manager on port %v died", rm.id, id)
			delete(rm.peers, id)
		}

		log.Printf("RMServer %v: Value added to replica manager on port %s: ", rm.id, id)
	}

	return rm.value, nil
}

func (RM *RMServer) AddToValue(ctx context.Context, AddRequest *distIncrement.AddRequest) (*distIncrement.Response, error) {
	RM.value = RM.value + 1

	println("AddToValue: Added value. New value on RM: ", RM.value)
	log.Printf("RMServer %v: Added value. New value on RM: ", RM.id, RM.value)

	return &distIncrement.Response{Response: RM.value}, nil
}

func setLogRMServer() *os.File {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return f
}
