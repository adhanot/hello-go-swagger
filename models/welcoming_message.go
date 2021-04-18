// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WelcomingMessage welcoming message
//
// swagger:model WelcomingMessage
type WelcomingMessage struct {

	// message
	// Example: Hello, welcome on the go-swagger framework
	Message string `json:"message,omitempty"`
}

// Validate validates this welcoming message
func (m *WelcomingMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this welcoming message based on context it is used
func (m *WelcomingMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WelcomingMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WelcomingMessage) UnmarshalBinary(b []byte) error {
	var res WelcomingMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}