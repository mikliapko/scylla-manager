// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/agent/models"
)

// OperationsCopyfileReader is a Reader for the OperationsCopyfile structure.
type OperationsCopyfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationsCopyfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationsCopyfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOperationsCopyfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOperationsCopyfileOK creates a OperationsCopyfileOK with default headers values
func NewOperationsCopyfileOK() *OperationsCopyfileOK {
	return &OperationsCopyfileOK{}
}

/*OperationsCopyfileOK handles this case with default header values.

Job
*/
type OperationsCopyfileOK struct {
	Payload *models.Jobid
	JobID   int64
}

func (o *OperationsCopyfileOK) GetPayload() *models.Jobid {
	return o.Payload
}

func (o *OperationsCopyfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Jobid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewOperationsCopyfileDefault creates a OperationsCopyfileDefault with default headers values
func NewOperationsCopyfileDefault(code int) *OperationsCopyfileDefault {
	return &OperationsCopyfileDefault{
		_statusCode: code,
	}
}

/*OperationsCopyfileDefault handles this case with default header values.

Server error
*/
type OperationsCopyfileDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the operations copyfile default response
func (o *OperationsCopyfileDefault) Code() int {
	return o._statusCode
}

func (o *OperationsCopyfileDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsCopyfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *OperationsCopyfileDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
