/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends
// DnsQuery : 
type DnsQuery string

// List of DnsQuery
const (
	A_RECORD DnsQuery = "ARecord"
	CNAME_RECORD DnsQuery = "CnameRecord"
	MX_RECORD DnsQuery = "MxRecord"
	NS_RECORD DnsQuery = "NsRecord"
	TXT_RECORD DnsQuery = "TxtRecord"
	SOA_RECORD DnsQuery = "SoaRecord"
	ROOT_SERVER DnsQuery = "RootServer"
	AAAA_RECORD DnsQuery = "AaaaRecord"
	SRV_RECORD DnsQuery = "SrvRecord"
)