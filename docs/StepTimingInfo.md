# StepTimingInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | [optional] [default to null]
**StartUtc** | **interface{}** |  | [default to null]
**EndUtc** | **interface{}** |  | [default to null]
**ElapsedMilliseconds** | **int64** |  | [default to null]
**DelayMilliseconds** | **int64** |  | [default to null]
**SubTimingInfos** | [**[]StepTimingInfo**](StepTimingInfo.md) |  | [optional] [default to null]
**IsValid** | **bool** | If true, this TimingInfo should be counted as part of the sum of its siblings. If false, the TimingInfo should be discarded (e.g. for PreDelays we don&#39;t want to count). | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


