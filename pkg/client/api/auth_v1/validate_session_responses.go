// Code generated by go-swagger; DO NOT EDIT.

package auth_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ValidateSessionReader is a Reader for the ValidateSession structure.
type ValidateSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewValidateSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewValidateSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateSessionOK creates a ValidateSessionOK with default headers values
func NewValidateSessionOK() *ValidateSessionOK {
	return &ValidateSessionOK{}
}

/*ValidateSessionOK handles this case with default header values.

The token is valid
*/
type ValidateSessionOK struct {
}

func (o *ValidateSessionOK) Error() string {
	return fmt.Sprintf("[GET /auth/sessions][%d] validateSessionOK ", 200)
}

func (o *ValidateSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateSessionUnauthorized creates a ValidateSessionUnauthorized with default headers values
func NewValidateSessionUnauthorized() *ValidateSessionUnauthorized {
	return &ValidateSessionUnauthorized{}
}

/*ValidateSessionUnauthorized handles this case with default header values.

The token is invalid
*/
type ValidateSessionUnauthorized struct {
}

func (o *ValidateSessionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth/sessions][%d] validateSessionUnauthorized ", 401)
}

func (o *ValidateSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}