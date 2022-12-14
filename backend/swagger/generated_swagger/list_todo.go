// Code generated by go-swagger; DO NOT EDIT.

package generated_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListTodo IndexTodo
//
// swagger:model ListTodo
type ListTodo struct {

	// todos
	Todos []*Todo `json:"todos"`
}

// Validate validates this list todo
func (m *ListTodo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTodos(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListTodo) validateTodos(formats strfmt.Registry) error {
	if swag.IsZero(m.Todos) { // not required
		return nil
	}

	for i := 0; i < len(m.Todos); i++ {
		if swag.IsZero(m.Todos[i]) { // not required
			continue
		}

		if m.Todos[i] != nil {
			if err := m.Todos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("todos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("todos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list todo based on the context it is used
func (m *ListTodo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTodos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListTodo) contextValidateTodos(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Todos); i++ {

		if m.Todos[i] != nil {
			if err := m.Todos[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("todos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("todos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListTodo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListTodo) UnmarshalBinary(b []byte) error {
	var res ListTodo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
