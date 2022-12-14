/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// Monitor check attributes
type MonitorCheckAttributes struct {
	// Monitor identifier
	MonitorGuid string `json:"MonitorGuid"`
	// Date/time stamp of the check
	Timestamp interface{} `json:"Timestamp"`
	// The numeric Uptrends error code in case of an error result, or 0 in case of an OK result.
	ErrorCode int32 `json:"ErrorCode"`
	// The number of milliseconds needed to complete the monitor check.
	TotalTime float64 `json:"TotalTime"`
	// The number of milliseconds needed to perform the DNS query for this check, when appropriate.
	ResolveTime float64 `json:"ResolveTime"`
	// The number of milliseconds needed to establish a connection.
	ConnectionTime float64 `json:"ConnectionTime"`
	// The number of milliseconds needed to download the response data.
	DownloadTime float64 `json:"DownloadTime"`
	// The number of downloaded bytes for this check.
	TotalBytes int32 `json:"TotalBytes,omitempty"`
	// The IP address that was found for the specified domain name as part of this monitor check.
	ResolvedIpAddress string                            `json:"ResolvedIpAddress,omitempty"`
	ErrorLevel        *MonitorCheckAttributesErrorLevel `json:"ErrorLevel"`
	// A text value that describes the error that was found, or OK if no error was found.
	ErrorDescription string `json:"ErrorDescription,omitempty"`
	// Any additional error information, if available.
	ErrorMessage string `json:"ErrorMessage,omitempty"`
	// Did the check come from a staging monitor?
	StagingMode bool `json:"StagingMode"`
	// The Id of the Uptrends checkpoint server that performed this check.
	ServerId int32 `json:"ServerId,omitempty"`
	// The HTTP status code returned (if applicable)
	HttpStatusCode int32 `json:"HttpStatusCode,omitempty"`
	// This is a partial concurrent measurement, part of a concurrent check
	IsPartialCheck bool `json:"IsPartialCheck"`
	// Is this a master concurrent check record
	IsConcurrentCheck bool `json:"IsConcurrentCheck,omitempty"`
}
