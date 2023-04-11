/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// ApiComparisonType :
type ApiComparisonType string

// List of ApiComparisonType
const (
	EQUAL                 ApiComparisonType = "Equal"
	DOES_NOT_EQUAL        ApiComparisonType = "DoesNotEqual"
	CONTAINS_TEXT         ApiComparisonType = "ContainsText"
	DOES_NOT_CONTAIN_TEXT ApiComparisonType = "DoesNotContainText"
	SHOULD_BE_IGNORED     ApiComparisonType = "ShouldBeIgnored"
	LESS_THAN             ApiComparisonType = "LessThan"
	GREATER_THAN          ApiComparisonType = "GreaterThan"
	LESS_THAN_OR_EQUAL    ApiComparisonType = "LessThanOrEqual"
	GREATER_THAN_OR_EQUAL ApiComparisonType = "GreaterThanOrEqual"
	IS_NOT_EMPTY          ApiComparisonType = "IsNotEmpty"
	IS_EMPTY              ApiComparisonType = "IsEmpty"
	IS_NOT_NULL           ApiComparisonType = "IsNotNull"
	IS_NULL               ApiComparisonType = "IsNull"
	DOES_NOT_EXIST        ApiComparisonType = "DoesNotExist"
	EXISTS                ApiComparisonType = "Exists"
)
