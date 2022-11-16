/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends
// EngineHashAlgorithm : 
type EngineHashAlgorithm string

// List of EngineHashAlgorithm
const (
	MD5 EngineHashAlgorithm = "MD5"
	SHA1 EngineHashAlgorithm = "SHA1"
	SHA256 EngineHashAlgorithm = "SHA256"
	SHA512 EngineHashAlgorithm = "SHA512"
	HMACSHA1 EngineHashAlgorithm = "HMACSHA1"
	HMACSHA256 EngineHashAlgorithm = "HMACSHA256"
	HMACSHA512 EngineHashAlgorithm = "HMACSHA512"
)
