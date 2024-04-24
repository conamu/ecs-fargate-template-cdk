package main

import (
	"github.com/conamu/cdk-fargate-template/internal/pkg/helloworld"
	"github.com/conamu/cdk-fargate-template/internal/pkg/logging"
	"log"
	"net/http"
)

func main() {

	logger := logging.NewLogger()

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", helloworld.IndexHandler)

	logger.Info("Serving on port 8080", "port", "8080")
	log.Fatal(server.ListenAndServe())

}
