/*
 * GeoServer Data Stores
 *
 * A data store contains vector format spatial data. It can be a file (such as a shapefile), a database (such as PostGIS), or a server (such as a remote Web Feature Service).
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// connection parameter key-value pair
type Entry struct {
	// Connection parameter value
	Value string `json:"value,omitempty"`
	// Connection parameter key
	Key string `json:"key,omitempty"`
}
