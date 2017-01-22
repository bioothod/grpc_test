package main

import (
	"fmt"
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"

	pb "github.com/bioothod/grpc_test/grtest"
)

type server struct {
}

func (ts *server) PingRequest(ctx context.Context, req *pb.Ping) (*pb.Pong, error) {
	reply := &pb.Pong {
		Pong:		"reply",
		Aux:		req.Aux,
	}
	fmt.Printf("ping: %s -> %s\n", req.Ping, reply.Pong)

	return reply, nil
}

func (ts *server) Stream(req *pb.Ping, stream pb.TestService_StreamServer) error {
	fmt.Printf("stream: %s:%s\n", req.Ping, req.Aux)

	var prev_time time.Time
	counter := 0
	for {
		reply := &pb.Pong {
			Pong:		"reply",
			Aux:		fmt.Sprintf("%s:%d", req.Aux, counter),
		}
		err := stream.Send(reply)
		if err != nil {
			err = fmt.Errorf("stream: send error: %v", err)
			fmt.Println(err.Error())
			return err
		}

		counter++

		t := time.Now()
		if t.After(prev_time.Add(10 * time.Second)) {
			fmt.Printf("%s: sending reply %s:%s\n", t.String(), reply.Pong, reply.Aux)
			prev_time = t
		}
	}

	return nil
}

func main() {
	listen := flag.String("listen", "localhost:12345", "Address to listen at")
	flag.Parse()

	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("Could not start listening at %s: %v\n", *listen, err)
	}

	ts := &server {
	}

	grpc.EnableTracing = true
	srv := grpc.NewServer()
	pb.RegisterTestServiceServer(srv, ts)
	reflection.Register(srv)
	log.Printf("Starting to listen at %s\n", *listen)
	srv.Serve(lis)
}
