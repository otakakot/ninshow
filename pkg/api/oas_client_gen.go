// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// Health invokes health operation.
	//
	// Health.
	//
	// GET /health
	Health(ctx context.Context) (HealthRes, error)
	// IdpSignin invokes idpSignin operation.
	//
	// Sign In.
	//
	// POST /idp/signin
	IdpSignin(ctx context.Context, request OptIdPSigninRequestSchema) (IdpSigninRes, error)
	// IdpSignup invokes idpSignup operation.
	//
	// Sign Up.
	//
	// POST /idp/signup
	IdpSignup(ctx context.Context, request OptIdPSignupRequestSchema) (IdpSignupRes, error)
	// OpAuthorize invokes opAuthorize operation.
	//
	// Authentication Request.
	//
	// GET /op/authorize
	OpAuthorize(ctx context.Context, params OpAuthorizeParams) (OpAuthorizeRes, error)
	// OpCallback invokes opCallback operation.
	//
	// OP Callback.
	//
	// GET /op/callback
	OpCallback(ctx context.Context, params OpCallbackParams) (OpCallbackRes, error)
	// OpCerts invokes opCerts operation.
	//
	// Https://openid-foundation-japan.github.io/rfc7517.ja.html.
	//
	// GET /op/certs
	OpCerts(ctx context.Context) (OpCertsRes, error)
	// OpLogin invokes opLogin operation.
	//
	// OP Login.
	//
	// POST /op/login
	OpLogin(ctx context.Context) (OpLoginRes, error)
	// OpLoginView invokes opLoginView operation.
	//
	// OP Login.
	//
	// GET /op/login
	OpLoginView(ctx context.Context, params OpLoginViewParams) (OpLoginViewRes, error)
	// OpOpenIDConfiguration invokes opOpenIDConfiguration operation.
	//
	// OpenID Provider Configuration.
	//
	// GET /op/.well-known/openid-configuration
	OpOpenIDConfiguration(ctx context.Context) (OpOpenIDConfigurationRes, error)
	// OpRevoke invokes opRevoke operation.
	//
	// Https://openid.net/specs/openid-connect-core-1_0.html#Revocation.
	//
	// POST /op/revoke
	OpRevoke(ctx context.Context, request *OPRevokeRequestSchema) (OpRevokeRes, error)
	// OpToken invokes opToken operation.
	//
	// Https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenRequest.
	//
	// POST /op/token
	OpToken(ctx context.Context, request *OPTokenRequestSchema) (OpTokenRes, error)
	// OpUserinfo invokes opUserinfo operation.
	//
	// Https://openid.net/specs/openid-connect-core-1_0.html#UserInfo.
	//
	// GET /op/userinfo
	OpUserinfo(ctx context.Context) (OpUserinfoRes, error)
	// RpCallback invokes rpCallback operation.
	//
	// RP Callback.
	//
	// GET /rp/callback
	RpCallback(ctx context.Context, params RpCallbackParams) (RpCallbackRes, error)
	// RpLogin invokes rpLogin operation.
	//
	// RP Login.
	//
	// GET /rp/login
	RpLogin(ctx context.Context) (RpLoginRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		sec:        sec,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// Health invokes health operation.
//
// Health.
//
// GET /health
func (c *Client) Health(ctx context.Context) (HealthRes, error) {
	res, err := c.sendHealth(ctx)
	return res, err
}

func (c *Client) sendHealth(ctx context.Context) (res HealthRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("health"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/health"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "Health",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/health"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeHealthResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// IdpSignin invokes idpSignin operation.
//
// Sign In.
//
// POST /idp/signin
func (c *Client) IdpSignin(ctx context.Context, request OptIdPSigninRequestSchema) (IdpSigninRes, error) {
	res, err := c.sendIdpSignin(ctx, request)
	return res, err
}

func (c *Client) sendIdpSignin(ctx context.Context, request OptIdPSigninRequestSchema) (res IdpSigninRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("idpSignin"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/idp/signin"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "IdpSignin",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/idp/signin"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeIdpSigninRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeIdpSigninResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// IdpSignup invokes idpSignup operation.
//
// Sign Up.
//
// POST /idp/signup
func (c *Client) IdpSignup(ctx context.Context, request OptIdPSignupRequestSchema) (IdpSignupRes, error) {
	res, err := c.sendIdpSignup(ctx, request)
	return res, err
}

func (c *Client) sendIdpSignup(ctx context.Context, request OptIdPSignupRequestSchema) (res IdpSignupRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("idpSignup"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/idp/signup"),
	}
	// Validate request before sending.
	if err := func() error {
		if value, ok := request.Get(); ok {
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
		return res, errors.Wrap(err, "validate")
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "IdpSignup",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/idp/signup"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeIdpSignupRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeIdpSignupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OpAuthorize invokes opAuthorize operation.
//
// Authentication Request.
//
// GET /op/authorize
func (c *Client) OpAuthorize(ctx context.Context, params OpAuthorizeParams) (OpAuthorizeRes, error) {
	res, err := c.sendOpAuthorize(ctx, params)
	return res, err
}

func (c *Client) sendOpAuthorize(ctx context.Context, params OpAuthorizeParams) (res OpAuthorizeRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("opAuthorize"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/op/authorize"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "OpAuthorize",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/op/authorize"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "response_type" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "response_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(params.ResponseType)))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "scope" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "scope",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range params.Scope {
					if err := func() error {
						return e.EncodeValue(conv.StringToString(string(item)))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "client_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "client_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.URLToString(params.ClientID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "redirect_uri" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "redirect_uri",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.URLToString(params.RedirectURI))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "state" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "state",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.State.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "nonce" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "nonce",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Nonce.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOpAuthorizeResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OpCallback invokes opCallback operation.
//
// OP Callback.
//
// GET /op/callback
func (c *Client) OpCallback(ctx context.Context, params OpCallbackParams) (OpCallbackRes, error) {
	res, err := c.sendOpCallback(ctx, params)
	return res, err
}

func (c *Client) sendOpCallback(ctx context.Context, params OpCallbackParams) (res OpCallbackRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("opCallback"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/op/callback"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "OpCallback",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/op/callback"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.ID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOpCallbackResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OpCerts invokes opCerts operation.
//
// Https://openid-foundation-japan.github.io/rfc7517.ja.html.
//
// GET /op/certs
func (c *Client) OpCerts(ctx context.Context) (OpCertsRes, error) {
	res, err := c.sendOpCerts(ctx)
	return res, err
}

func (c *Client) sendOpCerts(ctx context.Context) (res OpCertsRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("opCerts"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/op/certs"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "OpCerts",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/op/certs"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOpCertsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OpLogin invokes opLogin operation.
//
// OP Login.
//
// POST /op/login
func (c *Client) OpLogin(ctx context.Context) (OpLoginRes, error) {
	res, err := c.sendOpLogin(ctx)
	return res, err
}

func (c *Client) sendOpLogin(ctx context.Context) (res OpLoginRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("opLogin"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/op/login"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "OpLogin",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/op/login"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOpLoginResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OpLoginView invokes opLoginView operation.
//
// OP Login.
//
// GET /op/login
func (c *Client) OpLoginView(ctx context.Context, params OpLoginViewParams) (OpLoginViewRes, error) {
	res, err := c.sendOpLoginView(ctx, params)
	return res, err
}

func (c *Client) sendOpLoginView(ctx context.Context, params OpLoginViewParams) (res OpLoginViewRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("opLoginView"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/op/login"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "OpLoginView",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/op/login"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "auth_request_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "auth_request_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AuthRequestID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOpLoginViewResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OpOpenIDConfiguration invokes opOpenIDConfiguration operation.
//
// OpenID Provider Configuration.
//
// GET /op/.well-known/openid-configuration
func (c *Client) OpOpenIDConfiguration(ctx context.Context) (OpOpenIDConfigurationRes, error) {
	res, err := c.sendOpOpenIDConfiguration(ctx)
	return res, err
}

func (c *Client) sendOpOpenIDConfiguration(ctx context.Context) (res OpOpenIDConfigurationRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("opOpenIDConfiguration"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/op/.well-known/openid-configuration"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "OpOpenIDConfiguration",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/op/.well-known/openid-configuration"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOpOpenIDConfigurationResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OpRevoke invokes opRevoke operation.
//
// Https://openid.net/specs/openid-connect-core-1_0.html#Revocation.
//
// POST /op/revoke
func (c *Client) OpRevoke(ctx context.Context, request *OPRevokeRequestSchema) (OpRevokeRes, error) {
	res, err := c.sendOpRevoke(ctx, request)
	return res, err
}

func (c *Client) sendOpRevoke(ctx context.Context, request *OPRevokeRequestSchema) (res OpRevokeRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("opRevoke"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/op/revoke"),
	}
	// Validate request before sending.
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "OpRevoke",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/op/revoke"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeOpRevokeRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOpRevokeResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OpToken invokes opToken operation.
//
// Https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenRequest.
//
// POST /op/token
func (c *Client) OpToken(ctx context.Context, request *OPTokenRequestSchema) (OpTokenRes, error) {
	res, err := c.sendOpToken(ctx, request)
	return res, err
}

func (c *Client) sendOpToken(ctx context.Context, request *OPTokenRequestSchema) (res OpTokenRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("opToken"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/op/token"),
	}
	// Validate request before sending.
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "OpToken",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/op/token"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeOpTokenRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOpTokenResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OpUserinfo invokes opUserinfo operation.
//
// Https://openid.net/specs/openid-connect-core-1_0.html#UserInfo.
//
// GET /op/userinfo
func (c *Client) OpUserinfo(ctx context.Context) (OpUserinfoRes, error) {
	res, err := c.sendOpUserinfo(ctx)
	return res, err
}

func (c *Client) sendOpUserinfo(ctx context.Context) (res OpUserinfoRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("opUserinfo"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/op/userinfo"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "OpUserinfo",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/op/userinfo"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:Bearer"
			switch err := c.securityBearer(ctx, "OpUserinfo", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"Bearer\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOpUserinfoResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RpCallback invokes rpCallback operation.
//
// RP Callback.
//
// GET /rp/callback
func (c *Client) RpCallback(ctx context.Context, params RpCallbackParams) (RpCallbackRes, error) {
	res, err := c.sendRpCallback(ctx, params)
	return res, err
}

func (c *Client) sendRpCallback(ctx context.Context, params RpCallbackParams) (res RpCallbackRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("rpCallback"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/rp/callback"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "RpCallback",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/rp/callback"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "code" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "code",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Code))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "state" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "state",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.State))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeRpCallbackResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RpLogin invokes rpLogin operation.
//
// RP Login.
//
// GET /rp/login
func (c *Client) RpLogin(ctx context.Context) (RpLoginRes, error) {
	res, err := c.sendRpLogin(ctx)
	return res, err
}

func (c *Client) sendRpLogin(ctx context.Context) (res RpLoginRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("rpLogin"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/rp/login"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "RpLogin",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/rp/login"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeRpLoginResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
