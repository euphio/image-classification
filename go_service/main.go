package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	pb "go_service/proto"

	"google.golang.org/grpc"
)

const pythonServiceAddress = "python_service:50051"

func labelImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	conn, err := grpc.Dial(pythonServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to python service: %v", err)
	}
	defer conn.Close()

	client := pb.NewImageLabelServiceClient(conn)
	resp, err := client.GetLabels(context.Background(), &pb.ImageRequest{ImageData: imageBytes})
	if err != nil {
		log.Printf("gRPC GetLabels failed: %v\n", err)
		http.Error(w, "Failed to get labels", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Labels: %v", resp.Labels)
}

func main() {
	http.HandleFunc("/label-image", labelImageHandler)
	log.Println("Go service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
