package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	// this has to be the same as the go.mod module,
	// followed by the path to the folder the proto file is in.
	distIncrement "github.com/suvihanninen/DistIncrement.git/grpc"
	"google.golang.org/grpc"
)

type FEServer struct {
	distIncrement.UnimplementedIncrementValueServer        // You need this line if you have a server
	port                                            string // Not required but useful if your server needs to know what port it's listening to
	primaryServer                                   distIncrement.IncrementValueClient
	ctx                                             context.Context
}

var serverToDial int

func main() {
	port := os.Args[1] //give it a port and input the same port to the client
	address := ":" + port
	list, err := net.Listen("tcp", address)

	if err != nil {
		log.Printf("FEServer %s: Server on port %s: Failed to listen on port %s: %v", port, port, address, err) //If it fails to listen on the port, run launchServer method again with the next value/port in ports array
		return
	}
	grpcServer := grpc.NewServer()

	//log to file instead of console
	f := setLogFEServer()
	defer f.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := &FEServer{
		port:          os.Args[1],
		primaryServer: nil,
		ctx:           ctx,
	}

	distIncrement.RegisterIncrementValueServer(grpcServer, server) //Registers the server to the gRPC server.

	log.Printf("FEServer %s: Server on port %s: Listening at %v\n", server.port, port, list.Addr())
	go func() {
		log.Printf("FEServer %s: We are trying to listen calls from client: %s", server.port, port)

		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to serve %v", err)
		}

		log.Printf("FEServer %s: We have started to listen calls from client: %s", server.port, port)
	}()

	serverToDial = 5001
	conn := server.DialToPR(serverToDial)

	defer conn.Close()

	for {
	}

}

func (FE *FEServer) DialToPR(serverToDial int) *grpc.ClientConn {
	//Dialing to the primary replica manager
	portToDial := ":" + strconv.Itoa(serverToDial)
	connection, err := grpc.Dial(portToDial, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	log.Printf("FEServer %s: Connection established with Primary Replica.", FE.port)
	primServer := distIncrement.NewIncrementValueClient(connection)
	FE.primaryServer = primServer
	return connection
}

func (FE *FEServer) Increment(ctx context.Context, AddRequest *distIncrement.AddRequest) (*distIncrement.Response, error) {
	result, err := FE.primaryServer.Increment(ctx, AddRequest)
	if err != nil {
		log.Printf("FEServer %s: Error: %s", FE.port, err)
		//if we get an error we need to Dial to another port
		redialOutcome := FE.Redial(AddRequest.GetValue())
		result = &distIncrement.Response{Response: redialOutcome}
	}

	return result, nil
}

func (FE *FEServer) Redial(value int32) int32 {

	portNumber := int64(serverToDial) + int64(1)

	log.Printf("FEServer %s: Dialing to new PrimaryReplica on port ", FE.port, portNumber)
	FE.DialToPR(int(portNumber))
	//redial

	AddRequest := &distIncrement.AddRequest{
		Value: value,
	}
	result, _ := FE.Increment(FE.ctx, AddRequest)
	return result.Response

}

// sets the logger to use a log.txt file instead of the console
func setLogFEServer() *os.File {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return f
}
