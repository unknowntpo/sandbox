package main

import (
	pb "../healthcheck"
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
	c := pb.NewHealthcheckerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Healthcheck(ctx, &pb.HealthcheckRequest{Message: "Are you alive?"})
	if err != nil {

		log.Fatalf("could not do healthcheck: %v", err)
	}
	log.Printf("Healtheck: %s", r.GetMessage())
}
