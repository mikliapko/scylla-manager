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

// StorageServiceSampleKeyRangeGetReader is a Reader for the StorageServiceSampleKeyRangeGet structure.
type StorageServiceSampleKeyRangeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceSampleKeyRangeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceSampleKeyRangeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceSampleKeyRangeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceSampleKeyRangeGetOK creates a StorageServiceSampleKeyRangeGetOK with default headers values
func NewStorageServiceSampleKeyRangeGetOK() *StorageServiceSampleKeyRangeGetOK {
	return &StorageServiceSampleKeyRangeGetOK{}
}

/*StorageServiceSampleKeyRangeGetOK handles this case with default header values.

StorageServiceSampleKeyRangeGetOK storage service sample key range get o k
*/
type StorageServiceSampleKeyRangeGetOK struct {
	Payload []string
}

func (o *StorageServiceSampleKeyRangeGetOK) GetPayload() []string {
	return o.Payload
}

func (o *StorageServiceSampleKeyRangeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceSampleKeyRangeGetDefault creates a StorageServiceSampleKeyRangeGetDefault with default headers values
func NewStorageServiceSampleKeyRangeGetDefault(code int) *StorageServiceSampleKeyRangeGetDefault {
	return &StorageServiceSampleKeyRangeGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceSampleKeyRangeGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceSampleKeyRangeGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service sample key range get default response
func (o *StorageServiceSampleKeyRangeGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceSampleKeyRangeGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceSampleKeyRangeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceSampleKeyRangeGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
