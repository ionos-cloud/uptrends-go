/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends
// ScheduleType : 
type ScheduleType string

// List of ScheduleType
const (
	DAILY ScheduleType = "Daily"
	WEEKLY ScheduleType = "Weekly"
	MONTHLY ScheduleType = "Monthly"
	QUARTERLY ScheduleType = "Quarterly"
	YEARLY ScheduleType = "Yearly"
)
