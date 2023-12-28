// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// IdpOIDCCallbackParams is parameters of idpOIDCCallback operation.
type IdpOIDCCallbackParams struct {
	// Code.
	Code string
	// State.
	State string
}

func unpackIdpOIDCCallbackParams(packed middleware.Parameters) (params IdpOIDCCallbackParams) {
	{
		key := middleware.ParameterKey{
			Name: "code",
			In:   "query",
		}
		params.Code = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "state",
			In:   "query",
		}
		params.State = packed[key].(string)
	}
	return params
}

func decodeIdpOIDCCallbackParams(args [0]string, argsEscaped bool, r *http.Request) (params IdpOIDCCallbackParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: code.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "code",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Code = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "code",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: state.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "state",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.State = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "state",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// IdpOIDCLoginParams is parameters of idpOIDCLogin operation.
type IdpOIDCLoginParams struct {
	// Op.
	Op IdpOIDCLoginOp
}

func unpackIdpOIDCLoginParams(packed middleware.Parameters) (params IdpOIDCLoginParams) {
	{
		key := middleware.ParameterKey{
			Name: "op",
			In:   "query",
		}
		params.Op = packed[key].(IdpOIDCLoginOp)
	}
	return params
}

func decodeIdpOIDCLoginParams(args [0]string, argsEscaped bool, r *http.Request) (params IdpOIDCLoginParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: op.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "op",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Op = IdpOIDCLoginOp(c)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Op.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "op",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// OpAuthorizeParams is parameters of opAuthorize operation.
type OpAuthorizeParams struct {
	// Response_type.
	ResponseType OpAuthorizeResponseType
	// Scope.
	Scope string
	// Client_id.
	ClientID url.URL
	// Http://localhost:5555/rp/callback.
	RedirectURI url.URL
	// State.
	State OptString
	// Nonce.
	Nonce OptString
}

func unpackOpAuthorizeParams(packed middleware.Parameters) (params OpAuthorizeParams) {
	{
		key := middleware.ParameterKey{
			Name: "response_type",
			In:   "query",
		}
		params.ResponseType = packed[key].(OpAuthorizeResponseType)
	}
	{
		key := middleware.ParameterKey{
			Name: "scope",
			In:   "query",
		}
		params.Scope = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "client_id",
			In:   "query",
		}
		params.ClientID = packed[key].(url.URL)
	}
	{
		key := middleware.ParameterKey{
			Name: "redirect_uri",
			In:   "query",
		}
		params.RedirectURI = packed[key].(url.URL)
	}
	{
		key := middleware.ParameterKey{
			Name: "state",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.State = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "nonce",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Nonce = v.(OptString)
		}
	}
	return params
}

func decodeOpAuthorizeParams(args [0]string, argsEscaped bool, r *http.Request) (params OpAuthorizeParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: response_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "response_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.ResponseType = OpAuthorizeResponseType(c)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.ResponseType.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "response_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: scope.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "scope",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Scope = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "scope",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: client_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "client_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToURL(val)
				if err != nil {
					return err
				}

				params.ClientID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "client_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: redirect_uri.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "redirect_uri",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToURL(val)
				if err != nil {
					return err
				}

				params.RedirectURI = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "redirect_uri",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: state.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "state",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStateVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStateVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.State.SetTo(paramsDotStateVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "state",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: nonce.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "nonce",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotNonceVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotNonceVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Nonce.SetTo(paramsDotNonceVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "nonce",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// OpCallbackParams is parameters of opCallback operation.
type OpCallbackParams struct {
	// Id.
	ID string
}

func unpackOpCallbackParams(packed middleware.Parameters) (params OpCallbackParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "query",
		}
		params.ID = packed[key].(string)
	}
	return params
}

func decodeOpCallbackParams(args [0]string, argsEscaped bool, r *http.Request) (params OpCallbackParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// OpLoginViewParams is parameters of opLoginView operation.
type OpLoginViewParams struct {
	// Auth request id.
	AuthRequestID string
}

func unpackOpLoginViewParams(packed middleware.Parameters) (params OpLoginViewParams) {
	{
		key := middleware.ParameterKey{
			Name: "auth_request_id",
			In:   "query",
		}
		params.AuthRequestID = packed[key].(string)
	}
	return params
}

func decodeOpLoginViewParams(args [0]string, argsEscaped bool, r *http.Request) (params OpLoginViewParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: auth_request_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "auth_request_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.AuthRequestID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "auth_request_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// RpCallbackParams is parameters of rpCallback operation.
type RpCallbackParams struct {
	// Code.
	Code string
	// State.
	QueryState  string
	CookieState string
}

func unpackRpCallbackParams(packed middleware.Parameters) (params RpCallbackParams) {
	{
		key := middleware.ParameterKey{
			Name: "code",
			In:   "query",
		}
		params.Code = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "state",
			In:   "query",
		}
		params.QueryState = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "state",
			In:   "cookie",
		}
		params.CookieState = packed[key].(string)
	}
	return params
}

func decodeRpCallbackParams(args [0]string, argsEscaped bool, r *http.Request) (params RpCallbackParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	c := uri.NewCookieDecoder(r)
	// Decode query: code.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "code",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Code = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "code",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: state.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "state",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.QueryState = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "state",
			In:   "query",
			Err:  err,
		}
	}
	// Decode cookie: state.
	if err := func() error {
		cfg := uri.CookieParameterDecodingConfig{
			Name:    "state",
			Explode: true,
		}
		if err := c.HasParam(cfg); err == nil {
			if err := c.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.CookieState = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "state",
			In:   "cookie",
			Err:  err,
		}
	}
	return params, nil
}
