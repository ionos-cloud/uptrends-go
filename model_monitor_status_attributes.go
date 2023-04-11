/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// Monitor Status attributes
type MonitorStatusAttributes struct {
	ErrorLevel *MonitorStatusAttributesErrorLevel `json:"ErrorLevel"`
	// Last checked timeStamp for this monitor
	LastCheck interface{} `json:"LastCheck,omitempty"`
	// Checkpoint id for the monitor status
	CheckpointId int32 `json:"CheckpointId,omitempty"`
	// Checkpoint name for the monitor status
	CheckpointName string `json:"CheckpointName,omitempty"`
	// Error description for the monitor status
	ErrorDescription string `json:"ErrorDescription,omitempty"`
	// Uptime percentage for the monitor status
	UptimePercentage float64 `json:"UptimePercentage"`
	// Error code for the monitor status
	ErrorCode int32 `json:"ErrorCode"`
	// Last monitor check id
	LastMonitorCheckId int64 `json:"LastMonitorCheckId,omitempty"`
	// Total time for the monitor status
	TotalTime int32 `json:"TotalTime,omitempty"`
}