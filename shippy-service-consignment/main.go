package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/mdesson/shippy/shippy-service-consignment/proto/consignment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository - Dummy repo, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc.
// to give you a better idea
type service struct {
	repo repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return matching `Response` message  we create din our protobuf definition
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	repo := &Repository{}

	// set up gRPC server
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %w\n", err)
	}
	server := grpc.NewServer()

	// Regiser our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protbuf defintion
	pb.RegisterShippingServiceServer(server, &service{repo})

	// Register reflection service on gRPC server
	reflection.Register(server)

	log.Println("Running on port: ", port)
	if err := server.Serve(listen); err != nil {
		log.Fatal("failed to serve: %w\n", err)
	}
}
