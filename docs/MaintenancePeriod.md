# MaintenancePeriod

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique ID of this maintenance period | [default to null]
**ScheduleMode** | [***MaintenancePeriodScheduleMode**](MaintenancePeriod_ScheduleMode.md) |  | [default to null]
**StartDateTime** | **interface{}** | The start date/time for this schedule (for one-time schedules only) | [optional] [default to null]
**EndDateTime** | **interface{}** | The end date/time for this maintenance period (for one-time maintenance periods only) | [optional] [default to null]
**WeekDay** | [***MaintenancePeriodWeekDay**](MaintenancePeriod_WeekDay.md) |  | [optional] [default to null]
**MonthDay** | **int32** | the month day for this maintenance period (for montly maintenance periods only) | [optional] [default to null]
**StartTime** | **string** | The start time of this maintenance period | [optional] [default to null]
**EndTime** | **string** | The end time of this maintenance period | [optional] [default to null]
**MaintenanceType** | [***MaintenancePeriodMaintenanceType**](MaintenancePeriod_MaintenanceType.md) |  | [default to null]
**Description** | **string** | The description for this maintenance period | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


