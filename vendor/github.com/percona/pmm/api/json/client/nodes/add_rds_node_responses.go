// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm/api/json/models"
)

// AddRDSNodeReader is a Reader for the AddRDSNode structure.
type AddRDSNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRDSNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddRDSNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddRDSNodeOK creates a AddRDSNodeOK with default headers values
func NewAddRDSNodeOK() *AddRDSNodeOK {
	return &AddRDSNodeOK{}
}

/*AddRDSNodeOK handles this case with default header values.

(empty)
*/
type AddRDSNodeOK struct {
	Payload *models.InventoryAddRDSNodeResponse
}

func (o *AddRDSNodeOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Nodes/AddRDSNode][%d] addRdsNodeOK  %+v", 200, o.Payload)
}

func (o *AddRDSNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryAddRDSNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}