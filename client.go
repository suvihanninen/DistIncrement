package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	// this has to be the same as the go.mod module,
	// followed by the path to the folder the proto file is in.
	distIncrement "github.com/suvihanninen/DistIncrement.git/grpc"
	"google.golang.org/grpc"
)

func main() {

	//Make connection to 5001
	port := ":" + os.Args[1]
	connection, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}

	//log to file instead of console
	f := setLogClient()
	defer f.Close()

	server := distIncrement.NewIncrementValueClient(connection) //creates a new client

	defer connection.Close()

	go func() {

		for i := 0; i < 20; i++ {
			min := 1
			max := 4
			delay := rand.Intn(max-min) + min
			time.Sleep(time.Duration(delay) * time.Second)

			addRequest := &distIncrement.AddRequest{
				Value: int32(1),
			}

			result := IncrementValue(addRequest, connection, server, port)

			log.Printf("The value: ", result)
			println("The value: ", result)

		}
	}()

	for {

	}

}

func IncrementValue(addRequest *distIncrement.AddRequest, connection *grpc.ClientConn, server distIncrement.IncrementValueClient, port string) int32 {
	var result int32

	response, err := server.Increment(context.Background(), addRequest)
	if err != nil {
		log.Printf("Client %s: Adding failed: ", port, err)
		log.Printf("Client %s: FEServer has died", port)
		connection, server = Redial(port)
		result = IncrementValue(addRequest, connection, server, port)
		return result
	}
	result = response.GetResponse()
	return result
}

func Redial(port string) (*grpc.ClientConn, distIncrement.IncrementValueClient) {
	log.Printf("Client: FEServer on port %s is not listening anymore. It has died", port)
	if port == ":4001" {
		port = ":4002"
	} else {
		port = ":4001"
	}
	log.Printf("Client: Redialing to new port: " + port)
	connection, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}

	server := distIncrement.NewIncrementValueClient(connection) //creates a new client
	log.Printf("Client: Client has connected to new FEServer on port %s", port)
	return connection, server
}

// sets the logger to use a log.txt file instead of the console
func setLogClient() *os.File {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return f
}
