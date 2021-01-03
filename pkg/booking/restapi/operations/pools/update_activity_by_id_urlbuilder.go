// Code generated by go-swagger; DO NOT EDIT.

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// UpdateActivityByIDURL generates an URL for the update activity by ID operation
type UpdateActivityByIDURL struct {
	ActivityID string
	PoolID     string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdateActivityByIDURL) WithBasePath(bp string) *UpdateActivityByIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdateActivityByIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *UpdateActivityByIDURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/pools/{pool_id}/activities/{activity_id}"

	activityID := o.ActivityID
	if activityID != "" {
		_path = strings.Replace(_path, "{activity_id}", activityID, -1)
	} else {
		return nil, errors.New("activityId is required on UpdateActivityByIDURL")
	}

	poolID := o.PoolID
	if poolID != "" {
		_path = strings.Replace(_path, "{pool_id}", poolID, -1)
	} else {
		return nil, errors.New("poolId is required on UpdateActivityByIDURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UpdateActivityByIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UpdateActivityByIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UpdateActivityByIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UpdateActivityByIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UpdateActivityByIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *UpdateActivityByIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
