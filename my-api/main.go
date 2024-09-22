package main

import (
	"context"

	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"scottschubert.dev/pubsub"
	"scottschubert.dev/uploader"
	"scottschubert.dev/websocketUtils"
)

var wsConn *websocket.Conn

func handleImageUpload(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
	}

	const topicName = "image-upload"
	const pubsubName = "scott-learning"
	const subscriptionName = "image-upload-res-subscription"

	req.ParseMultipartForm(0)
	file, header, err := req.FormFile("image")

	if err != nil {
		http.Error(w, "error parsing image", http.StatusBadRequest)
	}

	fmt.Println("file", file)
	fmt.Println("filename", &header.Filename)
	fmt.Println("err", err)

	uploadResChannel := make(chan uploader.UploadFileResponse)
	go func() {
		uploadResChannel <- uploader.UploadFile(w, "upload-test-bucket-scott", header.Filename)

	}()

	uploadRes := <-uploadResChannel
	fmt.Printf("uploadRes: %v\n", uploadRes)

	errMsg := uploadRes.Err
	if errMsg != nil {
		fmt.Println(errMsg)
		http.Error(w, errMsg.Error(), http.StatusBadRequest)
	}

	pubSubClient := pubsub.PubSubClient(pubsubName)

	topic := pubSubClient.Topic(topicName)

	// Create subscription
	createSubscriptionErr := pubsub.CreateSubscription(pubsubName, subscriptionName, topic)
	if createSubscriptionErr != nil {
		log.Fatalf("Create.Subscription: %v", createSubscriptionErr)
	}

	// Subscribe to messages
	pubsub.StreamingPullMsgs(w, pubsubName, subscriptionName)

	publishRes, err := pubsub.PublishToTopic(topic, pubsub.Message{Generation: uploadRes.Generation, Name: uploadRes.ObjectName})

	if err != nil {
		log.Fatalf("Error uploading message: %v", err)
	}

	msgID, err := publishRes.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msgID)
	str := fmt.Sprintf("%d", uploadRes.Generation)

	websocketUtils.Writer(wsConn, []byte(str))

	// // Copy the uploaded file to the server's file

	// _, err = io.Copy(outFile, file)
	// if err != nil {
	// 	http.Error(w, "Could not save file on server", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusNoContent)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	log.Println("Client connected")
	wsConn = conn

	websocketUtils.Reader(wsConn)
}

func main() {
	fmt.Println("The start of my API")

	// Create subscription
	// Subscribe to messages
	http.HandleFunc("/image", handleImageUpload)
	http.HandleFunc("/ws", wsEndpoint)

	resp, err := http.Get("https://gobyexample.com")

	if err != nil {
		fmt.Println("There was an error")
	}

	fmt.Println("Response status:", resp.Status)

	fmt.Println("Server starting on :8080...")

	fmt.Println(http.ListenAndServe(":8080", nil))
}
