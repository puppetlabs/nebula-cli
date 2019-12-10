// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowRevision workflow revision
// swagger:model WorkflowRevision
type WorkflowRevision struct {
	Lifecycle

	WorkflowRevisionSummary

	// context
	Context *WorkflowSourceContext `json:"context,omitempty"`

	// parameters
	Parameters WorkflowParameters `json:"parameters,omitempty"`

	// Raw YAML representation of the yaml
	// Required: true
	Raw *string `json:"raw"`

	// A list of parsed information about workflow steps
	// Required: true
	Steps []WorkflowStep `json:"steps"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowRevision) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Lifecycle
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Lifecycle = aO0

	// AO1
	var aO1 WorkflowRevisionSummary
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.WorkflowRevisionSummary = aO1

	// AO2
	var dataAO2 struct {
		Context *WorkflowSourceContext `json:"context,omitempty"`

		Parameters WorkflowParameters `json:"parameters,omitempty"`

		Raw *string `json:"raw"`

		Steps []WorkflowStep `json:"steps"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.Context = dataAO2.Context

	m.Parameters = dataAO2.Parameters

	m.Raw = dataAO2.Raw

	m.Steps = dataAO2.Steps

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowRevision) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.Lifecycle)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.WorkflowRevisionSummary)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	var dataAO2 struct {
		Context *WorkflowSourceContext `json:"context,omitempty"`

		Parameters WorkflowParameters `json:"parameters,omitempty"`

		Raw *string `json:"raw"`

		Steps []WorkflowStep `json:"steps"`
	}

	dataAO2.Context = m.Context

	dataAO2.Parameters = m.Parameters

	dataAO2.Raw = m.Raw

	dataAO2.Steps = m.Steps

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow revision
func (m *WorkflowRevision) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Lifecycle
	if err := m.Lifecycle.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with WorkflowRevisionSummary
	if err := m.WorkflowRevisionSummary.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowRevision) validateContext(formats strfmt.Registry) error {

	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {
		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowRevision) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if err := m.Parameters.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("parameters")
		}
		return err
	}

	return nil
}

func (m *WorkflowRevision) validateRaw(formats strfmt.Registry) error {

	if err := validate.Required("raw", "body", m.Raw); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowRevision) validateSteps(formats strfmt.Registry) error {

	if err := validate.Required("steps", "body", m.Steps); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowRevision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowRevision) UnmarshalBinary(b []byte) error {
	var res WorkflowRevision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
