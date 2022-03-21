/*
 * Connector Service Fleet Manager Private APIs
 *
 * Connector Service Fleet Manager apis that are used by internal services.
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ConnectorNamespaceTenantKind the model 'ConnectorNamespaceTenantKind'
type ConnectorNamespaceTenantKind string

// List of ConnectorNamespaceTenantKind
const (
	CONNECTORNAMESPACETENANTKIND_USER         ConnectorNamespaceTenantKind = "user"
	CONNECTORNAMESPACETENANTKIND_ORGANISATION ConnectorNamespaceTenantKind = "organisation"
)