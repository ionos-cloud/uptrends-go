/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type Checkpoint2 struct {
	Attributes *StatisticsAttributes `json:"Attributes,omitempty"`
	// Identifier
	Id int32 `json:"Id"`
	// Object type
	Type_ string `json:"Type,omitempty"`
	// Relationships of the object
	Relationships []RelationObject `json:"Relationships,omitempty"`
	// Links related to the object
	Links map[string]string `json:"Links,omitempty"`
}
