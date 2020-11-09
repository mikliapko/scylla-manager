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

// ColumnFamilyMetricsMemtableColumnsCountByNameGetReader is a Reader for the ColumnFamilyMetricsMemtableColumnsCountByNameGet structure.
type ColumnFamilyMetricsMemtableColumnsCountByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMemtableColumnsCountByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMemtableColumnsCountByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMemtableColumnsCountByNameGetOK creates a ColumnFamilyMetricsMemtableColumnsCountByNameGetOK with default headers values
func NewColumnFamilyMetricsMemtableColumnsCountByNameGetOK() *ColumnFamilyMetricsMemtableColumnsCountByNameGetOK {
	return &ColumnFamilyMetricsMemtableColumnsCountByNameGetOK{}
}

/*ColumnFamilyMetricsMemtableColumnsCountByNameGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableColumnsCountByNameGetOK column family metrics memtable columns count by name get o k
*/
type ColumnFamilyMetricsMemtableColumnsCountByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMemtableColumnsCountByNameGetDefault creates a ColumnFamilyMetricsMemtableColumnsCountByNameGetDefault with default headers values
func NewColumnFamilyMetricsMemtableColumnsCountByNameGetDefault(code int) *ColumnFamilyMetricsMemtableColumnsCountByNameGetDefault {
	return &ColumnFamilyMetricsMemtableColumnsCountByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsMemtableColumnsCountByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMemtableColumnsCountByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics memtable columns count by name get default response
func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMemtableColumnsCountByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
