/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// Monitor group
type MonitorGroup struct {
	// The unique identifier of this monitor group
	MonitorGroupGuid string `json:"MonitorGroupGuid,omitempty"`
	// The descriptive name of this probe group
	Description string `json:"Description,omitempty"`
	// Indicates whether this is the default group for all probes
	IsAll bool `json:"IsAll"`
	// Indicates whether the monitor quota is unlimited Only administrators can change this
	IsQuotaUnlimited bool `json:"IsQuotaUnlimited,omitempty"`
	// The basic monitor quota for the monitor group Only administrators can change this
	BasicMonitorQuota int32 `json:"BasicMonitorQuota,omitempty"`
	// The browser monitor quota for the monitor group Only administrators can change this
	BrowserMonitorQuota int32 `json:"BrowserMonitorQuota,omitempty"`
	// The transaction monitor quota for the monitor group Only administrators can change this
	TransactionMonitorQuota int32 `json:"TransactionMonitorQuota,omitempty"`
	// The api monitor quota for the monitor group Only administrators can change this
	ApiMonitorQuota int32 `json:"ApiMonitorQuota,omitempty"`
	// The classic quota for the monitor group Only administrators can change this
	ClassicQuota int32 `json:"ClassicQuota,omitempty"`
}
