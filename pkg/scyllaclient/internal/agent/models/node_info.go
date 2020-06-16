// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NodeInfo node info
//
// Information about Scylla node.
// swagger:model NodeInfo
type NodeInfo struct {

	// Address for Alternator API requests.
	AlternatorAddress string `json:"alternator_address,omitempty"`

	// Whether Alternator requires authentication.
	AlternatorEnforceAuthorization bool `json:"alternator_enforce_authorization,omitempty"`

	// Port for Alternator HTTPS API server.
	AlternatorHTTPSPort string `json:"alternator_https_port,omitempty"`

	// Port for Alternator API server.
	AlternatorPort string `json:"alternator_port,omitempty"`

	// Address for REST API requests.
	APIAddress string `json:"api_address,omitempty"`

	// Port for REST API server.
	APIPort string `json:"api_port,omitempty"`

	// Address that is broadcasted to tell other Scylla nodes to connect to. Related to listen_address.
	BroadcastAddress string `json:"broadcast_address,omitempty"`

	// Address that is broadcasted to tell the clients to connect to.
	BroadcastRPCAddress string `json:"broadcast_rpc_address,omitempty"`

	// Whether client encryption is enabled.
	ClientEncryptionEnabled bool `json:"client_encryption_enabled,omitempty"`

	// Whether client authorization is required.
	ClientEncryptionRequireAuth bool `json:"client_encryption_require_auth,omitempty"`

	// Whether CQL requires password authentication.
	CqlPasswordProtected bool `json:"cql_password_protected,omitempty"`

	// Address Scylla listens for connections from other nodes.
	ListenAddress string `json:"listen_address,omitempty"`

	// Port for the CQL native transport to listen for clients on.
	NativeTransportPort string `json:"native_transport_port,omitempty"`

	// Address for Prometheus queries.
	PrometheusAddress string `json:"prometheus_address,omitempty"`

	// Port for Prometheus server.
	PrometheusPort string `json:"prometheus_port,omitempty"`

	// Address on which Scylla is going to expect Thrift and CQL clients connections.
	RPCAddress string `json:"rpc_address,omitempty"`

	// Port for Thrift to listen for clients on.
	RPCPort string `json:"rpc_port,omitempty"`
}

// Validate validates this node info
func (m *NodeInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeInfo) UnmarshalBinary(b []byte) error {
	var res NodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
