/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package uptrends

type CustomizationInfo struct {
	MainColor                    string         `json:"MainColor,omitempty"`
	BackgroundColor              string         `json:"BackgroundColor,omitempty"`
	TextColor                    string         `json:"TextColor,omitempty"`
	TitleText                    string         `json:"TitleText,omitempty"`
	FooterText                   string         `json:"FooterText,omitempty"`
	MonitorNameOverrideFieldName string         `json:"MonitorNameOverrideFieldName,omitempty"`
	SortColumnsNewToOld          bool           `json:"SortColumnsNewToOld"`
	SortRowsProperty             *SortOrderEnum `json:"SortRowsProperty"`
	CommentTitle                 string         `json:"CommentTitle,omitempty"`
	CommentText                  string         `json:"CommentText,omitempty"`
}
