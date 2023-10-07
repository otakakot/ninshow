// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeIdpSigninRequest(
	req OptIdPSigninRequestSchema,
	r *http.Request,
) error {
	const contentType = "application/json"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := jx.GetEncoder()
	{
		if req.Set {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeIdpSignupRequest(
	req OptIdPSignupRequestSchema,
	r *http.Request,
) error {
	const contentType = "application/json"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := jx.GetEncoder()
	{
		if req.Set {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeOpRevokeRequest(
	req *OPRevokeRequestSchema,
	r *http.Request,
) error {
	const contentType = "application/x-www-form-urlencoded"
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "token" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "token",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Token))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "token_type_hint" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "token_type_hint",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.TokenTypeHint.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	encoded := q.Values().Encode()
	ht.SetBody(r, strings.NewReader(encoded), contentType)
	return nil
}

func encodeOpTokenRequest(
	req *OPTokenRequestSchema,
	r *http.Request,
) error {
	const contentType = "application/x-www-form-urlencoded"
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "grant_type" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "grant_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(request.GrantType)))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "code" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "code",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Code))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "redirect_uri" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "redirect_uri",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.URLToString(request.RedirectURI))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	encoded := q.Values().Encode()
	ht.SetBody(r, strings.NewReader(encoded), contentType)
	return nil
}