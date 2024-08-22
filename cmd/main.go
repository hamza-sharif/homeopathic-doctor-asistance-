package main

import (
	"github.com/rs/cors"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/spf13/viper"

	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/config"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	runT, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := handlers.NewHandler(runT, swaggerSpec)

	server := restapi.NewServer(api)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	server.Host = viper.GetString(config.ServerHost)

	if err != nil {
		panic(err)
	}
	// Apply CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins, or specify your domains
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	server.ConfigureAPI()

	myHandler := server.GetHandler()
	server.SetHandler(corsMiddleware.Handler(myHandler))

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
