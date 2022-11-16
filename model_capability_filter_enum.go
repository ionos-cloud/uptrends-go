/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// CapabilityFilterEnum :
type CapabilityFilterEnum string

// List of CapabilityFilterEnum
const (
	IPV6              CapabilityFilterEnum = "IPv6"
	PRIMARY_SERVER    CapabilityFilterEnum = "PrimaryServer"
	HIGH_AVAILABILITY CapabilityFilterEnum = "HighAvailability"
	FIDDLER_PROXY     CapabilityFilterEnum = "FiddlerProxy"
	INTERNET_EXPLORER CapabilityFilterEnum = "InternetExplorer"
)
