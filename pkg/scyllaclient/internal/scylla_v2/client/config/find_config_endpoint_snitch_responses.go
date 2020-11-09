// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigEndpointSnitchReader is a Reader for the FindConfigEndpointSnitch structure.
type FindConfigEndpointSnitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigEndpointSnitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigEndpointSnitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigEndpointSnitchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigEndpointSnitchOK creates a FindConfigEndpointSnitchOK with default headers values
func NewFindConfigEndpointSnitchOK() *FindConfigEndpointSnitchOK {
	return &FindConfigEndpointSnitchOK{}
}

/*FindConfigEndpointSnitchOK handles this case with default header values.

Config value
*/
type FindConfigEndpointSnitchOK struct {
	Payload string
}

func (o *FindConfigEndpointSnitchOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigEndpointSnitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigEndpointSnitchDefault creates a FindConfigEndpointSnitchDefault with default headers values
func NewFindConfigEndpointSnitchDefault(code int) *FindConfigEndpointSnitchDefault {
	return &FindConfigEndpointSnitchDefault{
		_statusCode: code,
	}
}

/*FindConfigEndpointSnitchDefault handles this case with default header values.

unexpected error
*/
type FindConfigEndpointSnitchDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config endpoint snitch default response
func (o *FindConfigEndpointSnitchDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigEndpointSnitchDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigEndpointSnitchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigEndpointSnitchDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
