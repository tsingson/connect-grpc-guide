package main

import (
	"context"
	"log"
	"net/http"

	v1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"

	"github.com/bufbuild/connect-go"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithGRPC(),
	)
	req := connect.NewRequest(&v1.GreetRequest{Name: "Jane"})
	req.Header().Set("Acme-Tenant-Id", "1234")

	res, err := client.Greet(
		context.Background(),
		req,
	)

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
