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

// StorageServiceSstablesByKeyspacePostReader is a Reader for the StorageServiceSstablesByKeyspacePost structure.
type StorageServiceSstablesByKeyspacePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceSstablesByKeyspacePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceSstablesByKeyspacePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceSstablesByKeyspacePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceSstablesByKeyspacePostOK creates a StorageServiceSstablesByKeyspacePostOK with default headers values
func NewStorageServiceSstablesByKeyspacePostOK() *StorageServiceSstablesByKeyspacePostOK {
	return &StorageServiceSstablesByKeyspacePostOK{}
}

/*StorageServiceSstablesByKeyspacePostOK handles this case with default header values.

StorageServiceSstablesByKeyspacePostOK storage service sstables by keyspace post o k
*/
type StorageServiceSstablesByKeyspacePostOK struct {
}

func (o *StorageServiceSstablesByKeyspacePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceSstablesByKeyspacePostDefault creates a StorageServiceSstablesByKeyspacePostDefault with default headers values
func NewStorageServiceSstablesByKeyspacePostDefault(code int) *StorageServiceSstablesByKeyspacePostDefault {
	return &StorageServiceSstablesByKeyspacePostDefault{
		_statusCode: code,
	}
}

/*StorageServiceSstablesByKeyspacePostDefault handles this case with default header values.

internal server error
*/
type StorageServiceSstablesByKeyspacePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service sstables by keyspace post default response
func (o *StorageServiceSstablesByKeyspacePostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceSstablesByKeyspacePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceSstablesByKeyspacePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceSstablesByKeyspacePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
