// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NewWorkflowDefinitionRequest new workflow definition request
// swagger:model NewWorkflowDefinitionRequest

type NewWorkflowDefinitionRequest struct {

	// manager
	Manager Manager `json:"manager,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// state machine
	StateMachine *SLStateMachine `json:"stateMachine,omitempty"`
}

/* polymorph NewWorkflowDefinitionRequest manager false */

/* polymorph NewWorkflowDefinitionRequest name false */

/* polymorph NewWorkflowDefinitionRequest stateMachine false */

// Validate validates this new workflow definition request
func (m *NewWorkflowDefinitionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManager(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStateMachine(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewWorkflowDefinitionRequest) validateManager(formats strfmt.Registry) error {

	if swag.IsZero(m.Manager) { // not required
		return nil
	}

	if err := m.Manager.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("manager")
		}
		return err
	}

	return nil
}

func (m *NewWorkflowDefinitionRequest) validateStateMachine(formats strfmt.Registry) error {

	if swag.IsZero(m.StateMachine) { // not required
		return nil
	}

	if m.StateMachine != nil {

		if err := m.StateMachine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stateMachine")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewWorkflowDefinitionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewWorkflowDefinitionRequest) UnmarshalBinary(b []byte) error {
	var res NewWorkflowDefinitionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
