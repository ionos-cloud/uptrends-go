/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type ThrottlingOptions struct {
	ThrottlingType      *ThrottlingType  `json:"ThrottlingType,omitempty"`
	ThrottlingValue     *ThrottlingValue `json:"ThrottlingValue,omitempty"`
	ThrottlingSpeedUp   int32            `json:"ThrottlingSpeedUp,omitempty"`
	ThrottlingSpeedDown int32            `json:"ThrottlingSpeedDown,omitempty"`
	ThrottlingLatency   int32            `json:"ThrottlingLatency,omitempty"`
}
