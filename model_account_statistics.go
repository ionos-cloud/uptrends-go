/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type AccountStatistics struct {
	AccountId string `json:"AccountId,omitempty"`
	ExpirationDate interface{} `json:"ExpirationDate,omitempty"`
	MonitorQuota *MonitorQuota `json:"MonitorQuota,omitempty"`
	OperatorQuota *OperatorQuota `json:"OperatorQuota,omitempty"`
	RemainingMessageCredits int32 `json:"RemainingMessageCredits,omitempty"`
}