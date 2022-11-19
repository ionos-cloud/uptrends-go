/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// Alert attributes
type AlertAttributes struct {
	AlertType *AlertAttributesAlertType `json:"AlertType"`
	// The monitor identifier.
	MonitorGuid string `json:"MonitorGuid"`
	// Date/time stamp of the alert.
	Timestamp interface{} `json:"Timestamp"`
	// Date/time stamp of the first monitor check.
	FirstError interface{} `json:"FirstError"`
	// The Id of the monitor check that triggered this alert.
	MonitorCheckId int64 `json:"MonitorCheckId"`
	// The Id of the first monitor check error that led to this alert.
	FirstErrorMonitorCheckId int64 `json:"FirstErrorMonitorCheckId"`
	// A text value that describes the error that was found, or OK if no error was found.
	ErrorDescription string `json:"ErrorDescription,omitempty"`
	// Any additional error information, if available.
	ErrorMessage string `json:"ErrorMessage,omitempty"`
	// The incident key of this alert.
	IncidentKey string `json:"IncidentKey,omitempty"`
}
