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

// RequestSessionByPoolIDReader is a Reader for the RequestSessionByPoolID structure.
type RequestSessionByPoolIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestSessionByPoolIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRequestSessionByPoolIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRequestSessionByPoolIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewRequestSessionByPoolIDPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRequestSessionByPoolIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRequestSessionByPoolIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRequestSessionByPoolIDOK creates a RequestSessionByPoolIDOK with default headers values
func NewRequestSessionByPoolIDOK() *RequestSessionByPoolIDOK {
	return &RequestSessionByPoolIDOK{}
}

/*RequestSessionByPoolIDOK handles this case with default header values.

RequestSessionByPoolIDOK request session by pool Id o k
*/
type RequestSessionByPoolIDOK struct {
	Payload *models.Activity
}

func (o *RequestSessionByPoolIDOK) Error() string {
	return fmt.Sprintf("[POST /pools/{pool_id}/sessions][%d] requestSessionByPoolIdOK  %+v", 200, o.Payload)
}

func (o *RequestSessionByPoolIDOK) GetPayload() *models.Activity {
	return o.Payload
}

func (o *RequestSessionByPoolIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Activity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestSessionByPoolIDUnauthorized creates a RequestSessionByPoolIDUnauthorized with default headers values
func NewRequestSessionByPoolIDUnauthorized() *RequestSessionByPoolIDUnauthorized {
	return &RequestSessionByPoolIDUnauthorized{}
}

/*RequestSessionByPoolIDUnauthorized handles this case with default header values.

Unauthorized
*/
type RequestSessionByPoolIDUnauthorized struct {
	Payload interface{}
}

func (o *RequestSessionByPoolIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pools/{pool_id}/sessions][%d] requestSessionByPoolIdUnauthorized  %+v", 401, o.Payload)
}

func (o *RequestSessionByPoolIDUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *RequestSessionByPoolIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestSessionByPoolIDPaymentRequired creates a RequestSessionByPoolIDPaymentRequired with default headers values
func NewRequestSessionByPoolIDPaymentRequired() *RequestSessionByPoolIDPaymentRequired {
	return &RequestSessionByPoolIDPaymentRequired{}
}

/*RequestSessionByPoolIDPaymentRequired handles this case with default header values.

Quota Exceeded
*/
type RequestSessionByPoolIDPaymentRequired struct {
	Payload interface{}
}

func (o *RequestSessionByPoolIDPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /pools/{pool_id}/sessions][%d] requestSessionByPoolIdPaymentRequired  %+v", 402, o.Payload)
}

func (o *RequestSessionByPoolIDPaymentRequired) GetPayload() interface{} {
	return o.Payload
}

func (o *RequestSessionByPoolIDPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestSessionByPoolIDNotFound creates a RequestSessionByPoolIDNotFound with default headers values
func NewRequestSessionByPoolIDNotFound() *RequestSessionByPoolIDNotFound {
	return &RequestSessionByPoolIDNotFound{}
}

/*RequestSessionByPoolIDNotFound handles this case with default header values.

Not Available
*/
type RequestSessionByPoolIDNotFound struct {
	Payload interface{}
}

func (o *RequestSessionByPoolIDNotFound) Error() string {
	return fmt.Sprintf("[POST /pools/{pool_id}/sessions][%d] requestSessionByPoolIdNotFound  %+v", 404, o.Payload)
}

func (o *RequestSessionByPoolIDNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RequestSessionByPoolIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestSessionByPoolIDInternalServerError creates a RequestSessionByPoolIDInternalServerError with default headers values
func NewRequestSessionByPoolIDInternalServerError() *RequestSessionByPoolIDInternalServerError {
	return &RequestSessionByPoolIDInternalServerError{}
}

/*RequestSessionByPoolIDInternalServerError handles this case with default header values.

Internal Error
*/
type RequestSessionByPoolIDInternalServerError struct {
	Payload interface{}
}

func (o *RequestSessionByPoolIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pools/{pool_id}/sessions][%d] requestSessionByPoolIdInternalServerError  %+v", 500, o.Payload)
}

func (o *RequestSessionByPoolIDInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *RequestSessionByPoolIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
