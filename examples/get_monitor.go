package main

import (
	"context"
	"fmt"

	sw "github.com/ionos-cloud/uptrends-go"
)

func main() {
	auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
		UserName: "",
		Password: "",
	})

	client := sw.NewAPIClient(sw.NewConfiguration())

	mon, _, err := client.MonitorApi.MonitorGetMonitor(
		auth, "", &sw.MonitorApiMonitorGetMonitorOpts{},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(mon.SelectedCheckpoints)
}
