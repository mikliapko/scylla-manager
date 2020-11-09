// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla/models"
)

// FailureDetectorEndpointsStatesByAddrGetReader is a Reader for the FailureDetectorEndpointsStatesByAddrGet structure.
type FailureDetectorEndpointsStatesByAddrGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FailureDetectorEndpointsStatesByAddrGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFailureDetectorEndpointsStatesByAddrGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFailureDetectorEndpointsStatesByAddrGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFailureDetectorEndpointsStatesByAddrGetOK creates a FailureDetectorEndpointsStatesByAddrGetOK with default headers values
func NewFailureDetectorEndpointsStatesByAddrGetOK() *FailureDetectorEndpointsStatesByAddrGetOK {
	return &FailureDetectorEndpointsStatesByAddrGetOK{}
}

/*FailureDetectorEndpointsStatesByAddrGetOK handles this case with default header values.

FailureDetectorEndpointsStatesByAddrGetOK failure detector endpoints states by addr get o k
*/
type FailureDetectorEndpointsStatesByAddrGetOK struct {
	Payload string
}

func (o *FailureDetectorEndpointsStatesByAddrGetOK) GetPayload() string {
	return o.Payload
}

func (o *FailureDetectorEndpointsStatesByAddrGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFailureDetectorEndpointsStatesByAddrGetDefault creates a FailureDetectorEndpointsStatesByAddrGetDefault with default headers values
func NewFailureDetectorEndpointsStatesByAddrGetDefault(code int) *FailureDetectorEndpointsStatesByAddrGetDefault {
	return &FailureDetectorEndpointsStatesByAddrGetDefault{
		_statusCode: code,
	}
}

/*FailureDetectorEndpointsStatesByAddrGetDefault handles this case with default header values.

internal server error
*/
type FailureDetectorEndpointsStatesByAddrGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the failure detector endpoints states by addr get default response
func (o *FailureDetectorEndpointsStatesByAddrGetDefault) Code() int {
	return o._statusCode
}

func (o *FailureDetectorEndpointsStatesByAddrGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FailureDetectorEndpointsStatesByAddrGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FailureDetectorEndpointsStatesByAddrGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
