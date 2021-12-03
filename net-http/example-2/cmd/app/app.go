package main

import (
	"context"
	"log"

	"github.com/kakoitouser/coursera-course-golang/net-http/example-2/internal/httpserver"
	"github.com/kakoitouser/coursera-course-golang/net-http/example-2/internal/store/inmemory"
)

func main() {
	// ctx, cansel := context.WithTimeout(context.Background(), 1*time.Hour*24*31)
	// defer cansel()
	srv := httpserver.NewServer(context.Background(), ":8080", inmemory.NewDB())
	log.Fatal(srv.Run())
	srv.WaitForGracefulTermination()
}
