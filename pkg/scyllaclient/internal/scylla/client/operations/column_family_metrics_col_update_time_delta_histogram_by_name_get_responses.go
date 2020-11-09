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

// ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGet structure.
type ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK creates a ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK() *ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK {
	return &ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK column family metrics col update time delta histogram by name get o k
*/
type ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault creates a ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault(code int) *ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault {
	return &ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics col update time delta histogram by name get default response
func (o *ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsColUpdateTimeDeltaHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
