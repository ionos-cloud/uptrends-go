/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type SelectedPeriod struct {
	SelectedPeriodType *SelectedPeriodSelectedPeriodType `json:"SelectedPeriodType,omitempty"`
	// The start of a custom period (can't be used together with the SelectedPeriodType parameter)
	Start interface{} `json:"Start,omitempty"`
	// The end of a custom period
	End interface{} `json:"End,omitempty"`
	PresetPeriod *SelectedPeriodPresetPeriod `json:"PresetPeriod,omitempty"`
}