package main

import (
	"context"
	"fmt"

	sw "github.com/ionos-cloud/uptrends-go"
)

func main() {
	auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
		UserName: "bfd3c47585824e9badd88a6b56235de1",
		Password: "TjqkuZqcwAQDoeiqZQrr/YbjRbgcF6aj",
	})

	client := sw.NewAPIClient(sw.NewConfiguration())

	mon, _, err := client.MonitorApi.MonitorGetMonitor(
		auth, "bc2ecb25-d69d-417a-b2f1-445849fbb2b7", &sw.MonitorApiMonitorGetMonitorOpts{},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(mon.SelectedCheckpoints)
}
