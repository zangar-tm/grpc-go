package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/zangar-tm/grpc-go/usermanage"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Zangar"] = 20
	new_users["Bek"] = 30
	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: int32(age)})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`User details:
		Name: %s
		Age: %d
		Id: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
	params := &pb.GetUsersParams{}
	r, err := c.GetUsers(ctx, params)
	if err != nil {
		log.Fatalf("could not retrieve users: %v", err)
	}
	log.Print("\n USER LIST: \n")
	fmt.Printf("r.GetUsers(): %v\n", r.GetUsers())
}
