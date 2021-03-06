/*
 * GeoServer Layers
 *
 * A layer is a published resource (feature type or coverage).
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// WMS attribution information to be drawn on each map
type LayerAttribution struct {
	// Data provider logo height
	LogoHeight int32 `json:"logoHeight,omitempty"`
	// Data provider logo width
	LogoWidth int32 `json:"logoWidth,omitempty"`
	// URL to data provider
	Href string `json:"href,omitempty"`
	// Human-readable text describing the data provider
	Title string `json:"title,omitempty"`
	// Format of data provider logo, example \"image/png\"
	LogoType string `json:"logoType,omitempty"`
	// Data provider logo
	LogoURL string `json:"logoURL,omitempty"`
}
