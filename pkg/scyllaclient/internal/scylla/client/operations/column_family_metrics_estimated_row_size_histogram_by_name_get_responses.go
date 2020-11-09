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

// ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGet structure.
type ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK creates a ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK() *ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK {
	return &ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK column family metrics estimated row size histogram by name get o k
*/
type ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault creates a ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault(code int) *ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault {
	return &ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics estimated row size histogram by name get default response
func (o *ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsEstimatedRowSizeHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
