package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func doHttpRunGet() {
	//transport := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	//}
	//resource := "http://localhost:8090/api/Runner/RunGet/Golang"
	resource := "https://localhost:8090/api/Runner/RunGet/Golang"
	request, err := http.NewRequest("GET", resource, nil)
	if err != nil {
		log.Fatalf("NewRequest: %v", err)
	}
	// HTTP headers that start with 'Grpc-Metadata-' are mapped to gRPC metadata after removing prefix 'Grpc-Metadata-'.
	request.Header.Add("Grpc-Metadata-Key", "value")
	client := &http.Client{
		//Transport: transport,
	}
	resp, err := client.Do(request)
	//resp, err := client.Get(resource)
	if err != nil {
		log.Fatalf("Do fail: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll fail: %v", err)
	}
	fmt.Println(string(body))
}

func doHttpRunPost() {
	//transport := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	//}
	//resource := "http://localhost:8090/api/Runner/RunPost"
	resource := "https://localhost:8090/api/Runner/RunPost"
	client := &http.Client{
		//Transport: transport,
	}
	resp, err := client.Post(
		resource,
		"application/json",
		strings.NewReader("{\"name\": \"Golang\"}"),
	)
	if err != nil {
		log.Fatalf("Do fail: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll fail: %v", err)
	}
	fmt.Println(string(body))
}
