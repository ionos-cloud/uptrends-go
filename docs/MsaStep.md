# MsaStep

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [optional] [default to null]
**Method** | [***HttpMethods**](HttpMethods.md) |  | [default to null]
**Body** | **string** | The body that will be send in the request. Only used if BodyType equals Raw | [optional] [default to null]
**BodyType** | [***MsaStepBodyType**](MsaStep_BodyType.md) |  | [optional] [default to null]
**VaultFileId** | **string** | The guid of the vaultfile that will be send in the request. Only used if BodyType equals VaultFiles | [optional] [default to null]
**RequestHeaders** | [**[]ApiHttpHeaderValue**](ApiHttpHeaderValue.md) |  | [optional] [default to null]
**Variables** | [**[]ApiVariableDefinition**](ApiVariableDefinition.md) |  | [optional] [default to null]
**Assertions** | [**[]ApiAssertion**](ApiAssertion.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**UseFixedClientCertificate** | **bool** |  | [default to null]
**Authentication** | [***ApiAuthenticationInfo**](ApiAuthenticationInfo.md) |  | [optional] [default to null]
**IgnoreCertificateErrors** | **bool** |  | [default to null]
**ClientCertificateVaultItem** | **string** |  | [optional] [default to null]
**Delay** | **int32** |  | [default to null]
**StepType** | [***MsaStepType**](MsaStepType.md) |  | [default to null]
**RetryUntilSuccessful** | **bool** |  | [default to null]
**MaxAttempts** | **int32** |  | [default to null]
**RetryWaitMilliseconds** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


