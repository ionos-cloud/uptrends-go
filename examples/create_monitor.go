package main

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	sw "github.com/ionos-cloud/uptrends-go"
	"github.com/ionos-cloud/uptrends-go/utils"
)

func main() {
	auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
		UserName: "bfd3c47585824e9badd88a6b56235de1",
		Password: "TjqkuZqcwAQDoeiqZQrr/YbjRbgcF6aj",
	})

	client := sw.NewAPIClient(sw.NewConfiguration())

	new := sw.Monitor{
		Name:          "ionos.com - Uptime",
		Url:           "https://ionos.com",
		MonitorType:   utils.PtrMonitor(sw.HTTP),
		CheckInterval: 5,
		SelectedCheckpoints: &sw.SelectedCheckpoints{
			Checkpoints:      []int32{},
			Regions:          []int32{},
			ExcludeLocations: []int32{},
		},
	}

	mon, _, err := client.MonitorApi.MonitorPostMonitor(
		auth, &sw.MonitorApiMonitorPostMonitorOpts{
			Monitor: optional.NewInterface(new),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(mon.MonitorGuid)
}
