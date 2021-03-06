/*
 * GeoServer Layers
 *
 * A layer is a published resource (feature type or coverage).
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LayerReference struct {
	// Name of layer
	Name string `json:"name,omitempty"`
	// URL to layer definition
	Link string `json:"link,omitempty"`
}
