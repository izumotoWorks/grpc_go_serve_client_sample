package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"bufio"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	hellopb "hello/src/grpc/gen"
)

var (
	scanner *bufio.Scanner
	client hellopb.GreeterClient
)

func main() {
	fmt.Println("start gRPC Client")

	scanner = bufio.NewScanner(os.Stdin)

	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Connection failed.")
		return
	}
	defer conn.Close()

	client = hellopb.NewGreeterClient(conn)

	for {
		fmt.Println("1. send Request")
		fmt.Println("2. exit")
		fmt.Print("please enter > ")
		
		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Hello()

		case "2":
			fmt.Println("bye.")
			goto M
		}
	}
M:
}

func Hello() {
	fmt.Print("please enter your name > ")
	scanner.Scan()
	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMessage())
	}
}