// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// ListClusterManifestsReader is a Reader for the ListClusterManifests structure.
type ListClusterManifestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterManifestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterManifestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClusterManifestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterManifestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListClusterManifestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewListClusterManifestsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListClusterManifestsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListClusterManifestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListClusterManifestsOK creates a ListClusterManifestsOK with default headers values
func NewListClusterManifestsOK() *ListClusterManifestsOK {
	return &ListClusterManifestsOK{}
}

/*ListClusterManifestsOK handles this case with default header values.

Success.
*/
type ListClusterManifestsOK struct {
	Payload models.ListManifests
}

func (o *ListClusterManifestsOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/manifests][%d] listClusterManifestsOK  %+v", 200, o.Payload)
}

func (o *ListClusterManifestsOK) GetPayload() models.ListManifests {
	return o.Payload
}

func (o *ListClusterManifestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterManifestsUnauthorized creates a ListClusterManifestsUnauthorized with default headers values
func NewListClusterManifestsUnauthorized() *ListClusterManifestsUnauthorized {
	return &ListClusterManifestsUnauthorized{}
}

/*ListClusterManifestsUnauthorized handles this case with default header values.

Unauthorized.
*/
type ListClusterManifestsUnauthorized struct {
	Payload *models.InfraError
}

func (o *ListClusterManifestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/manifests][%d] listClusterManifestsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListClusterManifestsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *ListClusterManifestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterManifestsForbidden creates a ListClusterManifestsForbidden with default headers values
func NewListClusterManifestsForbidden() *ListClusterManifestsForbidden {
	return &ListClusterManifestsForbidden{}
}

/*ListClusterManifestsForbidden handles this case with default header values.

Forbidden.
*/
type ListClusterManifestsForbidden struct {
	Payload *models.InfraError
}

func (o *ListClusterManifestsForbidden) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/manifests][%d] listClusterManifestsForbidden  %+v", 403, o.Payload)
}

func (o *ListClusterManifestsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *ListClusterManifestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterManifestsNotFound creates a ListClusterManifestsNotFound with default headers values
func NewListClusterManifestsNotFound() *ListClusterManifestsNotFound {
	return &ListClusterManifestsNotFound{}
}

/*ListClusterManifestsNotFound handles this case with default header values.

Error.
*/
type ListClusterManifestsNotFound struct {
	Payload *models.Error
}

func (o *ListClusterManifestsNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/manifests][%d] listClusterManifestsNotFound  %+v", 404, o.Payload)
}

func (o *ListClusterManifestsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClusterManifestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterManifestsMethodNotAllowed creates a ListClusterManifestsMethodNotAllowed with default headers values
func NewListClusterManifestsMethodNotAllowed() *ListClusterManifestsMethodNotAllowed {
	return &ListClusterManifestsMethodNotAllowed{}
}

/*ListClusterManifestsMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type ListClusterManifestsMethodNotAllowed struct {
	Payload *models.Error
}

func (o *ListClusterManifestsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/manifests][%d] listClusterManifestsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListClusterManifestsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClusterManifestsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterManifestsConflict creates a ListClusterManifestsConflict with default headers values
func NewListClusterManifestsConflict() *ListClusterManifestsConflict {
	return &ListClusterManifestsConflict{}
}

/*ListClusterManifestsConflict handles this case with default header values.

Error.
*/
type ListClusterManifestsConflict struct {
	Payload *models.Error
}

func (o *ListClusterManifestsConflict) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/manifests][%d] listClusterManifestsConflict  %+v", 409, o.Payload)
}

func (o *ListClusterManifestsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClusterManifestsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterManifestsInternalServerError creates a ListClusterManifestsInternalServerError with default headers values
func NewListClusterManifestsInternalServerError() *ListClusterManifestsInternalServerError {
	return &ListClusterManifestsInternalServerError{}
}

/*ListClusterManifestsInternalServerError handles this case with default header values.

Error.
*/
type ListClusterManifestsInternalServerError struct {
	Payload *models.Error
}

func (o *ListClusterManifestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/manifests][%d] listClusterManifestsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListClusterManifestsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClusterManifestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}