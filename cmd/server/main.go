package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	v1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[v1.GreetRequest],
) (*connect.Response[v1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	fmt.Println(req.Header().Get("Acme-Tenant-Id"))

	res := connect.NewResponse(&v1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")

	return res, nil
}

// func (s *GreetServer) GreetStream(ctx context.Context, stream *connect.BidiStream[v1.GreetRequest, v1.GreetResponse]) error {
//
//		for {
//			if err := ctx.Err(); err != nil {
//				return err
//			}
//			request, err := stream.Receive()
//			if err != nil && errors.Is(err, io.EOF) {
//				return nil
//			} else if err != nil {
//				return fmt.Errorf("receive request: %w", err)
//			}
//			reply := fmt.Sprintf("Hello, %s!", request.Name)
//			if err := stream.Send(&v1.GreetResponse{
//				Greeting: reply,
//			}); err != nil {
//				return fmt.Errorf("send response: %w", err)
//			}
//
//		}
//		return nil
//	}
func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
