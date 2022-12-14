/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

// SubStepType :
type SubStepType string

// List of SubStepType
const (
	NAVIGATE         SubStepType = "Navigate"
	CLICK            SubStepType = "Click"
	SET              SubStepType = "Set"
	TEST             SubStepType = "Test"
	SCRIPT           SubStepType = "Script"
	HOVER            SubStepType = "Hover"
	SCREENSHOT       SubStepType = "Screenshot"
	SCROLL_TO        SubStepType = "ScrollTo"
	WAIT_FOR_ELEMENT SubStepType = "WaitForElement"
	SWITCH_TO_FRAME  SubStepType = "SwitchToFrame"
	SWITCH_TO_TAB    SubStepType = "SwitchToTab"
)
