// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func encodeHealthResponse(response HealthRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *HealthOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *HealthInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeIdpSigninResponse(response IdpSigninRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *IdpSigninOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *IdpSigninUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *IdpSigninInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeIdpSignupResponse(response IdpSignupRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *IdpSignupOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *IdpSignupInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOpAuthorizeResponse(response OpAuthorizeRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *OpAuthorizeFound:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Location" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Location",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.Location.Get(); ok {
						return e.EncodeValue(conv.URLToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Location header")
				}
			}
		}
		w.WriteHeader(302)
		span.SetStatus(codes.Ok, http.StatusText(302))

		return nil

	case *OpAuthorizeInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOpCallbackResponse(response OpCallbackRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *OpCallbackFound:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Location" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Location",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.Location.Get(); ok {
						return e.EncodeValue(conv.URLToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Location header")
				}
			}
		}
		w.WriteHeader(302)
		span.SetStatus(codes.Ok, http.StatusText(302))

		return nil

	case *OpCallbackInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOpCertsResponse(response OpCertsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *OPJWKSetResponseSchema:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *OpCertsInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOpLoginResponse(response OpLoginRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *OpLoginOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *OpLoginInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOpLoginViewResponse(response OpLoginViewRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *OpLoginViewOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *OpLoginViewInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOpOpenIDConfigurationResponse(response OpOpenIDConfigurationRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *OPOpenIDConfigurationResponseSchema:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *OpOpenIDConfigurationInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOpRevokeResponse(response OpRevokeRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *OpRevokeOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *OpRevokeBadRequest:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *OpRevokeInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOpTokenResponse(response OpTokenRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *OPTokenResponseSchema:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *OpTokenBadRequest:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *OpTokenInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOpUserinfoResponse(response OpUserinfoRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *OPUserInfoResponseSchema:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *OpUserinfoInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeRpCallbackResponse(response RpCallbackRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *RpCallbackOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *RpCallbackInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeRpLoginResponse(response RpLoginRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *RpLoginFound:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Location" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Location",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.Location.Get(); ok {
						return e.EncodeValue(conv.URLToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Location header")
				}
			}
			// Encode "Set-Cookie" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Set-Cookie",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.SetCookie.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Set-Cookie header")
				}
			}
		}
		w.WriteHeader(302)
		span.SetStatus(codes.Ok, http.StatusText(302))

		return nil

	case *RpLoginInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
