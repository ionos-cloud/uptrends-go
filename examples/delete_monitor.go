package main

import (
	"context"

	sw "github.com/ionos-cloud/uptrends-go"
)

func main() {
	auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
		UserName: "",
		Password: "",
	})

	client := sw.NewAPIClient(sw.NewConfiguration())

	_, err := client.MonitorApi.MonitorDeleteMonitor(
		auth, "7ca2083c-4b9c-4e2f-a8cf-0b9a36268ca5",
	)
	if err != nil {
		panic(err)
	}
}
