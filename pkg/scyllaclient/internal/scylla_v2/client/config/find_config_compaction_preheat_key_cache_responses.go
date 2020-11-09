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

// FindConfigCompactionPreheatKeyCacheReader is a Reader for the FindConfigCompactionPreheatKeyCache structure.
type FindConfigCompactionPreheatKeyCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCompactionPreheatKeyCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCompactionPreheatKeyCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCompactionPreheatKeyCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCompactionPreheatKeyCacheOK creates a FindConfigCompactionPreheatKeyCacheOK with default headers values
func NewFindConfigCompactionPreheatKeyCacheOK() *FindConfigCompactionPreheatKeyCacheOK {
	return &FindConfigCompactionPreheatKeyCacheOK{}
}

/*FindConfigCompactionPreheatKeyCacheOK handles this case with default header values.

Config value
*/
type FindConfigCompactionPreheatKeyCacheOK struct {
	Payload bool
}

func (o *FindConfigCompactionPreheatKeyCacheOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigCompactionPreheatKeyCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCompactionPreheatKeyCacheDefault creates a FindConfigCompactionPreheatKeyCacheDefault with default headers values
func NewFindConfigCompactionPreheatKeyCacheDefault(code int) *FindConfigCompactionPreheatKeyCacheDefault {
	return &FindConfigCompactionPreheatKeyCacheDefault{
		_statusCode: code,
	}
}

/*FindConfigCompactionPreheatKeyCacheDefault handles this case with default header values.

unexpected error
*/
type FindConfigCompactionPreheatKeyCacheDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config compaction preheat key cache default response
func (o *FindConfigCompactionPreheatKeyCacheDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCompactionPreheatKeyCacheDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCompactionPreheatKeyCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCompactionPreheatKeyCacheDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
