# Integration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationGuid** | **string** | Guid of Integration in Alert Definition Escalation Level | [default to null]
**Name** | **string** | Name of Integration in Alert Definition Escalation Level | [optional] [default to null]
**Type_** | [***IntegrationType**](Integration_Type.md) |  | [default to null]
**ExtraEmailAddresses** | **string** | Extra emailadresses for this integration (if type &#x3D;&#x3D; email) | [optional] [default to null]
**StatusHubServiceList** | [**[]IntegrationServiceMap**](IntegrationServiceMap.md) | All statushubs for this integration (if type &#x3D;&#x3D; statushub) | [optional] [default to null]
**IntegrationServices** | **[]string** | All integrations services. | [optional] [default to null]
**Hash** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


