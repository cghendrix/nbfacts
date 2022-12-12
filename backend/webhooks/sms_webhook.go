package webhookingester

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

func init() {
	functions.HTTP("ProcessMessage", ProcessMessage)
}

// StartServer for local testing
func StartServer() {
	http.HandleFunc("/", ProcessMessage)
	http.ListenAndServe(":80", nil)
}

func ProcessMessage(w http.ResponseWriter, r *http.Request) {
	// respond to signalwire
	fmt.Println("SENDING RESPONSE")
	sendHTTPResponse(w)

	// parse all message values from incoming query string
	fmt.Println("GETTING MESSAGE VALUES")
	messageValues := parseMessageValues(r)
	fmt.Println(messageValues)

	// send to pubsub topic
	fmt.Println("PUBLISHING MESSAGE")
	publishMessage(messageValues)
}

func sendHTTPResponse(w http.ResponseWriter) {
	// send blank response (no response text)
	fmt.Fprint(w,
		`<?xml version="1.0" encoding="UTF-8"?>
	<Response>
		<Message>Your fact has been added!</Message>
	</Response>`)
}

func parseMessageValues(r *http.Request) map[string]string {
	// get body string from request
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	bodyStr := string(rawBody)
	// add all message attributes from incoming query string
	messageValues := map[string]string{}
	re := regexp.MustCompile(`([^=|^&]+)\=([^&]+)`)
	for _, match := range re.FindAllStringSubmatch(bodyStr, -1) {
		paramName := match[1]
		paramValue := match[2]
		if decodedValue, err := url.QueryUnescape(paramValue); err != nil {
			log.Fatalf("error unescaping value: " + paramValue)
		} else {
			messageValues[paramName] = decodedValue
		}
	}
	return messageValues
}

func publishMessage(messageValues map[string]string) {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, os.Getenv("PROJECT_NAME"))
	if err != nil {
		log.Fatalf("error creating pub sub client: " + err.Error())
	}
	defer pubsubClient.Close()
	topic := pubsubClient.Topic(os.Getenv("TOPIC"))
	jsonStr, err := json.Marshal(messageValues)
	res := topic.Publish(ctx, &pubsub.Message{Data: jsonStr})
	id, err := res.Get(ctx)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("success: " + id)
	}
}
