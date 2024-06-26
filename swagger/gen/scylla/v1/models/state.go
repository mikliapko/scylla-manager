// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// State State
//
// # Current session state
//
// swagger:model State
type State string

const (

	// StateINITIALIZED captures enum value "INITIALIZED"
	StateINITIALIZED State = "INITIALIZED"

	// StatePREPARING captures enum value "PREPARING"
	StatePREPARING State = "PREPARING"

	// StateSTREAMING captures enum value "STREAMING"
	StateSTREAMING State = "STREAMING"

	// StateWAITCOMPLETE captures enum value "WAIT_COMPLETE"
	StateWAITCOMPLETE State = "WAIT_COMPLETE"

	// StateCOMPLETE captures enum value "COMPLETE"
	StateCOMPLETE State = "COMPLETE"

	// StateFAILED captures enum value "FAILED"
	StateFAILED State = "FAILED"
)

// for schema
var stateEnum []interface{}

func init() {
	var res []State
	if err := json.Unmarshal([]byte(`["INITIALIZED","PREPARING","STREAMING","WAIT_COMPLETE","COMPLETE","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stateEnum = append(stateEnum, v)
	}
}

func (m State) validateStateEnum(path, location string, value State) error {
	if err := validate.EnumCase(path, location, value, stateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this state
func (m State) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
