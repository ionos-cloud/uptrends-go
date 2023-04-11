/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// Alert definition authorization
type AlertDefinitionAuthorization struct {
	// The unique ID of this authorization
	AuthorizationId string `json:"AuthorizationId,omitempty"`
	AuthorizationType *PspAuthorizationAuthorizationType `json:"AuthorizationType"`
	// The GUID of the operator (NULL if this authorization is for an operator group)
	OperatorGuid string `json:"OperatorGuid,omitempty"`
	// The GUID of the operator group (NULL if this authorization is for an operator)
	OperatorGroupGuid string `json:"OperatorGroupGuid,omitempty"`
}