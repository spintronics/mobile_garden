/*
 * mobile_garden_api_v1
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SoilReading struct {
	Id string `json:"id,omitempty"`

	Moisture float64 `json:"moisture,omitempty"`

	Date string `json:"date,omitempty"`
}
