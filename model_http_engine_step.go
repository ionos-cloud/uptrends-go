/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// HttpEngineStep describes the results of a step in a Multi-step Api monitor
type HttpEngineStep struct {
	// The name of the step
	StepName string `json:"StepName,omitempty"`
	// Url the step was executed on
	Url string `json:"Url,omitempty"`
	// The HTTP status code returned
	HttpStatusCode string `json:"HttpStatusCode,omitempty"`
	// Http method used in this step
	HttpMethod string `json:"HttpMethod,omitempty"`
	// Step description
	HttpStatusDescription string `json:"HttpStatusDescription,omitempty"`
	// Did the response complete?
	ResponseCompleted bool `json:"ResponseCompleted"`
	// Was this step executed?
	StepExecuted bool `json:"StepExecuted"`
	AssertionResultsInfo *HttpEngineStepAssertionResultsInfo `json:"AssertionResultsInfo,omitempty"`
	// Number of milliseconds it took for this step to succeed
	TotalTime int64 `json:"TotalTime"`
	// Response headers
	ResponseHeaders string `json:"ResponseHeaders,omitempty"`
	// Response body
	ResponseBody string `json:"ResponseBody,omitempty"`
	// Request headers send
	RequestHeaders string `json:"RequestHeaders,omitempty"`
	// Request body send
	RequestBody string `json:"RequestBody,omitempty"`
	// Resolved IP address
	ResolvedIpAddress string `json:"ResolvedIpAddress,omitempty"`
}
