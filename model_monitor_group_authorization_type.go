/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends
// MonitorGroupAuthorizationType : 
type MonitorGroupAuthorizationType string

// List of MonitorGroupAuthorizationType
const (
	VIEW_MONITOR_DATA_IN_GROUP MonitorGroupAuthorizationType = "ViewMonitorDataInGroup"
	VIEW_MONITORS_IN_GROUP MonitorGroupAuthorizationType = "ViewMonitorsInGroup"
	EDIT_MONITORS_IN_GROUP MonitorGroupAuthorizationType = "EditMonitorsInGroup"
	CREATE_AND_DELETE_MONITORS_IN_GROUP MonitorGroupAuthorizationType = "CreateAndDeleteMonitorsInGroup"
)
