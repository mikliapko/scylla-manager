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

// StorageServiceMovePostReader is a Reader for the StorageServiceMovePost structure.
type StorageServiceMovePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceMovePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceMovePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceMovePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceMovePostOK creates a StorageServiceMovePostOK with default headers values
func NewStorageServiceMovePostOK() *StorageServiceMovePostOK {
	return &StorageServiceMovePostOK{}
}

/*StorageServiceMovePostOK handles this case with default header values.

StorageServiceMovePostOK storage service move post o k
*/
type StorageServiceMovePostOK struct {
}

func (o *StorageServiceMovePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceMovePostDefault creates a StorageServiceMovePostDefault with default headers values
func NewStorageServiceMovePostDefault(code int) *StorageServiceMovePostDefault {
	return &StorageServiceMovePostDefault{
		_statusCode: code,
	}
}

/*StorageServiceMovePostDefault handles this case with default header values.

internal server error
*/
type StorageServiceMovePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service move post default response
func (o *StorageServiceMovePostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceMovePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceMovePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceMovePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
