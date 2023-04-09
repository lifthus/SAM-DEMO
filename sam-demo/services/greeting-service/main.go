package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	httpadapter "github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	demoGreeting, ok1 := os.LookupEnv("DEMO_GREETING")
	demoUser, ok2 := os.LookupEnv("DEMO_USER")
	if !ok1 || !ok2 {
		log.Fatal("DEMO_GREETING and DEMO_USER must be set")
	}
	http.DefaultServeMux.HandleFunc("/greeting/who", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		io.WriteString(w, "I am "+demoUser)
	})
	http.DefaultServeMux.HandleFunc("/greeting/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		io.WriteString(w, demoGreeting)
	})
	lambda.Start(httpadapter.NewV2(http.DefaultServeMux).ProxyWithContext)
}
