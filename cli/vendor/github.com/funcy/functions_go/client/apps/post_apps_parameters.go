package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/funcy/functions_go/models"
)

// NewPostAppsParams creates a new PostAppsParams object
// with the default values initialized.
func NewPostAppsParams() *PostAppsParams {
	var ()
	return &PostAppsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAppsParamsWithTimeout creates a new PostAppsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAppsParamsWithTimeout(timeout time.Duration) *PostAppsParams {
	var ()
	return &PostAppsParams{

		timeout: timeout,
	}
}

// NewPostAppsParamsWithContext creates a new PostAppsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAppsParamsWithContext(ctx context.Context) *PostAppsParams {
	var ()
	return &PostAppsParams{

		Context: ctx,
	}
}

// NewPostAppsParamsWithHTTPClient creates a new PostAppsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAppsParamsWithHTTPClient(client *http.Client) *PostAppsParams {
	var ()
	return &PostAppsParams{
		HTTPClient: client,
	}
}

/*PostAppsParams contains all the parameters to send to the API endpoint
for the post apps operation typically these are written to a http.Request
*/
type PostAppsParams struct {

	/*Body
	  App to post.

	*/
	Body *models.AppWrapper

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post apps params
func (o *PostAppsParams) WithTimeout(timeout time.Duration) *PostAppsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post apps params
func (o *PostAppsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post apps params
func (o *PostAppsParams) WithContext(ctx context.Context) *PostAppsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post apps params
func (o *PostAppsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post apps params
func (o *PostAppsParams) WithHTTPClient(client *http.Client) *PostAppsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post apps params
func (o *PostAppsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post apps params
func (o *PostAppsParams) WithBody(body *models.AppWrapper) *PostAppsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post apps params
func (o *PostAppsParams) SetBody(body *models.AppWrapper) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAppsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.AppWrapper)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}