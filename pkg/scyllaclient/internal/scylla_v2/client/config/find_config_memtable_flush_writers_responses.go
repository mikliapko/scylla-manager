// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigMemtableFlushWritersReader is a Reader for the FindConfigMemtableFlushWriters structure.
type FindConfigMemtableFlushWritersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMemtableFlushWritersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMemtableFlushWritersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMemtableFlushWritersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMemtableFlushWritersOK creates a FindConfigMemtableFlushWritersOK with default headers values
func NewFindConfigMemtableFlushWritersOK() *FindConfigMemtableFlushWritersOK {
	return &FindConfigMemtableFlushWritersOK{}
}

/*FindConfigMemtableFlushWritersOK handles this case with default header values.

Config value
*/
type FindConfigMemtableFlushWritersOK struct {
	Payload int64
}

func (o *FindConfigMemtableFlushWritersOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigMemtableFlushWritersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMemtableFlushWritersDefault creates a FindConfigMemtableFlushWritersDefault with default headers values
func NewFindConfigMemtableFlushWritersDefault(code int) *FindConfigMemtableFlushWritersDefault {
	return &FindConfigMemtableFlushWritersDefault{
		_statusCode: code,
	}
}

/*FindConfigMemtableFlushWritersDefault handles this case with default header values.

unexpected error
*/
type FindConfigMemtableFlushWritersDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config memtable flush writers default response
func (o *FindConfigMemtableFlushWritersDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMemtableFlushWritersDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMemtableFlushWritersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigMemtableFlushWritersDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
