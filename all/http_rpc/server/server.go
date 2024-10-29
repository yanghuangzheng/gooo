package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServer struct{}

func (s *HelloServer) Hello(req string, reply *string) error {
	*reply = "hello, " + req
	return nil
}

func main() {
	err := rpc.RegisterName("HelloService", new(HelloServer))
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe("1234", nil)
}
