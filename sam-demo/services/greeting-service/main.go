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
	// get environment variables
	demoGreeting, ok1 := os.LookupEnv("DEMO_GREETING")
	demoUser, ok2 := os.LookupEnv("DEMO_USER")
	demoEnv, ok3 := os.LookupEnv("DEMO_ENV")
	if !ok1 || !ok2 || !ok3 {
		log.Fatal("DEMO_GREETING and DEMO_USER must be set")
	}

	// separate basePath by environment
	var basePath string
	if demoEnv == "production" {
		basePath = "/demo/greeting"
	} else if demoEnv == "development" {
		basePath = "/greeting"
	} else {
		log.Fatal("wrong environment")
	}

	http.DefaultServeMux.HandleFunc(basePath+"/who", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		io.WriteString(w, "I am "+demoUser)
	})
	http.DefaultServeMux.HandleFunc(basePath+"/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		io.WriteString(w, demoGreeting)
	})

	// run the lambda function with adapter
	lambda.Start(httpadapter.NewV2(http.DefaultServeMux).ProxyWithContext)
}
