package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SLChoice s l choice
// swagger:model SLChoice
type SLChoice struct {

	// and
	And []*SLChoice `json:"And"`

	// boolean equals
	BooleanEquals *bool `json:"BooleanEquals,omitempty"`

	// next
	Next string `json:"Next,omitempty"`

	// not
	Not *SLChoice `json:"Not,omitempty"`

	// numeric equals
	NumericEquals *int64 `json:"NumericEquals,omitempty"`

	// numeric greater than
	NumericGreaterThan *float64 `json:"NumericGreaterThan,omitempty"`

	// numeric greater than equals
	NumericGreaterThanEquals *int64 `json:"NumericGreaterThanEquals,omitempty"`

	// numeric less than
	NumericLessThan *float64 `json:"NumericLessThan,omitempty"`

	// numeric less than equals
	NumericLessThanEquals *int64 `json:"NumericLessThanEquals,omitempty"`

	// or
	Or []*SLChoice `json:"Or"`

	// string equals
	StringEquals *string `json:"StringEquals,omitempty"`

	// string greater than
	StringGreaterThan *string `json:"StringGreaterThan,omitempty"`

	// string greater than equals
	StringGreaterThanEquals *string `json:"StringGreaterThanEquals,omitempty"`

	// string less than
	StringLessThan *string `json:"StringLessThan,omitempty"`

	// string less than equals
	StringLessThanEquals *string `json:"StringLessThanEquals,omitempty"`

	// timestamp equals
	TimestampEquals *strfmt.DateTime `json:"TimestampEquals,omitempty"`

	// timestamp greater than
	TimestampGreaterThan *strfmt.DateTime `json:"TimestampGreaterThan,omitempty"`

	// timestamp greater than equals
	TimestampGreaterThanEquals *strfmt.DateTime `json:"TimestampGreaterThanEquals,omitempty"`

	// timestamp less than
	TimestampLessThan *strfmt.DateTime `json:"TimestampLessThan,omitempty"`

	// timestamp less than equals
	TimestampLessThanEquals *strfmt.DateTime `json:"TimestampLessThanEquals,omitempty"`

	// variable
	Variable string `json:"Variable,omitempty"`
}

// Validate validates this s l choice
func (m *SLChoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNot(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOr(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLChoice) validateAnd(formats strfmt.Registry) error {

	if swag.IsZero(m.And) { // not required
		return nil
	}

	for i := 0; i < len(m.And); i++ {

		if swag.IsZero(m.And[i]) { // not required
			continue
		}

		if m.And[i] != nil {

			if err := m.And[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("And" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SLChoice) validateNot(formats strfmt.Registry) error {

	if swag.IsZero(m.Not) { // not required
		return nil
	}

	if m.Not != nil {

		if err := m.Not.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Not")
			}
			return err
		}
	}

	return nil
}

func (m *SLChoice) validateOr(formats strfmt.Registry) error {

	if swag.IsZero(m.Or) { // not required
		return nil
	}

	for i := 0; i < len(m.Or); i++ {

		if swag.IsZero(m.Or[i]) { // not required
			continue
		}

		if m.Or[i] != nil {

			if err := m.Or[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Or" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
