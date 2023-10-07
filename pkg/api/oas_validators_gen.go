// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *IdPSignupRequestSchema) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        true,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Email)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *OPJWKSetResponseSchema) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Keys == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "keys",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *OPRevokeRequestSchema) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.TokenTypeHint.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "token_type_hint",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s OPRevokeRequestSchemaTokenTypeHint) Validate() error {
	switch s {
	case "access_token":
		return nil
	case "refresh_token":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *OPTokenRequestSchema) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.GrantType.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "grant_type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s OPTokenRequestSchemaGrantType) Validate() error {
	switch s {
	case "authorization_code":
		return nil
	case "refresh_token":
		return nil
	case "client_credentials":
		return nil
	case "password":
		return nil
	case "urn:ietf:params:oauth:grant-type:device_code":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s OpAuthorizeResponseType) Validate() error {
	switch s {
	case "code":
		return nil
	case "id_token":
		return nil
	case "token":
		return nil
	case "code id_token":
		return nil
	case "code token":
		return nil
	case "id_token token":
		return nil
	case "code id_token token":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s OpAuthorizeScope) Validate() error {
	switch s {
	case "openid":
		return nil
	case "profile":
		return nil
	case "email":
		return nil
	case "address":
		return nil
	case "phone":
		return nil
	case "offline_access":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
