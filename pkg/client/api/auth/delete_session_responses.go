// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// DeleteSessionReader is a Reader for the DeleteSession structure.
type DeleteSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSessionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSessionOK creates a DeleteSessionOK with default headers values
func NewDeleteSessionOK() *DeleteSessionOK {
	return &DeleteSessionOK{}
}

/*DeleteSessionOK handles this case with default header values.

The operation completed successfully
*/
type DeleteSessionOK struct {
	Payload *DeleteSessionOKBody
}

func (o *DeleteSessionOK) Error() string {
	return fmt.Sprintf("[DELETE /auth/sessions][%d] deleteSessionOK  %+v", 200, o.Payload)
}

func (o *DeleteSessionOK) GetPayload() *DeleteSessionOKBody {
	return o.Payload
}

func (o *DeleteSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteSessionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSessionDefault creates a DeleteSessionDefault with default headers values
func NewDeleteSessionDefault(code int) *DeleteSessionDefault {
	return &DeleteSessionDefault{
		_statusCode: code,
	}
}

/*DeleteSessionDefault handles this case with default header values.

An error occurred
*/
type DeleteSessionDefault struct {
	_statusCode int

	Payload *DeleteSessionDefaultBody
}

// Code gets the status code for the delete session default response
func (o *DeleteSessionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSessionDefault) Error() string {
	return fmt.Sprintf("[DELETE /auth/sessions][%d] deleteSession default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSessionDefault) GetPayload() *DeleteSessionDefaultBody {
	return o.Payload
}

func (o *DeleteSessionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteSessionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteSessionDefaultBody Error response
swagger:model DeleteSessionDefaultBody
*/
type DeleteSessionDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this delete session default body
func (o *DeleteSessionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteSessionDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteSession default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteSessionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteSessionDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteSessionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteSessionOKBody Success response
swagger:model DeleteSessionOKBody
*/
type DeleteSessionOKBody struct {

	// Did this succeed?
	Success bool `json:"success,omitempty"`
}

// Validate validates this delete session o k body
func (o *DeleteSessionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteSessionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteSessionOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteSessionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}