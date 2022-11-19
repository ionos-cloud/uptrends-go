package utils

import (
	sw "github.com/ionos-cloud/uptrends-go"
)

func PtrMonitor(m sw.MonitorType) *sw.MonitorType {
	return &m
}
