// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoginSuccess login success
//
// swagger:model LoginSuccess
type LoginSuccess struct {

	// success
	Success bool `json:"success,omitempty"`

	// token
	// Example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJkYXRlX2NyZWF0ZWRfYXQiOiIyMDIxLTA0LTE4VDE5OjQ5OjE4LjE0MzU5MzkxWiIsImRhdGVfZXhwaXJlZF9hdCI6IjIwMjEtMDQtMThUMjE6NDk6MTguMTQzNTk0NDk2WiIsImVtYWlsIjoic2FsdHlANDIuZnIiLCJleHAiOjE2MTg3ODI1NTgsImZuYW1lIjoiYWRyaWVuIiwibG5hbWUiOiJoYW5vdCIsImxvZ2luIjoic2FsdHkifQ.WS5hy8jE-Hdd3NFQcENK4XKAoQBoSVRR-WNaexPtNs4
	Token string `json:"token,omitempty"`
}

// Validate validates this login success
func (m *LoginSuccess) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this login success based on context it is used
func (m *LoginSuccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginSuccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginSuccess) UnmarshalBinary(b []byte) error {
	var res LoginSuccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
