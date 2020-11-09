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

// StorageServiceForceTerminateRepairPostReader is a Reader for the StorageServiceForceTerminateRepairPost structure.
type StorageServiceForceTerminateRepairPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceForceTerminateRepairPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceForceTerminateRepairPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceForceTerminateRepairPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceForceTerminateRepairPostOK creates a StorageServiceForceTerminateRepairPostOK with default headers values
func NewStorageServiceForceTerminateRepairPostOK() *StorageServiceForceTerminateRepairPostOK {
	return &StorageServiceForceTerminateRepairPostOK{}
}

/*StorageServiceForceTerminateRepairPostOK handles this case with default header values.

StorageServiceForceTerminateRepairPostOK storage service force terminate repair post o k
*/
type StorageServiceForceTerminateRepairPostOK struct {
}

func (o *StorageServiceForceTerminateRepairPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceForceTerminateRepairPostDefault creates a StorageServiceForceTerminateRepairPostDefault with default headers values
func NewStorageServiceForceTerminateRepairPostDefault(code int) *StorageServiceForceTerminateRepairPostDefault {
	return &StorageServiceForceTerminateRepairPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceForceTerminateRepairPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceForceTerminateRepairPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service force terminate repair post default response
func (o *StorageServiceForceTerminateRepairPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceForceTerminateRepairPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceForceTerminateRepairPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceForceTerminateRepairPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
