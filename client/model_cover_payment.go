/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CoverPayment
type CoverPayment struct {
	// SwiftFieldTag
	SwiftFieldTag string `json:"swiftFieldTag,omitempty"`
	// SwiftLineOne
	SwiftLineOne string `json:"swiftLineOne,omitempty"`
	// SwiftLineTwo
	SwiftLineTwo string `json:"swiftLineTwo,omitempty"`
	// SwiftLineThree
	SwiftLineThree string `json:"swiftLineThree,omitempty"`
	// SwiftLineFour
	SwiftLineFour string `json:"swiftLineFour,omitempty"`
	// SwiftLineFive
	SwiftLineFive string `json:"swiftLineFive,omitempty"`
	// SwiftLineSix
	SwiftLineSix string `json:"swiftLineSix,omitempty"`
}