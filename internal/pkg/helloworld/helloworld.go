package helloworld

import (
	"cdk-fargate/internal/pkg/awsmeta"
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	meta := awsmeta.Get()

	w.Write([]byte(fmt.Sprintf("Hello World from %s!", meta.DockerID)))
}
