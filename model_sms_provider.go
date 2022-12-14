/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// SmsProvider :
type SmsProvider string

// List of SmsProvider
const (
	USE_ACCOUNT_SETTING        SmsProvider = "UseAccountSetting"
	SMS_PROVIDER_EUROPE        SmsProvider = "SmsProviderEurope"
	SMS_PROVIDER_INTERNATIONAL SmsProvider = "SmsProviderInternational"
	SMS_PROVIDER_EUROPE2       SmsProvider = "SmsProviderEurope2"
	SMS_PROVIDER_USA           SmsProvider = "SmsProviderUSA"
)
