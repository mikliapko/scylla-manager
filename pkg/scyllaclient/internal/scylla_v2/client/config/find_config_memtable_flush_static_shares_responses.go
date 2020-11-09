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

// FindConfigMemtableFlushStaticSharesReader is a Reader for the FindConfigMemtableFlushStaticShares structure.
type FindConfigMemtableFlushStaticSharesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMemtableFlushStaticSharesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMemtableFlushStaticSharesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMemtableFlushStaticSharesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMemtableFlushStaticSharesOK creates a FindConfigMemtableFlushStaticSharesOK with default headers values
func NewFindConfigMemtableFlushStaticSharesOK() *FindConfigMemtableFlushStaticSharesOK {
	return &FindConfigMemtableFlushStaticSharesOK{}
}

/*FindConfigMemtableFlushStaticSharesOK handles this case with default header values.

Config value
*/
type FindConfigMemtableFlushStaticSharesOK struct {
	Payload float64
}

func (o *FindConfigMemtableFlushStaticSharesOK) GetPayload() float64 {
	return o.Payload
}

func (o *FindConfigMemtableFlushStaticSharesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMemtableFlushStaticSharesDefault creates a FindConfigMemtableFlushStaticSharesDefault with default headers values
func NewFindConfigMemtableFlushStaticSharesDefault(code int) *FindConfigMemtableFlushStaticSharesDefault {
	return &FindConfigMemtableFlushStaticSharesDefault{
		_statusCode: code,
	}
}

/*FindConfigMemtableFlushStaticSharesDefault handles this case with default header values.

unexpected error
*/
type FindConfigMemtableFlushStaticSharesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config memtable flush static shares default response
func (o *FindConfigMemtableFlushStaticSharesDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMemtableFlushStaticSharesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMemtableFlushStaticSharesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigMemtableFlushStaticSharesDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
