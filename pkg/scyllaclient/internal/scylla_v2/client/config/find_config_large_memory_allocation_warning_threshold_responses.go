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

// FindConfigLargeMemoryAllocationWarningThresholdReader is a Reader for the FindConfigLargeMemoryAllocationWarningThreshold structure.
type FindConfigLargeMemoryAllocationWarningThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigLargeMemoryAllocationWarningThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigLargeMemoryAllocationWarningThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigLargeMemoryAllocationWarningThresholdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigLargeMemoryAllocationWarningThresholdOK creates a FindConfigLargeMemoryAllocationWarningThresholdOK with default headers values
func NewFindConfigLargeMemoryAllocationWarningThresholdOK() *FindConfigLargeMemoryAllocationWarningThresholdOK {
	return &FindConfigLargeMemoryAllocationWarningThresholdOK{}
}

/*FindConfigLargeMemoryAllocationWarningThresholdOK handles this case with default header values.

Config value
*/
type FindConfigLargeMemoryAllocationWarningThresholdOK struct {
	Payload int64
}

func (o *FindConfigLargeMemoryAllocationWarningThresholdOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigLargeMemoryAllocationWarningThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigLargeMemoryAllocationWarningThresholdDefault creates a FindConfigLargeMemoryAllocationWarningThresholdDefault with default headers values
func NewFindConfigLargeMemoryAllocationWarningThresholdDefault(code int) *FindConfigLargeMemoryAllocationWarningThresholdDefault {
	return &FindConfigLargeMemoryAllocationWarningThresholdDefault{
		_statusCode: code,
	}
}

/*FindConfigLargeMemoryAllocationWarningThresholdDefault handles this case with default header values.

unexpected error
*/
type FindConfigLargeMemoryAllocationWarningThresholdDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config large memory allocation warning threshold default response
func (o *FindConfigLargeMemoryAllocationWarningThresholdDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigLargeMemoryAllocationWarningThresholdDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigLargeMemoryAllocationWarningThresholdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigLargeMemoryAllocationWarningThresholdDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
