// Code generated by go-swagger; DO NOT EDIT.

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timdrysdale/relay/pkg/bc/models"
)

// AddActivityByPoolIDReader is a Reader for the AddActivityByPoolID structure.
type AddActivityByPoolIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddActivityByPoolIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddActivityByPoolIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddActivityByPoolIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddActivityByPoolIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddActivityByPoolIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddActivityByPoolIDOK creates a AddActivityByPoolIDOK with default headers values
func NewAddActivityByPoolIDOK() *AddActivityByPoolIDOK {
	return &AddActivityByPoolIDOK{}
}

/*AddActivityByPoolIDOK handles this case with default header values.

AddActivityByPoolIDOK add activity by pool Id o k
*/
type AddActivityByPoolIDOK struct {
	Payload *models.ID
}

func (o *AddActivityByPoolIDOK) Error() string {
	return fmt.Sprintf("[POST /pools/{pool_id}/activities][%d] addActivityByPoolIdOK  %+v", 200, o.Payload)
}

func (o *AddActivityByPoolIDOK) GetPayload() *models.ID {
	return o.Payload
}

func (o *AddActivityByPoolIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddActivityByPoolIDUnauthorized creates a AddActivityByPoolIDUnauthorized with default headers values
func NewAddActivityByPoolIDUnauthorized() *AddActivityByPoolIDUnauthorized {
	return &AddActivityByPoolIDUnauthorized{}
}

/*AddActivityByPoolIDUnauthorized handles this case with default header values.

Unauthorized
*/
type AddActivityByPoolIDUnauthorized struct {
	Payload interface{}
}

func (o *AddActivityByPoolIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pools/{pool_id}/activities][%d] addActivityByPoolIdUnauthorized  %+v", 401, o.Payload)
}

func (o *AddActivityByPoolIDUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *AddActivityByPoolIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddActivityByPoolIDNotFound creates a AddActivityByPoolIDNotFound with default headers values
func NewAddActivityByPoolIDNotFound() *AddActivityByPoolIDNotFound {
	return &AddActivityByPoolIDNotFound{}
}

/*AddActivityByPoolIDNotFound handles this case with default header values.

Not Available
*/
type AddActivityByPoolIDNotFound struct {
	Payload interface{}
}

func (o *AddActivityByPoolIDNotFound) Error() string {
	return fmt.Sprintf("[POST /pools/{pool_id}/activities][%d] addActivityByPoolIdNotFound  %+v", 404, o.Payload)
}

func (o *AddActivityByPoolIDNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AddActivityByPoolIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddActivityByPoolIDInternalServerError creates a AddActivityByPoolIDInternalServerError with default headers values
func NewAddActivityByPoolIDInternalServerError() *AddActivityByPoolIDInternalServerError {
	return &AddActivityByPoolIDInternalServerError{}
}

/*AddActivityByPoolIDInternalServerError handles this case with default header values.

Internal Error
*/
type AddActivityByPoolIDInternalServerError struct {
	Payload interface{}
}

func (o *AddActivityByPoolIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pools/{pool_id}/activities][%d] addActivityByPoolIdInternalServerError  %+v", 500, o.Payload)
}

func (o *AddActivityByPoolIDInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *AddActivityByPoolIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
