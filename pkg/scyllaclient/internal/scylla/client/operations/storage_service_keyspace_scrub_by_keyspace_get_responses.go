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

// StorageServiceKeyspaceScrubByKeyspaceGetReader is a Reader for the StorageServiceKeyspaceScrubByKeyspaceGet structure.
type StorageServiceKeyspaceScrubByKeyspaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceKeyspaceScrubByKeyspaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceKeyspaceScrubByKeyspaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceKeyspaceScrubByKeyspaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceKeyspaceScrubByKeyspaceGetOK creates a StorageServiceKeyspaceScrubByKeyspaceGetOK with default headers values
func NewStorageServiceKeyspaceScrubByKeyspaceGetOK() *StorageServiceKeyspaceScrubByKeyspaceGetOK {
	return &StorageServiceKeyspaceScrubByKeyspaceGetOK{}
}

/*StorageServiceKeyspaceScrubByKeyspaceGetOK handles this case with default header values.

StorageServiceKeyspaceScrubByKeyspaceGetOK storage service keyspace scrub by keyspace get o k
*/
type StorageServiceKeyspaceScrubByKeyspaceGetOK struct {
	Payload int32
}

func (o *StorageServiceKeyspaceScrubByKeyspaceGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageServiceKeyspaceScrubByKeyspaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceKeyspaceScrubByKeyspaceGetDefault creates a StorageServiceKeyspaceScrubByKeyspaceGetDefault with default headers values
func NewStorageServiceKeyspaceScrubByKeyspaceGetDefault(code int) *StorageServiceKeyspaceScrubByKeyspaceGetDefault {
	return &StorageServiceKeyspaceScrubByKeyspaceGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceKeyspaceScrubByKeyspaceGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceKeyspaceScrubByKeyspaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service keyspace scrub by keyspace get default response
func (o *StorageServiceKeyspaceScrubByKeyspaceGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceKeyspaceScrubByKeyspaceGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceKeyspaceScrubByKeyspaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceKeyspaceScrubByKeyspaceGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
