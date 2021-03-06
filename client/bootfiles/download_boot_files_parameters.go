// Code generated by go-swagger; DO NOT EDIT.

package bootfiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDownloadBootFilesParams creates a new DownloadBootFilesParams object
// with the default values initialized.
func NewDownloadBootFilesParams() *DownloadBootFilesParams {
	var ()
	return &DownloadBootFilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadBootFilesParamsWithTimeout creates a new DownloadBootFilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadBootFilesParamsWithTimeout(timeout time.Duration) *DownloadBootFilesParams {
	var ()
	return &DownloadBootFilesParams{

		timeout: timeout,
	}
}

// NewDownloadBootFilesParamsWithContext creates a new DownloadBootFilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadBootFilesParamsWithContext(ctx context.Context) *DownloadBootFilesParams {
	var ()
	return &DownloadBootFilesParams{

		Context: ctx,
	}
}

// NewDownloadBootFilesParamsWithHTTPClient creates a new DownloadBootFilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadBootFilesParamsWithHTTPClient(client *http.Client) *DownloadBootFilesParams {
	var ()
	return &DownloadBootFilesParams{
		HTTPClient: client,
	}
}

/*DownloadBootFilesParams contains all the parameters to send to the API endpoint
for the download boot files operation typically these are written to a http.Request
*/
type DownloadBootFilesParams struct {

	/*FileType
	  The file type to download.

	*/
	FileType string
	/*OpenshiftVersion
	  The corresponding OpenShift version for the boot file.

	*/
	OpenshiftVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download boot files params
func (o *DownloadBootFilesParams) WithTimeout(timeout time.Duration) *DownloadBootFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download boot files params
func (o *DownloadBootFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download boot files params
func (o *DownloadBootFilesParams) WithContext(ctx context.Context) *DownloadBootFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download boot files params
func (o *DownloadBootFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download boot files params
func (o *DownloadBootFilesParams) WithHTTPClient(client *http.Client) *DownloadBootFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download boot files params
func (o *DownloadBootFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileType adds the fileType to the download boot files params
func (o *DownloadBootFilesParams) WithFileType(fileType string) *DownloadBootFilesParams {
	o.SetFileType(fileType)
	return o
}

// SetFileType adds the fileType to the download boot files params
func (o *DownloadBootFilesParams) SetFileType(fileType string) {
	o.FileType = fileType
}

// WithOpenshiftVersion adds the openshiftVersion to the download boot files params
func (o *DownloadBootFilesParams) WithOpenshiftVersion(openshiftVersion string) *DownloadBootFilesParams {
	o.SetOpenshiftVersion(openshiftVersion)
	return o
}

// SetOpenshiftVersion adds the openshiftVersion to the download boot files params
func (o *DownloadBootFilesParams) SetOpenshiftVersion(openshiftVersion string) {
	o.OpenshiftVersion = openshiftVersion
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadBootFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param file_type
	qrFileType := o.FileType
	qFileType := qrFileType
	if qFileType != "" {
		if err := r.SetQueryParam("file_type", qFileType); err != nil {
			return err
		}
	}

	// query param openshift_version
	qrOpenshiftVersion := o.OpenshiftVersion
	qOpenshiftVersion := qrOpenshiftVersion
	if qOpenshiftVersion != "" {
		if err := r.SetQueryParam("openshift_version", qOpenshiftVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
