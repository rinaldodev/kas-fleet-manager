/*
 * Connector Service Fleet Manager Private APIs
 *
 * Connector Service Fleet Manager apis that are used by internal services.
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ConnectorNamespaceRequestMeta struct for ConnectorNamespaceRequestMeta
type ConnectorNamespaceRequestMeta struct {
	Name        string                                     `json:"name,omitempty"`
	Annotations []ConnectorNamespaceRequestMetaAnnotations `json:"annotations,omitempty"`
}