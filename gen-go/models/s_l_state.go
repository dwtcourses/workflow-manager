package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// SLState s l state
// swagger:model SLState
type SLState struct {

	// comment
	Comment string `json:"Comment,omitempty"`

	// end
	End bool `json:"End,omitempty"`

	// heartbeat seconds
	HeartbeatSeconds int64 `json:"HeartbeatSeconds,omitempty"`

	// input path
	InputPath string `json:"InputPath,omitempty"`

	// next
	Next string `json:"Next,omitempty"`

	// output path
	OutputPath string `json:"OutputPath,omitempty"`

	// resource
	Resource string `json:"Resource,omitempty"`

	// result path
	ResultPath string `json:"ResultPath,omitempty"`

	// retry
	Retry []*SLRetrier `json:"Retry"`

	// timeout seconds
	TimeoutSeconds int64 `json:"TimeoutSeconds,omitempty"`

	// type
	Type SLStateType `json:"Type,omitempty"`
}

// Validate validates this s l state
func (m *SLState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRetry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLState) validateRetry(formats strfmt.Registry) error {

	if swag.IsZero(m.Retry) { // not required
		return nil
	}

	for i := 0; i < len(m.Retry); i++ {

		if swag.IsZero(m.Retry[i]) { // not required
			continue
		}

		if m.Retry[i] != nil {

			if err := m.Retry[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SLState) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		return err
	}

	return nil
}