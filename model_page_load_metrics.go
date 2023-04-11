/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type PageLoadMetrics struct {
	// Measures the sum total of all individual layout shift scores for every unexpected layout shift that occurs during the entire lifespan of the page. A score lower than 0.1 is considered a good score. Between 0.1 and 0.25 could use some improvement. 0.25 or higher is a poor score. https://web.dev/cls/?utm_source=devtools  In case the data was inconclusive null is returned.
	CumulativeLayoutShift float32 `json:"CumulativeLayoutShift,omitempty"`
	// Measures the time from when the page starts loading to when any of the page's content is rendered on screen. A score lower than 1000ms is considered a good score. https://web.dev/fcp/  In case the data was inconclusive null is returned.
	FirstContentfulPaint int32 `json:"FirstContentfulPaint,omitempty"`
	// The render timestamp of the largest image or text block visible within the viewport. A score lower than 2500ms is good. Between 2500ms and 4000ms needs improvement. 4000ms or higher is poor https://web.dev/lcp/  In case the data was inconclusive null is returned.
	LargestContentfulPaint int32 `json:"LargestContentfulPaint,omitempty"`
	// The Total Blocking Time (TBT) metric measures the total amount of time between First Contentful Paint (FCP) and Time to Interactive (TTI) where the main thread was blocked for long enough to prevent input responsiveness.  https://web.dev/tbt/  In case the data was inconclusive null is returned.
	TotalBlockingTime int32 `json:"TotalBlockingTime,omitempty"`
	// TTI measures how long it takes a page to become fully interactive.   https://web.dev/interactive/  In case the data was inconclusive null is returned.
	TimeToInteractive int32 `json:"TimeToInteractive,omitempty"`
}