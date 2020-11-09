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

// StorageServiceMetricsExceptionsGetReader is a Reader for the StorageServiceMetricsExceptionsGet structure.
type StorageServiceMetricsExceptionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceMetricsExceptionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceMetricsExceptionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceMetricsExceptionsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceMetricsExceptionsGetOK creates a StorageServiceMetricsExceptionsGetOK with default headers values
func NewStorageServiceMetricsExceptionsGetOK() *StorageServiceMetricsExceptionsGetOK {
	return &StorageServiceMetricsExceptionsGetOK{}
}

/*StorageServiceMetricsExceptionsGetOK handles this case with default header values.

StorageServiceMetricsExceptionsGetOK storage service metrics exceptions get o k
*/
type StorageServiceMetricsExceptionsGetOK struct {
	Payload int32
}

func (o *StorageServiceMetricsExceptionsGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageServiceMetricsExceptionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceMetricsExceptionsGetDefault creates a StorageServiceMetricsExceptionsGetDefault with default headers values
func NewStorageServiceMetricsExceptionsGetDefault(code int) *StorageServiceMetricsExceptionsGetDefault {
	return &StorageServiceMetricsExceptionsGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceMetricsExceptionsGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceMetricsExceptionsGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service metrics exceptions get default response
func (o *StorageServiceMetricsExceptionsGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceMetricsExceptionsGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceMetricsExceptionsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceMetricsExceptionsGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
