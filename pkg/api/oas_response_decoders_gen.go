// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"mime"
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func decodeHealthResponse(resp *http.Response) (res HealthRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &HealthOK{}, nil
	case 500:
		// Code 500.
		return &HealthInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeIdpOIDCCallbackResponse(resp *http.Response) (res IdpOIDCCallbackRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "text/html":
			reader := resp.Body
			b, err := io.ReadAll(reader)
			if err != nil {
				return res, err
			}

			response := IdpOIDCCallbackOK{Data: bytes.NewReader(b)}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		// Code 500.
		return &IdpOIDCCallbackInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeIdpOIDCLoginResponse(resp *http.Response) (res IdpOIDCLoginRes, _ error) {
	switch resp.StatusCode {
	case 302:
		// Code 302.
		var wrapper IdpOIDCLoginFound
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse "Location" header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "Location",
				Explode: false,
			}
			if err := func() error {
				if err := h.HasParam(cfg); err == nil {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotLocationVal url.URL
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToURL(val)
							if err != nil {
								return err
							}

							wrapperDotLocationVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.Location.SetTo(wrapperDotLocationVal)
						return nil
					}); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "parse Location header")
			}
		}
		return &wrapper, nil
	case 400:
		// Code 400.
		return &IdpOIDCLoginBadRequest{}, nil
	case 500:
		// Code 500.
		return &IdpOIDCLoginInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeIdpSigninResponse(resp *http.Response) (res IdpSigninRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &IdpSigninOK{}, nil
	case 401:
		// Code 401.
		return &IdpSigninUnauthorized{}, nil
	case 500:
		// Code 500.
		return &IdpSigninInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeIdpSignupResponse(resp *http.Response) (res IdpSignupRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &IdpSignupOK{}, nil
	case 500:
		// Code 500.
		return &IdpSignupInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOpAuthorizeResponse(resp *http.Response) (res OpAuthorizeRes, _ error) {
	switch resp.StatusCode {
	case 302:
		// Code 302.
		var wrapper OpAuthorizeFound
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse "Location" header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "Location",
				Explode: false,
			}
			if err := func() error {
				if err := h.HasParam(cfg); err == nil {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotLocationVal url.URL
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToURL(val)
							if err != nil {
								return err
							}

							wrapperDotLocationVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.Location.SetTo(wrapperDotLocationVal)
						return nil
					}); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "parse Location header")
			}
		}
		return &wrapper, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OpAuthorizeBadRequest
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 401:
		// Code 401.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OpAuthorizeUnauthorized
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OpAuthorizeForbidden
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		// Code 500.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OpAuthorizeInternalServerError
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOpCallbackResponse(resp *http.Response) (res OpCallbackRes, _ error) {
	switch resp.StatusCode {
	case 302:
		// Code 302.
		var wrapper OpCallbackFound
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse "Location" header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "Location",
				Explode: false,
			}
			if err := func() error {
				if err := h.HasParam(cfg); err == nil {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotLocationVal url.URL
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToURL(val)
							if err != nil {
								return err
							}

							wrapperDotLocationVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.Location.SetTo(wrapperDotLocationVal)
						return nil
					}); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "parse Location header")
			}
		}
		return &wrapper, nil
	case 500:
		// Code 500.
		return &OpCallbackInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOpCertsResponse(resp *http.Response) (res OpCertsRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OPJWKSetResponseSchema
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		// Code 500.
		return &OpCertsInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOpLoginResponse(resp *http.Response) (res OpLoginRes, _ error) {
	switch resp.StatusCode {
	case 302:
		// Code 302.
		var wrapper OpLoginFound
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse "Location" header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "Location",
				Explode: false,
			}
			if err := func() error {
				if err := h.HasParam(cfg); err == nil {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotLocationVal url.URL
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToURL(val)
							if err != nil {
								return err
							}

							wrapperDotLocationVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.Location.SetTo(wrapperDotLocationVal)
						return nil
					}); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "parse Location header")
			}
		}
		return &wrapper, nil
	case 500:
		// Code 500.
		return &OpLoginInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOpLoginViewResponse(resp *http.Response) (res OpLoginViewRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "text/html":
			reader := resp.Body
			b, err := io.ReadAll(reader)
			if err != nil {
				return res, err
			}

			response := OpLoginViewOK{Data: bytes.NewReader(b)}
			var wrapper OpLoginViewOKHeaders
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "X-Request-Id" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "X-Request-Id",
					Explode: false,
				}
				if err := func() error {
					if err := h.HasParam(cfg); err == nil {
						if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
							var wrapperDotXRequestIDVal string
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToString(val)
								if err != nil {
									return err
								}

								wrapperDotXRequestIDVal = c
								return nil
							}(); err != nil {
								return err
							}
							wrapper.XRequestID.SetTo(wrapperDotXRequestIDVal)
							return nil
						}); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse X-Request-Id header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		// Code 500.
		return &OpLoginViewInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOpOpenIDConfigurationResponse(resp *http.Response) (res OpOpenIDConfigurationRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OPOpenIDConfigurationResponseSchema
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		// Code 500.
		return &OpOpenIDConfigurationInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOpRevokeResponse(resp *http.Response) (res OpRevokeRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &OpRevokeOK{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OpRevokeBadRequest
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		// Code 500.
		return &OpRevokeInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOpTokenResponse(resp *http.Response) (res OpTokenRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OPTokenResponseSchema
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			var wrapper OPTokenResponseSchemaHeaders
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Cache-Control" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Cache-Control",
					Explode: false,
				}
				if err := func() error {
					if err := h.HasParam(cfg); err == nil {
						if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
							var wrapperDotCacheControlVal string
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToString(val)
								if err != nil {
									return err
								}

								wrapperDotCacheControlVal = c
								return nil
							}(); err != nil {
								return err
							}
							wrapper.CacheControl.SetTo(wrapperDotCacheControlVal)
							return nil
						}); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Cache-Control header")
				}
			}
			// Parse "Pragma" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Pragma",
					Explode: false,
				}
				if err := func() error {
					if err := h.HasParam(cfg); err == nil {
						if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
							var wrapperDotPragmaVal string
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToString(val)
								if err != nil {
									return err
								}

								wrapperDotPragmaVal = c
								return nil
							}(); err != nil {
								return err
							}
							wrapper.Pragma.SetTo(wrapperDotPragmaVal)
							return nil
						}); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Pragma header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OpTokenBadRequest
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		// Code 500.
		return &OpTokenInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOpUserinfoResponse(resp *http.Response) (res OpUserinfoRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OPUserInfoResponseSchema
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 500:
		// Code 500.
		return &OpUserinfoInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeRpCallbackResponse(resp *http.Response) (res RpCallbackRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "text/html":
			reader := resp.Body
			b, err := io.ReadAll(reader)
			if err != nil {
				return res, err
			}

			response := RpCallbackOK{Data: bytes.NewReader(b)}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		// Code 400.
		return &RpCallbackBadRequest{}, nil
	case 500:
		// Code 500.
		return &RpCallbackInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeRpLoginResponse(resp *http.Response) (res RpLoginRes, _ error) {
	switch resp.StatusCode {
	case 302:
		// Code 302.
		var wrapper RpLoginFound
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse "Location" header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "Location",
				Explode: false,
			}
			if err := func() error {
				if err := h.HasParam(cfg); err == nil {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotLocationVal url.URL
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToURL(val)
							if err != nil {
								return err
							}

							wrapperDotLocationVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.Location.SetTo(wrapperDotLocationVal)
						return nil
					}); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "parse Location header")
			}
		}
		// Parse "Set-Cookie" header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "Set-Cookie",
				Explode: false,
			}
			if err := func() error {
				if err := h.HasParam(cfg); err == nil {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotSetCookieVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							wrapperDotSetCookieVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.SetCookie.SetTo(wrapperDotSetCookieVal)
						return nil
					}); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "parse Set-Cookie header")
			}
		}
		return &wrapper, nil
	case 500:
		// Code 500.
		return &RpLoginInternalServerError{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
