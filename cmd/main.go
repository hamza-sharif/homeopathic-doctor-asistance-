package main

import (
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/spf13/viper"
	"golang.org/x/net/html"

	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/config"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/handlers"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
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
	//addMadd(runT)

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

func addMadd(rt *runtime.Runtime) {
	var alphabets []string

	// Loop through ASCII values for A-Z
	for i := 'a'; i <= 'z'; i++ {
		alphabets = append(alphabets, string(i))
	}
	for _, alphabet := range alphabets {
		resp, err := http.Get("http://homeoint.org/books/boericmm/" + alphabet + ".htm")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		// Parse the HTML response
		doc, err := html.Parse(strings.NewReader(string(body)))
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			return
		}

		// Extract text content and print it
		text := extractText(doc)

		re := regexp.MustCompile(`------>\s*([^\n]+)`)

		// Find all matches
		matches := re.FindAllStringSubmatch(text, -1)

		// Print the matched values after removing leading/trailing whitespace
		for _, match := range matches {
			if len(match) > 1 {
				name := strings.TrimSpace(match[1])
				err := rt.Svc.AddMedicine(&models.Medicine{Name: name})
				if err != nil {
					fmt.Println("Error adding medicine ===> :", name)
					println(err.Error())
				}
			}
		}

	}
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return text
}
