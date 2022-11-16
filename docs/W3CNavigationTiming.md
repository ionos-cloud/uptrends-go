# W3CNavigationTiming

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestStart** | **int32** | Equal to the requestStart as defined by the W3C.  It is a timestamp indicating when the browser starts requesting the resource from the webserver after the DNS lookup and TCP connection. | [default to null]
**TimeToFirstByte** | **int32** | Equal to the difference between navigationStart and responseStart as defined by the W3C.  In short, it&#39;s the time between when the first request was sent from browser to server, and when the first bytes of the following response were received by the browser. | [default to null]
**DomInteractive** | **int32** | Equal to domInteractive as defined by W3C.  It is a timestamp, indicating the document readiness is set to &#39;interactive&#39;, to indicate that the browser has stopped parsing the page and the user can start interacting with it.  Resources such as scripts, images, stylesheets, or frames may still be loading. | [default to null]
**DomComplete** | **int32** | Equal to the domComplete as defined by W3C.  It is a timestamp, indicating when the main document has been parsed, the DOM has been fully loaded, and the page readiness is set to &#39;complete&#39;. | [default to null]
**LoadEvent** | **int32** | Equal to loadEventEnd as defined by W3C.  It is a timestamp, indicating when the load event of the current document has completed, including all dependent resources such as stylesheets and images. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


