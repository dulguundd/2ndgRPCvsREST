package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "goproject/2ndgRPCvsREST/gRPCbenchmark/usermgmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	start := time.Now()
	log.Printf("Recieved: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	serviceLatencyLogger(start)
	return &pb.User{
		Name:        in.GetName(),
		Age:         in.GetAge(),
		Id:          user_id,
		Inactendd:   "01/01/2038",
		Error:       "",
		Retailer:    "2",
		Class:       "PRE_Hybrid_14900_N",
		Actendd:     "01/01/2038",
		Adminst:     "1",
		Creditvioce: "5490500",
		Code:        "0",
		Phone:       "94300048",
		Rbal:        "5490500",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func serviceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	log.Println(logMessage)
}
