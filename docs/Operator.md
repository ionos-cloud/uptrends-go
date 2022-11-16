# Operator

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatorGuid** | **string** | The unique identifier of this operator | [optional] [default to null]
**Hash** | **string** | The hash of this operator. | [optional] [default to null]
**Password** | **string** | The password is a required field if AllowNativeLogin is set to True | [optional] [default to null]
**FullName** | **string** | The full name of this operator | [optional] [default to null]
**Email** | **string** | The email address of this operator. This also serves as the username | [optional] [default to null]
**MobilePhone** | **string** | The phone number of this operator to which SMS and phone alerts can be sent. Start with a plus (+) sign and your country code | [optional] [default to null]
**OutgoingPhoneNumberId** | **int32** | The id of the phone number that will be used to send phone alerts (See /OutgoingPhoneNumber API under Miscellaneous for available ids) | [optional] [default to null]
**OutgoingPhoneNumberIdSpecified** | **bool** |  | [optional] [default to null]
**IsAccountAdministrator** | **bool** | This indicates if the operator is the account administrator. | [optional] [default to null]
**BackupEmail** | **string** | The backup email address of this operator | [optional] [default to null]
**IsOnDuty** | **bool** | Indicates whether the operator is currently active | [optional] [default to null]
**CultureName** | **string** | If ommitted the operator will use the account culture. If set it will override the account default | [optional] [default to null]
**CultureNameSpecified** | **bool** |  | [optional] [default to null]
**TimeZoneId** | **int32** | The id of the timezone of this operator (See /Timezone API under Miscellaneous for available timezones) | [optional] [default to null]
**TimeZoneIdSpecified** | **bool** |  | [optional] [default to null]
**SmsProvider** | [***OperatorSmsProvider**](Operator_SmsProvider.md) |  | [optional] [default to null]
**UseNumericSender** | **bool** | Set to True to override the default behavior of sending SMS alerts with textual sender ID | [optional] [default to null]
**UseNumericSenderSpecified** | **bool** |  | [optional] [default to null]
**AllowNativeLogin** | **bool** | This can only be set to false if the account has SSO enabled. Ommitting or providing null will use the account default | [optional] [default to null]
**AllowNativeLoginSpecified** | **bool** |  | [optional] [default to null]
**AllowSingleSignon** | **bool** | This can only be set to true if the account has SSO enabled. Ommitting or providing null will use the account default | [optional] [default to null]
**AllowSingleSignonSpecified** | **bool** |  | [optional] [default to null]
**DefaultDashboard** | **string** | This is used to set the default dashboard for the operator.  Valid options are: UseAccountSpecifiedDashboard (This will use the dashboard specified for the account) Any built-in dashboard: e.g. AccountOverview Any custom dashboard to which the operator has access to, defined by the guid of this dashboard | [optional] [default to null]
**SetupMode** | [***OperatorSetupMode**](OperatorSetupMode.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


