// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestoreTarget restore target
//
// swagger:model RestoreTarget
type RestoreTarget struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// dc
	Dc []string `json:"dc"`

	// host
	Host string `json:"host,omitempty"`

	// location
	Location []string `json:"location"`
}

// Validate validates this restore target
func (m *RestoreTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreTarget) UnmarshalBinary(b []byte) error {
	var res RestoreTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
