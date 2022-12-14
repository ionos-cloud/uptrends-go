/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// JwtAlgorithm :
type JwtAlgorithm string

// List of JwtAlgorithm
const (
	HS256 JwtAlgorithm = "HS256"
	HS384 JwtAlgorithm = "HS384"
	HS512 JwtAlgorithm = "HS512"
)
