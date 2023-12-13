// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'h': // Prefix: "health"
				if l := len("health"); len(elem) >= l && elem[0:l] == "health" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleHealthRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'i': // Prefix: "idp/"
				if l := len("idp/"); len(elem) >= l && elem[0:l] == "idp/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'o': // Prefix: "oidc"
					if l := len("oidc"); len(elem) >= l && elem[0:l] == "oidc" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleIdpOIDCRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 's': // Prefix: "sign"
					if l := len("sign"); len(elem) >= l && elem[0:l] == "sign" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'i': // Prefix: "in"
						if l := len("in"); len(elem) >= l && elem[0:l] == "in" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleIdpSigninRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 'u': // Prefix: "up"
						if l := len("up"); len(elem) >= l && elem[0:l] == "up" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleIdpSignupRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				}
			case 'o': // Prefix: "op/"
				if l := len("op/"); len(elem) >= l && elem[0:l] == "op/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '.': // Prefix: ".well-known/openid-configuration"
					if l := len(".well-known/openid-configuration"); len(elem) >= l && elem[0:l] == ".well-known/openid-configuration" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleOpOpenIDConfigurationRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'a': // Prefix: "authorize"
					if l := len("authorize"); len(elem) >= l && elem[0:l] == "authorize" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleOpAuthorizeRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'c': // Prefix: "c"
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "allback"
						if l := len("allback"); len(elem) >= l && elem[0:l] == "allback" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleOpCallbackRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 'e': // Prefix: "erts"
						if l := len("erts"); len(elem) >= l && elem[0:l] == "erts" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleOpCertsRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				case 'l': // Prefix: "login"
					if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleOpLoginViewRequest([0]string{}, elemIsEscaped, w, r)
						case "POST":
							s.handleOpLoginRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET,POST")
						}

						return
					}
				case 'r': // Prefix: "revoke"
					if l := len("revoke"); len(elem) >= l && elem[0:l] == "revoke" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleOpRevokeRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 't': // Prefix: "token"
					if l := len("token"); len(elem) >= l && elem[0:l] == "token" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleOpTokenRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'u': // Prefix: "userinfo"
					if l := len("userinfo"); len(elem) >= l && elem[0:l] == "userinfo" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleOpUserinfoRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'r': // Prefix: "rp/"
				if l := len("rp/"); len(elem) >= l && elem[0:l] == "rp/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "callback"
					if l := len("callback"); len(elem) >= l && elem[0:l] == "callback" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleRpCallbackRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'l': // Prefix: "login"
					if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleRpLoginRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'h': // Prefix: "health"
				if l := len("health"); len(elem) >= l && elem[0:l] == "health" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: Health
						r.name = "Health"
						r.summary = "Health"
						r.operationID = "health"
						r.pathPattern = "/health"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'i': // Prefix: "idp/"
				if l := len("idp/"); len(elem) >= l && elem[0:l] == "idp/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'o': // Prefix: "oidc"
					if l := len("oidc"); len(elem) >= l && elem[0:l] == "oidc" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: IdpOIDC
							r.name = "IdpOIDC"
							r.summary = "OpenID Connect"
							r.operationID = "idpOIDC"
							r.pathPattern = "/idp/oidc"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 's': // Prefix: "sign"
					if l := len("sign"); len(elem) >= l && elem[0:l] == "sign" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'i': // Prefix: "in"
						if l := len("in"); len(elem) >= l && elem[0:l] == "in" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: IdpSignin
								r.name = "IdpSignin"
								r.summary = "Sign In"
								r.operationID = "idpSignin"
								r.pathPattern = "/idp/signin"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'u': // Prefix: "up"
						if l := len("up"); len(elem) >= l && elem[0:l] == "up" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: IdpSignup
								r.name = "IdpSignup"
								r.summary = "Sign Up"
								r.operationID = "idpSignup"
								r.pathPattern = "/idp/signup"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'o': // Prefix: "op/"
				if l := len("op/"); len(elem) >= l && elem[0:l] == "op/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '.': // Prefix: ".well-known/openid-configuration"
					if l := len(".well-known/openid-configuration"); len(elem) >= l && elem[0:l] == ".well-known/openid-configuration" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: OpOpenIDConfiguration
							r.name = "OpOpenIDConfiguration"
							r.summary = "OpenID Provider Configuration"
							r.operationID = "opOpenIDConfiguration"
							r.pathPattern = "/op/.well-known/openid-configuration"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'a': // Prefix: "authorize"
					if l := len("authorize"); len(elem) >= l && elem[0:l] == "authorize" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: OpAuthorize
							r.name = "OpAuthorize"
							r.summary = "Authentication Request"
							r.operationID = "opAuthorize"
							r.pathPattern = "/op/authorize"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'c': // Prefix: "c"
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "allback"
						if l := len("allback"); len(elem) >= l && elem[0:l] == "allback" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: OpCallback
								r.name = "OpCallback"
								r.summary = "OP Callback"
								r.operationID = "opCallback"
								r.pathPattern = "/op/callback"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'e': // Prefix: "erts"
						if l := len("erts"); len(elem) >= l && elem[0:l] == "erts" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: OpCerts
								r.name = "OpCerts"
								r.summary = "OP JWK Set"
								r.operationID = "opCerts"
								r.pathPattern = "/op/certs"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 'l': // Prefix: "login"
					if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: OpLoginView
							r.name = "OpLoginView"
							r.summary = "OP Login"
							r.operationID = "opLoginView"
							r.pathPattern = "/op/login"
							r.args = args
							r.count = 0
							return r, true
						case "POST":
							// Leaf: OpLogin
							r.name = "OpLogin"
							r.summary = "OP Login"
							r.operationID = "opLogin"
							r.pathPattern = "/op/login"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'r': // Prefix: "revoke"
					if l := len("revoke"); len(elem) >= l && elem[0:l] == "revoke" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: OpRevoke
							r.name = "OpRevoke"
							r.summary = "OP Revocation Request"
							r.operationID = "opRevoke"
							r.pathPattern = "/op/revoke"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 't': // Prefix: "token"
					if l := len("token"); len(elem) >= l && elem[0:l] == "token" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: OpToken
							r.name = "OpToken"
							r.summary = "OP Token Request"
							r.operationID = "opToken"
							r.pathPattern = "/op/token"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'u': // Prefix: "userinfo"
					if l := len("userinfo"); len(elem) >= l && elem[0:l] == "userinfo" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: OpUserinfo
							r.name = "OpUserinfo"
							r.summary = "UserInfo Request"
							r.operationID = "opUserinfo"
							r.pathPattern = "/op/userinfo"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'r': // Prefix: "rp/"
				if l := len("rp/"); len(elem) >= l && elem[0:l] == "rp/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "callback"
					if l := len("callback"); len(elem) >= l && elem[0:l] == "callback" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: RpCallback
							r.name = "RpCallback"
							r.summary = "RP Callback"
							r.operationID = "rpCallback"
							r.pathPattern = "/rp/callback"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'l': // Prefix: "login"
					if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: RpLogin
							r.name = "RpLogin"
							r.summary = "RP Login"
							r.operationID = "rpLogin"
							r.pathPattern = "/rp/login"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			}
		}
	}
	return r, false
}
