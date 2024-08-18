package main

import (
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/rs/cors"
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

	// Configure CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Set to false if credentials should not be allowed
	}).Handler

	server.SetHandler(corsHandler(server.GetHandler()))

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
