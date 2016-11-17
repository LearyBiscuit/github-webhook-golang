package main

import (
	"log"
	"net/http"

	"github.com/LearyBiscuit/github-webhook-golang"
)

func main() {
	// To verify webhook's payload, set secret by SetSecret().
	webhook.SetSecret([]byte("abcdefgh"))

	http.HandleFunc("/", webhook.HandleHook(func(ev *webhook.Event) {
		// ev.Header.EventType contains webhook event type
		switch ev.Header.EventType {
			case "push": 
				push := ev.PushEvent()
				if push == nil {
					return
				}
				log.Printf("push: verified=%v %#v", ev.Verified, push)
			case "pull_request":
				pull_request := ev.PullRequest()
				if pull_request == nil {
					return
				}
				log.Printf("pull_request: verified=%v %#v", ev.Verified, pull_request)

		}
	}))
	
	// Add a HandlerFunc to process webhook.
	// http.HandleFunc("/", webhook.HandlePush(func(ev *webhook.Event) {
	// 	push := ev.PushEvent()
	// 	if push == nil {
	// 		return
	// 	}
	// 	log.Printf("push: verified=%v %#v", ev.Verified, push)
	// }))

	// Start web server.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
