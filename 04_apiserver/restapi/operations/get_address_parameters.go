// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAddressParams creates a new GetAddressParams object
// with the default values initialized.
func NewGetAddressParams() GetAddressParams {

	var (
		// initialize parameters with default values

		zipcodeDefault = float64(1.638001e+06)
	)

	return GetAddressParams{
		Zipcode: &zipcodeDefault,
	}
}

// GetAddressParams contains all the bound params for the get address operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAddress
type GetAddressParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Default: 1.638001e+06
	*/
	Zipcode *float64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAddressParams() beforehand.
func (o *GetAddressParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qZipcode, qhkZipcode, _ := qs.GetOK("zipcode")
	if err := o.bindZipcode(qZipcode, qhkZipcode, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindZipcode binds and validates parameter Zipcode from query.
func (o *GetAddressParams) bindZipcode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAddressParams()
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("zipcode", "query", "float64", raw)
	}
	o.Zipcode = &value

	return nil
}
