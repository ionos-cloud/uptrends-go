# PageLoadMetrics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CumulativeLayoutShift** | **float32** | Measures the sum total of all individual layout shift scores for every unexpected layout shift that occurs during the entire lifespan of the page. A score lower than 0.1 is considered a good score. Between 0.1 and 0.25 could use some improvement. 0.25 or higher is a poor score. https://web.dev/cls/?utm_source&#x3D;devtools  In case the data was inconclusive null is returned. | [optional] [default to null]
**FirstContentfulPaint** | **int32** | Measures the time from when the page starts loading to when any of the page&#39;s content is rendered on screen. A score lower than 1000ms is considered a good score. https://web.dev/fcp/  In case the data was inconclusive null is returned. | [optional] [default to null]
**LargestContentfulPaint** | **int32** | The render timestamp of the largest image or text block visible within the viewport. A score lower than 2500ms is good. Between 2500ms and 4000ms needs improvement. 4000ms or higher is poor https://web.dev/lcp/  In case the data was inconclusive null is returned. | [optional] [default to null]
**TotalBlockingTime** | **int32** | The Total Blocking Time (TBT) metric measures the total amount of time between First Contentful Paint (FCP) and Time to Interactive (TTI) where the main thread was blocked for long enough to prevent input responsiveness.  https://web.dev/tbt/  In case the data was inconclusive null is returned. | [optional] [default to null]
**TimeToInteractive** | **int32** | TTI measures how long it takes a page to become fully interactive.   https://web.dev/interactive/  In case the data was inconclusive null is returned. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


