/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type Monitor struct {
	// The unique key of this monitor.
	MonitorGuid string `json:"MonitorGuid,omitempty"`
	// The name of this monitor.
	Name string `json:"Name,omitempty"`
	MonitorType *MonitorType `json:"MonitorType,omitempty"`
	// Indicates whether this monitor should generate alerts.
	GenerateAlert bool `json:"GenerateAlert,omitempty"`
	// Indicates whether this monitor is actively running according to the monitoring interval.
	IsActive bool `json:"IsActive,omitempty"`
	// Indicates whether this monitor is locked.
	IsLocked bool `json:"IsLocked,omitempty"`
	CustomFields []CustomField `json:"CustomFields,omitempty"`
	SelectedCheckpoints *SelectedCheckpoints `json:"SelectedCheckpoints,omitempty"`
	UsePrimaryCheckpointsOnly bool `json:"UsePrimaryCheckpointsOnly,omitempty"`
	// Indicates the interval in seconds
	CheckInterval int32 `json:"CheckInterval,omitempty"`
	MonitorMode *MonitorMonitorMode `json:"MonitorMode,omitempty"`
	Notes string `json:"Notes,omitempty"`
	// Hash corresponding with this monitor.
	Hash string `json:"Hash,omitempty"`
	RequestHeaders []RequestHeader `json:"RequestHeaders,omitempty"`
	UserAgent string `json:"UserAgent,omitempty"`
	LoadTimeLimit1 int32 `json:"LoadTimeLimit1,omitempty"`
	AlertOnLoadTimeLimit1 bool `json:"AlertOnLoadTimeLimit1,omitempty"`
	LoadTimeLimit2 int32 `json:"LoadTimeLimit2,omitempty"`
	AlertOnLoadTimeLimit2 bool `json:"AlertOnLoadTimeLimit2,omitempty"`
	Username string `json:"Username,omitempty"`
	AuthenticationType *ApiHttpAuthenticationType `json:"AuthenticationType,omitempty"`
	CheckCertificateErrors bool `json:"CheckCertificateErrors,omitempty"`
	IpVersion *IpVersion `json:"IpVersion,omitempty"`
	AlertOnMinimumBytes bool `json:"AlertOnMinimumBytes,omitempty"`
	MinimumBytes int32 `json:"MinimumBytes,omitempty"`
	HttpMethod *HttpMethod `json:"HttpMethod,omitempty"`
	TlsVersion *TlsVersion `json:"TlsVersion,omitempty"`
	RequestBody string `json:"RequestBody,omitempty"`
	MatchPatterns []PatternMatch `json:"MatchPatterns,omitempty"`
	Url string `json:"Url,omitempty"`
	Credits int32 `json:"Credits,omitempty"`
	PredefinedVariables []PredefinedVariable `json:"PredefinedVariables,omitempty"`
	MsaSteps []MsaStep `json:"MsaSteps,omitempty"`
	UserDefinedFunctions []UserDefinedFunction `json:"UserDefinedFunctions,omitempty"`
	CustomMetrics []CustomMetric `json:"CustomMetrics,omitempty"`
	SelfServiceTransactionScript string `json:"SelfServiceTransactionScript,omitempty"`
	MultiStepApiTransactionScript string `json:"MultiStepApiTransactionScript,omitempty"`
	BlockGoogleAnalytics bool `json:"BlockGoogleAnalytics,omitempty"`
	BlockUptrendsRum bool `json:"BlockUptrendsRum,omitempty"`
	BlockUrls []string `json:"BlockUrls,omitempty"`
	Password string `json:"Password,omitempty"`
	NameForPhoneAlerts string `json:"NameForPhoneAlerts,omitempty"`
	ThrottlingOptions *ThrottlingOptions `json:"ThrottlingOptions,omitempty"`
	DnsBypasses []DnsBypass `json:"DnsBypasses,omitempty"`
	TransactionStepDefinition *MonitorTransactionStepDefinition `json:"TransactionStepDefinition,omitempty"`
	CertificateName string `json:"CertificateName,omitempty"`
	CertificateOrganization string `json:"CertificateOrganization,omitempty"`
	CertificateOrganizationalUnit string `json:"CertificateOrganizationalUnit,omitempty"`
	CertificateSerialNumber string `json:"CertificateSerialNumber,omitempty"`
	CertificateFingerprint string `json:"CertificateFingerprint,omitempty"`
	CertificateIssuerName string `json:"CertificateIssuerName,omitempty"`
	CertificateIssuerCompanyName string `json:"CertificateIssuerCompanyName,omitempty"`
	CertificateIssuerOrganizationalUnit string `json:"CertificateIssuerOrganizationalUnit,omitempty"`
	CertificateExpirationWarningDays int32 `json:"CertificateExpirationWarningDays,omitempty"`
	AlertOnMaximumBytes bool `json:"AlertOnMaximumBytes,omitempty"`
	MaximumBytes int32 `json:"MaximumBytes,omitempty"`
	AlertOnMaximumSize bool `json:"AlertOnMaximumSize,omitempty"`
	ElementMaximumSize int32 `json:"ElementMaximumSize,omitempty"`
	IgnoreExternalElements bool `json:"IgnoreExternalElements,omitempty"`
	AlertOnPercentageFail bool `json:"AlertOnPercentageFail,omitempty"`
	FailedObjectPercentage int32 `json:"FailedObjectPercentage,omitempty"`
	DomainGroupGuid string `json:"DomainGroupGuid,omitempty"`
	DomainGroupGuidSpecified bool `json:"DomainGroupGuidSpecified,omitempty"`
	DnsServer string `json:"DnsServer,omitempty"`
	DnsQuery *DnsQuery `json:"DnsQuery,omitempty"`
	DnsExpectedResult string `json:"DnsExpectedResult,omitempty"`
	DnsTestValue string `json:"DnsTestValue,omitempty"`
	Port int32 `json:"Port,omitempty"`
	DatabaseName string `json:"DatabaseName,omitempty"`
	NetworkAddress string `json:"NetworkAddress,omitempty"`
	ImapSecureConnection bool `json:"ImapSecureConnection,omitempty"`
	SftpAction *SftpAction `json:"SftpAction,omitempty"`
	SftpActionPath string `json:"SftpActionPath,omitempty"`
	ExpectedHttpStatusCode int32 `json:"ExpectedHttpStatusCode,omitempty"`
	ExpectedHttpStatusCodeSpecified bool `json:"ExpectedHttpStatusCodeSpecified,omitempty"`
	BrowserType *BrowserType `json:"BrowserType,omitempty"`
	BrowserWindowDimensions *BrowserWindowDimensions `json:"BrowserWindowDimensions,omitempty"`
	UseConcurrentMonitoring bool `json:"UseConcurrentMonitoring,omitempty"`
	ConcurrentUnconfirmedErrorThreshold int32 `json:"ConcurrentUnconfirmedErrorThreshold,omitempty"`
	ConcurrentConfirmedErrorThreshold int32 `json:"ConcurrentConfirmedErrorThreshold,omitempty"`
	ErrorConditions []ErrorCondition `json:"ErrorConditions,omitempty"`
	CreatedDate interface{} `json:"CreatedDate,omitempty"`
}