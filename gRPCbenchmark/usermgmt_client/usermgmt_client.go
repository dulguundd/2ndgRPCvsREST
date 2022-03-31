package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "goproject/2ndgRPCvsREST/gRPCbenchmark/usermgmt"
	"log"
	"time"
)

const address = "172.22.2.215:50051"

func main() {
	start := time.Now()
	count := 1
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	start2 := time.Now()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Alice"] = 43
	new_users["Bob"] = 30

	runningcount := 0
	for runningcount < count {
		for name, age := range new_users {
			r, err := c.CreateNewUser(ctx, &pb.NewUser{
				Name: name,
				Age:  age,
			})
			if err != nil {
				log.Fatalf("could not create user: %v", err)
			}
			log.Printf(`User Detail:
NAME: %s
AGE: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())
		}
		runningcount++
	}
	serviceLatencyLogger(start2)
	serviceLatencyLogger(start)
}

func serviceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	log.Println(logMessage)
}
