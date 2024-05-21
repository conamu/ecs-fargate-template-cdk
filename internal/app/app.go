package main

import (
	"cdk-fargate/internal/pkg/awsmeta"
	"cdk-fargate/internal/pkg/helloworld"
	"cdk-fargate/internal/pkg/logging"
	"log"
	"net/http"
)

func main() {

	logger := logging.NewLogger()
	awsmeta.Init(logger)

	server := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/", helloworld.IndexHandler)

	logger.Info("Serving on port 80", "port", "80")

	log.Fatal(server.ListenAndServe())

}
