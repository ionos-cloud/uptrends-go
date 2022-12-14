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
		auth, "",
	)
	if err != nil {
		panic(err)
	}
}
