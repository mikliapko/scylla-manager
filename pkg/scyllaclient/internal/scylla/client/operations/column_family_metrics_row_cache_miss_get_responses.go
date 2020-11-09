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

// ColumnFamilyMetricsRowCacheMissGetReader is a Reader for the ColumnFamilyMetricsRowCacheMissGet structure.
type ColumnFamilyMetricsRowCacheMissGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsRowCacheMissGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsRowCacheMissGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsRowCacheMissGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsRowCacheMissGetOK creates a ColumnFamilyMetricsRowCacheMissGetOK with default headers values
func NewColumnFamilyMetricsRowCacheMissGetOK() *ColumnFamilyMetricsRowCacheMissGetOK {
	return &ColumnFamilyMetricsRowCacheMissGetOK{}
}

/*ColumnFamilyMetricsRowCacheMissGetOK handles this case with default header values.

ColumnFamilyMetricsRowCacheMissGetOK column family metrics row cache miss get o k
*/
type ColumnFamilyMetricsRowCacheMissGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsRowCacheMissGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsRowCacheMissGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsRowCacheMissGetDefault creates a ColumnFamilyMetricsRowCacheMissGetDefault with default headers values
func NewColumnFamilyMetricsRowCacheMissGetDefault(code int) *ColumnFamilyMetricsRowCacheMissGetDefault {
	return &ColumnFamilyMetricsRowCacheMissGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsRowCacheMissGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsRowCacheMissGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics row cache miss get default response
func (o *ColumnFamilyMetricsRowCacheMissGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsRowCacheMissGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsRowCacheMissGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsRowCacheMissGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
