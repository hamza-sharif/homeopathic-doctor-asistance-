package main

import (
	"flag"
	"log"
	"net/http"

	"example/gen/restapi"
	"example/gen/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewExampleServerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Command line options
	server.Host = "localhost"
	server.Port = 8080

	// Handler function for the /hello endpoint
	api.GetHelloHandler = operations.GetHelloHandlerFunc(func(params operations.GetHelloParams) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, pr strfmt.Producer) {
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("Hello, world!"))
		})
	})

	// Parse flags
	flag.Parse()
	server.ConfigureFlags()
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
