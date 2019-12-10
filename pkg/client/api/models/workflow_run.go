// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowRun workflow run
// swagger:model WorkflowRun
type WorkflowRun struct {
	WorkflowRunIdentifier

	Lifecycle

	// error
	Error *Error `json:"error,omitempty"`

	// parameters
	Parameters WorkflowRunParameters `json:"parameters,omitempty"`

	// revision
	// Required: true
	Revision *WorkflowRevisionSummary `json:"revision"`

	// Secret names provided to the run, both used and unused
	Secrets []*WorkflowSecretSummary `json:"secrets"`

	// state
	// Required: true
	State *WorkflowRunState `json:"state"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowRun) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowRunIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowRunIdentifier = aO0

	// AO1
	var aO1 Lifecycle
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Lifecycle = aO1

	// AO2
	var dataAO2 struct {
		Error *Error `json:"error,omitempty"`

		Parameters WorkflowRunParameters `json:"parameters,omitempty"`

		Revision *WorkflowRevisionSummary `json:"revision"`

		Secrets []*WorkflowSecretSummary `json:"secrets"`

		State *WorkflowRunState `json:"state"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.Error = dataAO2.Error

	m.Parameters = dataAO2.Parameters

	m.Revision = dataAO2.Revision

	m.Secrets = dataAO2.Secrets

	m.State = dataAO2.State

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowRun) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.WorkflowRunIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.Lifecycle)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	var dataAO2 struct {
		Error *Error `json:"error,omitempty"`

		Parameters WorkflowRunParameters `json:"parameters,omitempty"`

		Revision *WorkflowRevisionSummary `json:"revision"`

		Secrets []*WorkflowSecretSummary `json:"secrets"`

		State *WorkflowRunState `json:"state"`
	}

	dataAO2.Error = m.Error

	dataAO2.Parameters = m.Parameters

	dataAO2.Revision = m.Revision

	dataAO2.Secrets = m.Secrets

	dataAO2.State = m.State

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow run
func (m *WorkflowRun) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowRunIdentifier
	if err := m.WorkflowRunIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with Lifecycle
	if err := m.Lifecycle.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecrets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowRun) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowRun) validateParameters(formats strfmt.Registry) error {

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

func (m *WorkflowRun) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowRun) validateSecrets(formats strfmt.Registry) error {

	if swag.IsZero(m.Secrets) { // not required
		return nil
	}

	for i := 0; i < len(m.Secrets); i++ {
		if swag.IsZero(m.Secrets[i]) { // not required
			continue
		}

		if m.Secrets[i] != nil {
			if err := m.Secrets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowRun) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowRun) UnmarshalBinary(b []byte) error {
	var res WorkflowRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
